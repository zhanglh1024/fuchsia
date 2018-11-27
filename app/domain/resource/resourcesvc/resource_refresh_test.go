package resourcesvc

import (
	"beginner-server/app/domain/resource/resourceobj"
	"fmt"
	"testing"
)

func TestGetResourceTime(t *testing.T) {
	resourceInfo := resourceobj.ResourceRefreshObj{}

	err := resourceInfo.InsertResourceRefreshTime("12",2,0)
	if err != nil{
		log.Error(err)
	}

	info,err := GetResourceTime("12", 2)
	if err != nil{
		t.Error(err)
	}
	fmt.Println(info)

	err = DeleteResourceTime("12")
	if err != nil{
		log.Error(err)
	}
}

func TestInsertResourceTime(t *testing.T) {
	err := InsertResourceTime("12",2,0)
	if err != nil{
		t.Error()
	}
	err = DeleteResourceTime("12")
	if err != nil{
		log.Error(err)
	}
}

func TestUpdateResourceRefresh(t *testing.T) {
	err := InsertResourceTime("12",2,0)
	if err != nil{
		t.Error()
	}
	err = UpdateResourceRefresh("12",2,1)
	if err != nil{
		t.Error(err)
	}
	err = DeleteResourceTime("12")
	if err != nil{
		log.Error(err)
	}
}

func TestIsExitRoleIdInTable(t *testing.T) {
	err := InsertResourceTime("12",2,0)
	if err != nil{
		t.Error()
	}
	exit := IsExitRoleIdInTable("12")
	if exit {
		log.Info("pass")
	}
	err = DeleteResourceTime("12")
	if err != nil{
		log.Error(err)
	}
}

func TestDeleteResourceTime(t *testing.T) {
	err := InsertResourceTime("12",2,0)
	if err != nil{
		t.Error()
	}
	err = DeleteResourceTime("12")
	if err != nil{
		log.Error(err)
	}
}
