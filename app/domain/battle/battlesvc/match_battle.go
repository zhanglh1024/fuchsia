package battlesvc

import (
	"beginner-server/app/domain/battle/battleobj"
	"beginner-server/app/model/dyn"
	"math/rand"
	"time"
)


//注册匹配大厅
func RegisterMatchBattleHall(roleId string, heroNo, soldierNum int) error {
	battle := dyn.MatchBattle{
		RoleId:roleId,
		HeroNo:heroNo,
		Soldier:soldierNum,
		Register:1,
	}
	return battleobj.UpdateRegisterMatchBattleData(battle)
}

//获取匹配玩家信息
func GetMatchBattleInfo(roleId string) (dyn.MatchBattle, error) {
	return battleobj.GetMatchBattleData(roleId)
}

//跟新注册大厅玩家信息，更新所有字段没穿入的更新为默认值默认
func UpdateMatchBattleInfo(info dyn.MatchBattle) error {
	err := battleobj.UpdateMatchBattleData(info)
	return err
}

//更新用户匹配数据只更新传入字段
func UPdateIntegralMatchBattleInfor(info dyn.MatchBattle) error {
	return  battleobj.UpdateRegisterMatchBattleData(info)
}


//随机匹配注册大厅防守玩家
func GetRandMatchBattleOpponent()(dyn.MatchBattle, error) {
	rand.Seed(time.Now().Unix())
	battleInfo, err := battleobj.GetAllMatchBattleInfo()
	matchInfo := dyn.MatchBattle{}
	if err != nil{
		return matchInfo, err
	}
	index := rand.Intn(len(battleInfo))
	matchInfo = battleInfo[index]
	return  matchInfo,nil
}

func InsertForTest(info dyn.MatchBattle)  error{
	return battleobj.InsertDateIntoMatchBattle(info)
}

func DeleteForTest(roleId string) error {
	return battleobj.DeleteMatchBattleData(roleId)
}
