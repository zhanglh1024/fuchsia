package actions

import (
	"github.com/cxr29/log"
	"beginner-server/app"
	"beginner-server/app/domain/battle/battlesvc"
	"beginner-server/app/domain/hero/herosvc"
	"beginner-server/app/domain/resource/resourcesvc"
	"beginner-server/app/model/dyn"
	"net/http"
	"sync"
)

//MatchFight路由的请求参数字段
type MatchFightParam struct {
	UserId string //用户ID
	HeroNo int    //出战英雄编号
}

//MatchFight路由的响应字段
type MatchFightResp struct {
	Code     		app.CodeErrorType	    	//返回码 1解析数据出错 2数据库处理出错
	TargetUserId 	string	//匹配到的对手的ID
	TargetSolider   int	 //对手当前登记剩余的士兵数
	IsSuccess 		int 	//是否战斗胜利，1为胜利，0为失败
	Resource        []app.ResourceInfo //	玩家当前拥有资源
}

type MatchBattleData struct {
	MatchInfo    dyn.MatchBattle 	//匹配到的用户数据
	MyData		 OwnData			//玩家的用户数据
}

type OwnData struct {
	RoleId     string
	FightCount int
	Soldier    int
	Integral   int
}


var (
	matchLock = MatchLock{
		lock:sync.Mutex{},
		isMatchedRole:make([]string, 0),
	}

	AllRoleId = []string{""}
)


//MatchFight路由的Handler
//	MatchFight的功能是，如果传入的MyName不为空字符传，则返回<问候语>+<请求者名字>，
//	若MyName为空字符串，则返回<问候语>+"Stranger"
func MatchFightHandler(w http.ResponseWriter, r *http.Request) {
	param := MatchFightParam{}
	resp := MatchFightResp{
		Code: 0,
		TargetUserId:"",
		TargetSolider:0,
		IsSuccess:1,
		Resource:[]app.ResourceInfo{},
	}
	err := parseRequest(r, &param)
	if err != nil {
		resp.Code = app.ParseRequestError
	}

	resp, err = MatchFightOperate(param.UserId, param.HeroNo)
	if err != nil{
		log.Error(err)
	}

	writeResponse(w, resp)
}

func MatchFightOperate(userId string, heroNo int) (MatchFightResp, error) {
	resp := MatchFightResp{}
	matchBattleData := MatchBattleData{}
	matchBattleData.MyData.RoleId = userId


	matchBattleData.SetOpponentDate()

	resp.TargetUserId = matchBattleData.MatchInfo.RoleId

	fightCount, err := herosvc.GetHeroFightByRoleIdAndHeroNo(userId,heroNo)
	if err != nil{
		resp.Code = app.SearchDataError
	}

	resourceInfo, err := resourcesvc.GetRoleResourceInfoByRoleId(userId)
	if err != nil{
		resp.Code = app.SearchDataError
	}

	OwnRegisterData, err := battlesvc.GetMatchBattleInfo(userId)
	if err != nil{
		resp.Code = app.SearchDataError
	}


	//玩家的英雄战斗力和士兵数和积分
	matchBattleData.MyData.FightCount = fightCount
	matchBattleData.MyData.Soldier = int(resourceInfo.Soldier)
	matchBattleData.MyData.Integral = OwnRegisterData.Integral

	win,err := matchBattleData.FightData(matchBattleData.MyData)

	resp.IsSuccess = win
	resp.TargetSolider = matchBattleData.MatchInfo.Soldier

	resp.Resource, err = resourcesvc.GetResourceInfoByRoleId(userId)

	matchLock.Unlock(matchBattleData.MatchInfo.RoleId)

	return resp, err
}

