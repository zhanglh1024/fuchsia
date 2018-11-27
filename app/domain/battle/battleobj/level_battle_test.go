package battleobj

import (
	"beginner-server/app/model/dyn"
	"fmt"
	"testing"
)



func TestInsertLevelBattleValue(t *testing.T) {
	info := dyn.LevelBattle{}
	info.RoleId = "33"
	info.RoundId = 1
	info.HeroFight = 15
	info.SoldierLeft = 80
	err := InsertLevelBattleValue(info)
	if err != nil{
		t.Error(err)
	}

	err = DeleteLevelBattleValue("33")
	if err != nil{
		t.Error(err)
	}
}

func TestGetLevelBattleValue(t *testing.T) {
	info, err := GetLevelBattleValue("23")
	if err != nil{
		t.Error(err)
	}
	fmt.Println(info)
}

func TestUpdateLevelBattleValue(t *testing.T) {
	info := dyn.LevelBattle{}
	info.RoleId = "33"
	info.RoundId = 1
	info.HeroFight = 15
	info.SoldierLeft = 80
	err := InsertLevelBattleValue(info)
	if err != nil{
		t.Error(err)
	}

	info.SoldierLeft = 72
	info.RoundId = 2
	err = UpdateLevelBattleValue(info)
	if err != nil{
		t.Error(err)
	}

	err = DeleteLevelBattleValue("33")
	if err != nil{
		t.Error(err)
	}
}

func TestDeleteLevelBattleValue(t *testing.T) {
	info := dyn.LevelBattle{}
	info.RoleId = "33"
	info.RoundId = 1
	info.HeroFight = 15
	info.SoldierLeft = 80
	err := InsertLevelBattleValue(info)
	if err != nil{
		t.Error(err)
	}

	err = DeleteLevelBattleValue("33")
	if err != nil{
		t.Error(err)
	}
}
