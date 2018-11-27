package actions

import (
	"github.com/cxr29/log"
	"errors"
	"beginner-server/app"
	"beginner-server/app/domain/battle/battlesvc"
	"beginner-server/app/domain/hero/herosvc"
	"beginner-server/app/domain/resource/resourcesvc"
	"beginner-server/app/model/dyn"
	"net/http"
)

//SectionFight路由的请求参数字段
type SectionFightParam struct {
	UserId string //用户ID
	HeroNo int 	  //出战英雄编号
}

//SectionFight路由的响应字段
type SectionFightResp struct {
	Code      app.CodeErrorType	    				//返回码 2 获取资源失败 9玩家已经通关
	IsSuccess int     				//是否战斗胜利，1为胜利，0为失败
	Resource  []app.ResourceInfo 	//玩家当前拥有资源
}

//SectionFight路由的Handler
//	SectionFight的功能是，如果传入正确的玩家编号和英雄编号操作成功后则返回是否战斗成功和玩家资源编号，
//	若出错则返回错误的返回码，不返回用户资源信息
func SectionFightHandler(w http.ResponseWriter, r *http.Request) {
	param := SectionFightParam{}
	resp := SectionFightResp{
		Code: 0,
		IsSuccess: 0,
		Resource: []app.ResourceInfo{},
	}
	err := parseRequest(r, &param)
	if err != nil {
		resp.Code = app.ParseRequestError
	}

	resp , err = SectionFightOperate(param.UserId, param.HeroNo)
	if err != nil{
		log.Error(err)
	}

	writeResponse(w, resp)
}

func SectionFightOperate(roleId string, herNo int)	(resp SectionFightResp, err error)  {
	resourceInfo , err := resourcesvc.GetRoleResourceInfoByRoleId(roleId)
	if err != nil{
		resp.Code = app.SearchDataError
		err = errors.New("獲取用戶資源信息失敗！！！")
		return
	}

	heroFight, err := herosvc.GetHeroFightByRoleIdAndHeroNo(roleId, herNo)

	if err != nil{
		resp.Code = app.SearchDataError
		err =  errors.New("獲取英雄失敗")
		return
	}

	levelInfo ,err := battlesvc.GetCurrentLevelInfo(roleId)
	if err != nil{
		resp.Code = app.SearchDataError
		err =  errors.New("獲取關卡信息失敗！！")
		return
	}

	if levelInfo.RoundId == 6{
		resp.Code = app.AlreadyPassAllLevel
		err = errors.New("玩家已经通关了")
		return
	}

	//玩家的战斗力
	roleFight := herosvc.CalculateHeroCombatPower(int(resourceInfo.Soldier),heroFight)
	levelFight := herosvc.CalculateHeroCombatPower(levelInfo.SoldierLeft, levelInfo.HeroFight)

	var err2,err1 error
	if roleFight > levelFight{
		log.Info("关卡战斗取胜了")
		resp.IsSuccess = 1

		roleCostSoldier, levelCostSoldier := herosvc.ConsumingSoldiers(int(resourceInfo.Soldier),
			heroFight,levelInfo.RoundId,levelInfo.SoldierLeft,levelInfo.HeroFight)
		err1,err2 = ConsumingSoldiersDataOperate(levelCostSoldier, roleCostSoldier,levelInfo,resourceInfo)

	}else {
		log.Info("关卡战斗失败了")
		resp.IsSuccess = 0
		levelCostSoldier, roleCostSoldier := herosvc.ConsumingSoldiers(levelInfo.SoldierLeft,
			levelInfo.HeroFight,levelInfo.RoundId,int(resourceInfo.Soldier),heroFight)
		err1,err2 = ConsumingSoldiersDataOperate(levelCostSoldier, roleCostSoldier,levelInfo,resourceInfo)
	}


	if err1 != nil || err2 != nil{
		resp.Code = app.UpdateDataError
		err =  errors.New("关卡或用户资源重新设置失败！！！")
		return
	}

	resp.Resource, err1 = resourcesvc.GetResourceInfoByRoleId(roleId)

	return
}

func ConsumingSoldiersDataOperate(levelCostSoldier, roleCostSoldier int,levelInfo dyn.LevelBattle,resourceInfo dyn.ResourceInfoDyn)(err,err1 error){
	if levelCostSoldier > levelInfo.SoldierLeft{
		log.Info("通关",levelCostSoldier,roleCostSoldier)
		err,_ = battlesvc.PassEditLevelInfo(levelInfo)
		err1 = resourcesvc.UpdateMatchBattleData(roleCostSoldier,resourceInfo.RoleId, int(resourceInfo.Soldier))
	}else {
		log.Info("未通关",levelCostSoldier,roleCostSoldier)
		err = battlesvc.EditCurrentLevel(levelCostSoldier, levelInfo)
		err1 = resourcesvc.UpdateMatchBattleData(roleCostSoldier,resourceInfo.RoleId, int(resourceInfo.Soldier))
	}
	return
}


