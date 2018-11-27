package resourceobj

import (
	"beginner-server/app/model/dyn"
	"fmt"
	"testing"
)

func TestInsertResource(t *testing.T) {
	resourceInfo:= dyn.ResourceInfoDyn{
			RoleId:"3",
			Food:12,
			Soldier:33,
			Coin :1000,
	}
	err := InsertResource(resourceInfo)
	if err != nil{
		log.Error(err)
	}

	err = DeleteResource("3")
	if err != nil{
		log.Error(err)
	}
}

func TestGetResourceInfoByRoleId(t *testing.T) {
	resourceInfo := dyn.ResourceInfoDyn{}
	resourceInfo.RoleId = "8"
	resourceInfo.Food = 1222
	resourceInfo.Soldier = 333
	resourceInfo.Coin = 1000

	err := InsertResource(resourceInfo)
	if err != nil{
		t.Error(err)
	}

	info, err := GetResourceInfoByRoleId("8")
	if err != nil{
		t.Error(err)
	}

	fmt.Println(info)

	err = DeleteResource("8")
	if err != nil{
		t.Error(err)
	}
}

func TestUpdateSoldierNum(t *testing.T) {
	resourceInfo := dyn.ResourceInfoDyn{}
	resourceInfo.RoleId = "8"
	resourceInfo.Food = 1222
	resourceInfo.Soldier = 333
	resourceInfo.Coin = 1000

	err := InsertResource(resourceInfo)
	if err != nil{
		t.Error(err)
	}

	resourceInfo.Soldier = 121
	err = UpdateSoldierNum("8",resourceInfo)
	if err != nil{
		t.Error(err)
	}

	err = DeleteResource("8")
	if err != nil{
		t.Error(err)
	}
}



func TestUpdateResource(t *testing.T) {
	resourceInfo := dyn.ResourceInfoDyn{}
	resourceInfo.RoleId = "8"
	resourceInfo.Food = 1222
	resourceInfo.Soldier = 333
	resourceInfo.Coin = 1000

	err := InsertResource(resourceInfo)
	if err != nil{
		t.Error(err)
	}

	resourceInfo.Food = 100

	err = UpdateResource("8",resourceInfo)
	if err != nil {
		log.Error(err)
	}

	err = DeleteResource("8")
	if err != nil{
		t.Error(err)
	}
}

func TestDeleteResource(t *testing.T) {
	resourceInfo:= dyn.ResourceInfoDyn{
		RoleId:"3",
		Food:12,
		Soldier:33,
		Coin :1000,
	}
	err := InsertResource(resourceInfo)
	if err != nil{
		t.Error(err)
	}
	err = DeleteResource("3")
	if err != nil{
		t.Error(err)
	}
}
