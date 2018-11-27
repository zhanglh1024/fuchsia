package resourceobj

import (
	"beginner-server/app"
	"beginner-server/app/model/dyn"
)


func InsertResource(dyn dyn.ResourceInfoDyn)error{
	err := app.DB.Save(&dyn).Error
	return err
}

func GetResourceInfoByRoleId(roleId string)(dyn.ResourceInfoDyn, error){
	resourceInfo := dyn.ResourceInfoDyn{}
	err := app.DB.Where("role_id = ?", roleId).Find(&resourceInfo).Error
	return resourceInfo, err
}


func UpdateResource(roleId string, resourceInfo dyn.ResourceInfoDyn)error{
	err := app.DB.Model(&resourceInfo).Where("role_id = ?", roleId).UpdateColumns(&resourceInfo).Error
	return err
}

func UpdateSoldierNum(roleId string, resource dyn.ResourceInfoDyn) error {
	return app.DB.Table("resource_info_dyns").Where("role_id = ?", roleId).Update("soldier",resource.Soldier).Error
}

func DeleteResource(roleId string) error {
	return app.DB.Table("resource_info_dyns").Where("role_id = ?", roleId).Delete(&dyn.ResourceInfoDyn{}).Error
}


