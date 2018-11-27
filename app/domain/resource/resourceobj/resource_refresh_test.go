package resourceobj

import (
	"fmt"
	"testing"
)

func TestResourceInsertResourceRefreshTime(t *testing.T) {
	resourceInfo := ResourceRefreshObj{}

	err := resourceInfo.InsertResourceRefreshTime("12",2,0)
	if err != nil{
		log.Error(err)
	}

	err = DeleteResourceRefreshTime("12")
	if err != nil{
		log.Error(err)
	}
}

func TestResourceRefreshObj_UpdateResourceRefreshTime(t *testing.T) {
	resourceInfo := ResourceRefreshObj{}
	err := resourceInfo.InsertResourceRefreshTime("12",2,0)
	if err != nil{
		log.Error(err)
	}

	err = resourceInfo.UpdateResourceRefreshTime("12",5,0)
	if err != nil{
		log.Error(err)
	}
	err = DeleteResourceRefreshTime("12")
	if err != nil{
		log.Error(err)
	}
}

func TestResourceRefreshObj_GetResourceRefreshTime(t *testing.T) {
	resourceInfo := ResourceRefreshObj{}
	err := resourceInfo.InsertResourceRefreshTime("12",5,0)
	if err != nil{
		log.Error(err)
	}
	resourceRefresh, err := resourceInfo.GetResourceRefreshTime("12",5)
	if err != nil {
		log.Error(err)
	}
	fmt.Println(resourceRefresh)
	err = DeleteResourceRefreshTime("12")
	if err != nil{
		log.Error(err)
	}
}

func TestIsIncludeResourceRefreshInfoRoleId(t *testing.T) {
	resourceInfo := ResourceRefreshObj{}

	err := resourceInfo.InsertResourceRefreshTime("12",2,0)
	if err != nil{
		log.Error(err)
	}
	flag := IsIncludeResourceRefreshInfoRoleId("12")
	if flag {
		log.Info(flag)
	}

	err = DeleteResourceRefreshTime("12")
	if err != nil{
		log.Error(err)
	}
}

func TestDeleteResourceRefreshTime(t *testing.T) {
	resourceInfo := ResourceRefreshObj{}

	err := resourceInfo.InsertResourceRefreshTime("12",2,0)
	if err != nil{
		log.Error(err)
	}
	err = DeleteResourceRefreshTime("12")
	if err != nil{
		log.Error(err)
	}
}
