package herosvc

import (
	"beginner-server/app/domain/hero/heroobj"
	"beginner-server/app/model/dyn"
	"fmt"
	"testing"
)

func TestGetLoginHeroInfo(t *testing.T) {

	hero := dyn.HeroInfoDyn{}
	hero.No = 4
	hero.RoleId = "3"
	hero.Lv = 2
	hero.Heroic = 12
	hero.Diligent = 8
	hero.Wise = 9
	hero.Loyal = 7

	err := heroobj.InsertHeroValue(hero)
	if err != nil{
		fmt.Println("存入数据失败:", err)
	}
	fmt.Println("存入成功")

	heroInfo, err := GetLoginHeroInfo("3")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(heroInfo)

	err = heroobj.DeleteHeroInfo("3",4)
	if err != nil{
		t.Error(err)
	}
}

func TestGetHeroInfoByRoleIdAndHeroNo(t *testing.T) {

	hero := dyn.HeroInfoDyn{}
	hero.No = 4
	hero.RoleId = "3"
	hero.Lv = 2
	hero.Heroic = 12
	hero.Diligent = 8
	hero.Wise = 9
	hero.Loyal = 7

	err := heroobj.InsertHeroValue(hero)
	if err != nil{
		fmt.Println("存入数据失败:", err)
	}
	fmt.Println("存入成功")

	info, err := GetHeroInfoByRoleIdAndHeroNo("3", 4)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(info)

	err = heroobj.DeleteHeroInfo("3",4)
	if err != nil{
		t.Error(err)
	}
}

func TestUpgradeHeroLv(t *testing.T) {

	hero := dyn.HeroInfoDyn{}
	hero.No = 4
	hero.RoleId = "3"
	hero.Lv = 2
	hero.Heroic = 12
	hero.Diligent = 8
	hero.Wise = 9
	hero.Loyal = 7

	err := heroobj.InsertHeroValue(hero)
	if err != nil{
		fmt.Println("存入数据失败:", err)
	}
	fmt.Println("存入成功")

	heroInfo := dyn.HeroInfoDyn{}
	heroInfo.Lv = 3
	heroInfo.No = 4
	heroInfo.RoleId = "3"
	err = UpgradeHeroLv(heroInfo)
	if err != nil {
		t.Error(err)
	}

	err = heroobj.DeleteHeroInfo("3",4)
	if err != nil{
		t.Error(err)
	}
}

func TestGetHeroFightByRoleIdAndHeroNo(t *testing.T) {
	hero := dyn.HeroInfoDyn{}
	hero.No = 4
	hero.RoleId = "3"
	hero.Lv = 2
	hero.Heroic = 12
	hero.Diligent = 8
	hero.Wise = 9
	hero.Loyal = 7

	err := heroobj.InsertHeroValue(hero)
	if err != nil{
		fmt.Println("存入数据失败:", err)
	}
	fmt.Println("存入成功")

	type Input struct{
		herNo int
		roleId string
	}

	type DataTest struct{
		input  Input
		expect int
	}


    dataTest := DataTest{
    	input: struct {
			herNo  int
			roleId string
		}{herNo: 4, roleId:"3" },
    	expect:36,
	}

	result,err := GetHeroFightByRoleIdAndHeroNo(dataTest.input.roleId,dataTest.input.herNo)
	if err != nil{
		t.Error(err)
	}
	if result != dataTest.expect{
		t.Errorf("GetHeroFightByRoleIdAndHeroNo error = %v, wantErr %v", err, dataTest.expect)
	}


	err = heroobj.DeleteHeroInfo("3",4)
	if err != nil{
		t.Error(err)
	}
}

func TestCalculateHeroCombatPower(t *testing.T) {

	tests := []struct {
		name         string
		soldierNum   int
		heroFight    int
		Expect       int
	}{
		// TODO: Add test cases.
		{
			name:    "CalculateHeroCombatPower",
			soldierNum:   3,
			heroFight :   4,
			Expect :      23,
		},
		{
			name:    "CalculateHeroCombatPower",
			soldierNum:   8,
			heroFight :   4,
			Expect :      28,
		},
	}


    for _,tt := range tests{
    	t.Run(tt.name ,func(t *testing.T) {
    		result := CalculateHeroCombatPower(tt.soldierNum, tt.heroFight)
			if result != tt.Expect {
				t.Errorf("CalculateHeroCombatPower result = %v, Expect %v", result, tt.Expect)
				return
			}
		})
	}

}

func TestConsumingSoldiers(t *testing.T) {
	type Expect struct {
		win int
		los int
	}

	tests := []struct {
		name         string
		soldierNum   int
		heroFight    int
		roundId		 int
		losSoldier	 int
		losHeroFight int
		expect       Expect
	}{
		// TODO: Add test cases.
		{
			name:    "ConsumingSoldiers",
			soldierNum:   3,
			heroFight :   4,
			roundId:      1,
			losSoldier:   5,
			losHeroFight: 8,
			expect :      Expect{win:8,los:10},
		},
		{
			name:    "ConsumingSoldiers",
			soldierNum:   8,
			heroFight :   4,
			roundId:      1,
			losSoldier:   5,
			losHeroFight: 8,
			expect :      Expect{win:15,los:10},
		},
	}

	for _,tt := range tests{
		t.Run(tt.name ,func(t *testing.T) {
			winCost, loseCost := ConsumingSoldiers(tt.soldierNum, tt.heroFight, tt.roundId, tt.losSoldier, tt.losHeroFight)
			if winCost != tt.expect.win || loseCost != tt.expect.los {
				t.Errorf("CalculateHeroCombatPower winCost = %v, winCost = %v, loseCost %v", winCost,loseCost, tt.expect)
				return
			}
		})
	}
}