package heroobj

import (
	"beginner-server/app/model/dyn"
	"fmt"
	"testing"
)

func TestInsertHeroValue(t *testing.T){
	hero := dyn.HeroInfoDyn{}
	hero.No = 4
	hero.RoleId = "3"
	hero.Lv = 2
	hero.Heroic = 12
	hero.Diligent = 8
	hero.Wise = 9
	hero.Loyal = 7

	err := InsertHeroValue(hero)
	if err != nil{
		fmt.Println("存入数据失败:", err)
	}
	fmt.Println("存入成功")

	err = DeleteHeroInfo("3",4)
	if err != nil{
		t.Error(err)
	}
}

func TestGetHeroInfoByRoleId(t *testing.T) {
	hero := dyn.HeroInfoDyn{}
	hero.No = 4
	hero.RoleId = "3"
	hero.Lv = 2
	hero.Heroic = 12
	hero.Diligent = 8
	hero.Wise = 9
	hero.Loyal = 7

	err := InsertHeroValue(hero)
	if err != nil{
		fmt.Println("存入数据失败:", err)
	}
	heroInfo, err := GetHeroInfoByRoleId("3")
	if err != nil {
		log.Error(err)
	}
	if len(heroInfo) == 0 {
		fmt.Println("获取到一个空值")
	}
	fmt.Println(heroInfo)
	err = DeleteHeroInfo("3",4)
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

	err := InsertHeroValue(hero)
	if err != nil{
		fmt.Println("存入数据失败:", err)
	}

	info, err:=GetHeroInfoByRoleIdAndHeroNo("3",4)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(info)

	err = DeleteHeroInfo("3",4)
	if err != nil{
		t.Error(err)
	}
}

func TestUpdateHeroInfo(t *testing.T) {

	hero := dyn.HeroInfoDyn{}
	hero.No = 4
	hero.RoleId = "3"
	hero.Lv = 2
	hero.Heroic = 12
	hero.Diligent = 8
	hero.Wise = 9
	hero.Loyal = 7

	err := InsertHeroValue(hero)
	if err != nil{
		fmt.Println("存入数据失败:", err)
	}

	hero.RoleId = "3"
	hero.No = 4
	hero.Lv = 4
	hero.Heroic = 14

	err = UpdateHeroInfo(hero)
	if err != nil{
		t.Error(err)
	}

	err = DeleteHeroInfo("3",4)
	if err != nil{
		t.Error(err)
	}
}

func TestDeleteHeroInfo(t *testing.T) {
	hero := dyn.HeroInfoDyn{}
	hero.No = 4
	hero.RoleId = "3"
	hero.Lv = 2
	hero.Heroic = 12
	hero.Diligent = 8
	hero.Wise = 9
	hero.Loyal = 7

	err := InsertHeroValue(hero)
	if err != nil{
		fmt.Println("存入数据失败:", err)
	}
	fmt.Println("存入成功")

	err = DeleteHeroInfo("3",4)
	if err != nil{
		t.Error(err)
	}
}