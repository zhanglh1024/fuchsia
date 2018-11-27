package battleobj

import (
	"beginner-server/app"
	"beginner-server/app/model/dyn"
)


func InsertLevelBattleValue(value dyn.LevelBattle) error  {
	return app.DB.Create(value).Error
}

func UpdateLevelBattleValue(value dyn.LevelBattle) error {
	err := app.DB.Table("level_battles").Where("role_id = ?", value.RoleId).Update(&value).Error
	return err
}

func GetLevelBattleValue(roleId string)(dyn.LevelBattle, error) {
	Info := dyn.LevelBattle{}
	err := app.DB.Table("level_battles").Where("role_id = ?", roleId).Find(&Info).Error
	return Info, err
}

func DeleteLevelBattleValue(roleId string)error{
	return app.DB.Table("level_battles").Where("role_id = ?", roleId).Delete(&dyn.LevelBattle{}).Error
}
