package battleobj

import (
	"beginner-server/app/model/dyn"
	"fmt"
	"testing"
)



func TestInsertDateIntoMatchBattle(t *testing.T) {
	info := dyn.MatchBattle{}
	info.RoleId = "1"
	info.HeroNo = 0
	info.Soldier = 0
	info.Register = 0
	info.Integral = 0
	err := InsertDateIntoMatchBattle(info)
	if err != nil {
		t.Error(err)
	}

	err = DeleteMatchBattleData("1")
	if err != nil{
		t.Error(err)
	}
}

func TestGetMatchBattleData(t *testing.T) {

	info := dyn.MatchBattle{}
	info.RoleId = "9"
	info.Soldier = 10
	info.Integral = 12
	info.Register = 0
	info.HeroNo = 9
	err := InsertDateIntoMatchBattle(info)
	if err != nil {
		t.Error(err)
	}

	infos, err := GetMatchBattleData("9")
	if err != nil{
		t.Error(err)
	}
	fmt.Println(infos)

	err = DeleteMatchBattleData("9")
	if err != nil{
		t.Error(err)
	}
}

func TestUpdateMatchBattleData(t *testing.T) {
	info := dyn.MatchBattle{}
	info.RoleId = "9"
	info.Soldier = 10
	info.Integral = 12
	info.Register = 0
	info.HeroNo = 9
	err := InsertDateIntoMatchBattle(info)
	if err != nil {
		t.Error(err)
	}

	info.Soldier = 99

	err = UpdateMatchBattleData(info)
	if err != nil{
		t.Error(err)
	}
	err = DeleteMatchBattleData("9")
	if err != nil{
		t.Error(err)
	}

}

func TestUpdateRegisterMatchBattleData(t *testing.T) {
	info := dyn.MatchBattle{}
	info.RoleId = "9"
	info.Soldier = 90
	info.Integral = 12
	info.Register = 1
	info.HeroNo = 10
	err := InsertDateIntoMatchBattle(info)
	if err != nil {
		t.Error(err)
	}

	info.Soldier = 129

	err = UpdateRegisterMatchBattleData(info)
	if err != nil{
		t.Error(err)
	}
	err = DeleteMatchBattleData("9")
	if err != nil{
		t.Error(err)
	}
}

func TestDeleteMatchBattleData(t *testing.T) {

	info := dyn.MatchBattle{}
	info.RoleId = "22"
	info.HeroNo = 0
	info.Soldier = 0
	info.Register = 0
	info.Integral = 0
	err := InsertDateIntoMatchBattle(info)
	if err != nil {
		t.Error(err)
	}

	err = DeleteMatchBattleData("22")
	if err != nil{
		t.Error(err)
	}
}

func TestGetAllMatchBattleInfo(t *testing.T) {
	info, err := GetAllMatchBattleInfo()
	if err != nil{
		t.Error(err)
	}
	fmt.Println(info)
}