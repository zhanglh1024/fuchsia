package battlesvc

import (
	"beginner-server/app/model/dyn"
	"fmt"
	"testing"
)

func TestGetMatchBattleInfo(t *testing.T) {

	info := dyn.MatchBattle{}
	info.RoleId = "2"
	info.HeroNo = 0
	info.Soldier = 0
	info.Register = 0
	info.Integral = 0
	err := InsertForTest(info)
	if err != nil {
		t.Error(err)
	}
	infos, err := GetMatchBattleInfo("2")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(infos)

	err = DeleteForTest("2")
	if err != nil{
		t.Error(err)
	}
}

func TestRegisterMatchBattleHall(t *testing.T) {

	info := dyn.MatchBattle{}
	info.RoleId = "2"
	info.HeroNo = 0
	info.Soldier = 0
	info.Register = 0
	info.Integral = 0
	err := InsertForTest(info)
	if err != nil {
		t.Error(err)
	}

	err = RegisterMatchBattleHall("2",1,200)
	if err != nil{
		t.Error(err)
	}

	err = DeleteForTest("2")
	if err != nil{
		t.Error(err)
	}
}

func TestGetRandMatchBattleOpponent(t *testing.T) {
	info, err := GetRandMatchBattleOpponent()
	if err != nil{
		t.Error(err)
	}
	fmt.Println(info)
}

func TestUpdateMatchBattleInfo(t *testing.T) {

	info := dyn.MatchBattle{}
	info.RoleId = "2"
	info.HeroNo = 0
	info.Soldier = 0
	info.Register = 0
	info.Integral = 0
	err := InsertForTest(info)
	if err != nil {
		t.Error(err)
	}

	info.Soldier = 100
	info.Integral=10

	err = UpdateMatchBattleInfo(info)

	err = DeleteForTest("2")
	if err != nil{
		t.Error(err)
	}
}

func TestInsertForTest(t *testing.T) {
	info := dyn.MatchBattle{}
	info.RoleId = "2"
	info.HeroNo = 0
	info.Soldier = 0
	info.Register = 0
	info.Integral = 0
	err := InsertForTest(info)
	if err != nil {
		t.Error(err)
	}

	err = DeleteForTest("2")
	if err != nil{
		t.Error(err)
	}
}

func TestDeleteForTest(t *testing.T) {
	info := dyn.MatchBattle{}
	info.RoleId = "2"
	info.HeroNo = 0
	info.Soldier = 0
	info.Register = 0
	info.Integral = 0
	err := InsertForTest(info)
	if err != nil {
		t.Error(err)
	}

	err = DeleteForTest("2")
	if err != nil{
		t.Error(err)
	}
}
