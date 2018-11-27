package battlesvc

import (
	"beginner-server/app/domain/battle/battleobj"
	"beginner-server/app/model/dyn"
	"fmt"
	"testing"
)

func TestEditCurrentLevel(t *testing.T) {
	info := dyn.LevelBattle{}
	info.RoleId = "33"
	info.RoundId = 1
	info.HeroFight = 15
	info.SoldierLeft = 80
	err := battleobj.InsertLevelBattleValue(info)
	if err != nil{
		t.Error(err)
	}


	err = EditCurrentLevel(2,info)
	if err != nil{
		t.Error(err)
	}


	err = battleobj.DeleteLevelBattleValue("33")
	if err != nil{
		t.Error(err)
	}
}

func TestGetCurrentLevelInfo(t *testing.T) {
	info := dyn.LevelBattle{}
	info.RoleId = "33"
	info.RoundId = 1
	info.HeroFight = 15
	info.SoldierLeft = 80
	err := battleobj.InsertLevelBattleValue(info)
	if err != nil{
		t.Error(err)
	}

	infos, err := GetCurrentLevelInfo("23")
	if err != nil{
		t.Error(err)
	}
	fmt.Println(infos)

	err = battleobj.DeleteLevelBattleValue("33")
	if err != nil{
		t.Error(err)
	}
}

func TestPassEditLevelInfo(t *testing.T) {

	infos := []dyn.LevelBattle{
		{
			RoleId: "33",
			RoundId: 1,
			HeroFight: 15,
			SoldierLeft:  80,
		},
		{
			RoleId: "34",
			RoundId: 5,
			HeroFight: 15,
			SoldierLeft:  80,
		},
	}

	for _,info := range infos{
		err := battleobj.InsertLevelBattleValue(info)
		if err != nil{
			t.Error(err)
		}


		err, code := PassEditLevelInfo(info)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(code)

		err = battleobj.DeleteLevelBattleValue(info.RoleId)
		if err != nil{
			t.Error(err)
		}

	}

}

