package resourceobj

import (
	"beginner-server/app"
	"beginner-server/app/model/dyn"
	"time"
)

type ResourceRefreshObj struct {
}

func (this *ResourceRefreshObj)InsertResourceRefreshTime( roleId string, resourceType app.ResourceType, first int) error {
	resourceTime := dyn.ResourceRefresh{}
	resourceTime.RoleId = roleId
	resourceTime.TypeId = int16(resourceType)
	resourceTime.First = first
	resourceTime.RefreshTime = time.Now()
	return  app.DB.Save(&resourceTime).Error

}

func (this *ResourceRefreshObj)UpdateResourceRefreshTime( roleId string, resourceType app.ResourceType, first int)error{
	resourceTime := dyn.ResourceRefresh{}
	resourceTime.RoleId = roleId
	resourceTime.TypeId = int16(resourceType)
	resourceTime.First = first
	resourceTime.RefreshTime = time.Now()
	err:= app.DB.Table("resource_refreshes").Where("role_id = ? and type_id = ?",roleId, resourceType).Update(&resourceTime).Error
	return err
}

func (this *ResourceRefreshObj)GetResourceRefreshTime(roleId string, resourceType app.ResourceType) (dyn.ResourceRefresh, error) {
	resourceTime := dyn.ResourceRefresh{}
	err := app.DB.Table("resource_refreshes").Select("*").Where("role_id = ? and type_id = ?",roleId, resourceType).Find(&resourceTime).Error
	return resourceTime, err
}


func IsIncludeResourceRefreshInfoRoleId(roleId string) bool {
	include := app.DB.Table("resource_refreshes").Where("role_id = ?",roleId).First(&dyn.ResourceRefresh{}).RecordNotFound()
	return include
}

func DeleteResourceRefreshTime(roleId string) error {
	return app.DB.Table("resource_refreshes").Where("role_id = ?", roleId).Delete(&dyn.ResourceRefresh{}).Error
}