package resourcesvc

import (
	"beginner-server/app"
	"beginner-server/app/domain/resource/resourceobj"
	"beginner-server/app/model/dyn"
	"fmt"
	"testing"
)

func TestGetResourceInfoByRoleId(t *testing.T) {

	resourceInfo := dyn.ResourceInfoDyn{}
	resourceInfo.RoleId = "8"
	resourceInfo.Food = 1222
	resourceInfo.Soldier = 333
	resourceInfo.Coin = 1000
	err := resourceobj.InsertResource(resourceInfo)
	if err != nil{
		t.Error(err)
	}

	info, err := GetResourceInfoByRoleId("8")
	if err!= nil{
		t.Errorf("get resource error:%s", err)
	}
	fmt.Println(info)

	err = resourceobj.DeleteResource("8")
	if err!= nil{
		t.Errorf("get resource error:%s", err)
	}
}

func TestGetRoleResourceInfoByRoleId(t *testing.T) {
	resourceInfo := dyn.ResourceInfoDyn{}
	resourceInfo.RoleId = "8"
	resourceInfo.Food = 1222
	resourceInfo.Soldier = 333
	resourceInfo.Coin = 1000
	err := resourceobj.InsertResource(resourceInfo)
	if err != nil{
		t.Error(err)
	}
	info, err := GetRoleResourceInfoByRoleId("8")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(info)
	err = resourceobj.DeleteResource("8")
	if err!= nil{
		t.Errorf("get resource error:%s", err)
	}
}

func TestUpdateResourceByRoleId(t *testing.T)  {

	resourceInfo := dyn.ResourceInfoDyn{}
	resourceInfo.RoleId = "8"
	resourceInfo.Food = 1222
	resourceInfo.Soldier = 333
	resourceInfo.Coin = 1000
	err := resourceobj.InsertResource(resourceInfo)
	if err != nil{
		t.Error(err)
	}

	type input struct{
		roleId string
		resourceType app.ResourceType
	}

	inputs := []input{
		{
			roleId:"22",
			resourceType:app.FOOD,
		},
		{
			roleId:"22",
			resourceType:app.GOLD,
		},
		{
			roleId:"22",
			resourceType:app.SOLDIER,
		},
	}

	for _,value := range inputs{
		err := UpdateResourceByRoleId(value.roleId,value.resourceType)
		if err != nil{
			t.Errorf("更新资源出错:%s", err)
		}
	}

	err = resourceobj.DeleteResource("8")
	if err!= nil{
		t.Errorf("get resource error:%s", err)
	}
}

func TestUpgradeHeroLvCostResource(t *testing.T) {
	resourceInfo := dyn.ResourceInfoDyn{}
	resourceInfo.Food = 1000
	resourceInfo.Soldier = 1000
	err:=UpgradeHeroLvCostResource("123",resourceInfo)
	if err != nil{
		t.Errorf("更新资源出错:%s", err)
	}
}

func TestUpdateMatchBattleData(t *testing.T) {
	err := UpdateMatchBattleData(90,"8",67)
	if err != nil{
		t.Error(err)
	}
}