func (OD *MatchBattleData)FightData(ownInfo OwnData) (int,error) {


	var IsSuccess = 0
	fightCount, err := herosvc.GetHeroFightByRoleIdAndHeroNo(OD.MatchInfo.RoleId, OD.MatchInfo.HeroNo)
	if err != nil {
		log.Error(err)
	}

	//玩家的战斗力
	roleFight := herosvc.CalculateHeroCombatPower(ownInfo.Soldier,ownInfo.FightCount)
	matchFight := herosvc.CalculateHeroCombatPower(OD.MatchInfo.Soldier, fightCount)

	//胜利扣除拥有总士兵数 * (5 / 英雄战斗力) + 5 * 关卡关数
	//失败扣除拥有总士兵数 * (8 / 英雄战斗力) + 5 * 关卡关数
	if roleFight > matchFight{
		IsSuccess = 1

		WinCost, LostCost := herosvc.ConsumingSoldiers(ownInfo.Soldier,
			ownInfo.FightCount,1,OD.MatchInfo.Soldier,fightCount)

		OD.FightResultSave(LostCost,IsSuccess)
		resourcesvc.UpdateMatchBattleData(WinCost,ownInfo.RoleId,ownInfo.Soldier)
	}else {
		IsSuccess = 0

		WinCost, LostCost := herosvc.ConsumingSoldiers(OD.MatchInfo.Soldier,
			fightCount,1,ownInfo.Soldier, ownInfo.FightCount)

		OD.FightResultSave(WinCost,IsSuccess)
		resourcesvc.UpdateMatchBattleData(LostCost,ownInfo.RoleId,ownInfo.Soldier)
	}

	//战斗结果数据保存到数据库
	err = battlesvc.UpdateMatchBattleInfo(OD.GetMatchBattleInfo())
	if err != nil{
		log.Errorf("战斗数据保存到数据库失败:%s", err)
	}

	ownData := dyn.MatchBattle{}
	ownData.RoleId = OD.MyData.RoleId
	ownData.Integral = OD.MyData.Integral
	err = battlesvc.UPdateIntegralMatchBattleInfor(ownData)
	if err != nil{
		log.Errorf("战斗数据保存到数据库失败:%s", err)
	}

	return IsSuccess ,err

}

//匹配游戏对手
func (OD *MatchBattleData)SetOpponentDate()  {

	matchInfo, err := battlesvc.GetRandMatchBattleOpponent()
	if err != nil {
		log.Errorf("随机匹配对手出错:%s", err)
	}

	if matchInfo.RoleId == OD.MyData.RoleId{
		log.Info("匹配到自己账号，请重新匹配！")
		OD.SetOpponentDate()
	}else 	if !matchLock.Lock(matchInfo.RoleId) {
		log.Info("该用户已被匹配，请重新匹配！")
		OD.SetOpponentDate()
	}else {
		log.Info("匹配用户成功")
		OD.MatchInfo = matchInfo
	}


}


//战斗结果数据保存
func (OD *MatchBattleData)FightResultSave(costSoldier int, isSuccess int)  {
	//士兵消耗完推出防守大厅
	var RefreshMatchInfo dyn.MatchBattle
	if costSoldier >= OD.MatchInfo.Soldier{
		integral := OD.MatchInfo.Integral

		if isSuccess == 0{
			integral += 1
		}else {
			OD.MyData.Integral += 1
		}
		RefreshMatchInfo = dyn.MatchBattle{
			RoleId:   OD.MatchInfo.RoleId,
			HeroNo:   0,
			Soldier:  0,
			Register: 0,
			Integral: integral,
		}
	}else {
		integral := OD.MatchInfo.Integral

		if isSuccess == 0{
			integral += 1
		}else {
			OD.MyData.Integral += 1
		}

		leftSoldier := OD.MatchInfo.Soldier - costSoldier


		RefreshMatchInfo = dyn.MatchBattle{
			RoleId:   OD.MatchInfo.RoleId,
			HeroNo:   OD.MatchInfo.HeroNo,
			Soldier:  leftSoldier,
			Register: OD.MatchInfo.Register,
			Integral: integral,
		}
	}

	OD.MatchInfo = RefreshMatchInfo

}

func (OD *MatchBattleData)GetMatchBattleInfo() dyn.MatchBattle {
	return OD.MatchInfo
}



