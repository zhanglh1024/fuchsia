package battleobj

import (
	"beginner-server/app"
	"beginner-server/app/model/dyn"
)


func InsertDateIntoMatchBattle(info dyn.MatchBattle) error {
	return app.DB.Table("match_battles").Create(&info).Error
}

func GetMatchBattleData( roleId string) (dyn.MatchBattle, error) {
	info := dyn.MatchBattle{}
	err := app.DB.Table("match_battles").Where("role_id = ?", roleId).Find(&info).Error
	return info, err
}

func DeleteMatchBattleData(roleId string)error  {
	info := dyn.MatchBattle{}
	return app.DB.Table("match_battles").Where("role_id = ?", roleId).Delete(&info).Error
}


//需要把数据设置成初始化数据跳这跳更新方法
func UpdateMatchBattleData(info dyn.MatchBattle) error {
	return app.DB.Table("match_battles").Where("role_id = ?", info.RoleId).Update(
		map[string]interface{}{"hero_no": info.HeroNo, "soldier": info.Soldier, "register": info.Register,"integral":info.Integral}).Error

}

//只更新传入数值的掉这条
func UpdateRegisterMatchBattleData(info dyn.MatchBattle) error {
	return app.DB.Table("match_battles").Where("role_id = ?", info.RoleId).Update(&info).Error
}

func GetAllMatchBattleInfo() ([]dyn.MatchBattle, error){
	Info := []dyn.MatchBattle{}
	err:=app.DB.Where("register = 1").Find(&Info).Error
	return Info, err
}


