package roleobj

import (
	"beginner-server/app"
	"beginner-server/app/model/dyn"
	"log"
)


type RoleObj struct {
	m 	*dyn.RoleInfoDyn

}


func (this *RoleObj)InsertRoleInfo( useId, name, passWord string)int{
	var roleInfo =  dyn.RoleInfoDyn{}

	roleInfo.RoleId = useId
	roleInfo.UserName = name
	roleInfo.Sex = 0
	roleInfo.Password = passWord
	err := app.DB.Table("role_info_dyns").Save(&roleInfo).Error
	if err != nil{
		log.Fatal("用户注册失败")
		return 4
	}
	return 0

}

func (this *RoleObj)GetRoleInfoByRoleId(roleId string)(dyn.RoleInfoDyn, error){
	RoleInfo := dyn.RoleInfoDyn{}
	err := app.DB.Where("role_id = ?", roleId).Find(&RoleInfo).Error
	return RoleInfo, err
}

func (this *RoleObj)GetAllRoleInfo()[]dyn.RoleInfoDyn{

	var roleInfos =  []dyn.RoleInfoDyn{}
	app.DB.Find(&roleInfos)
	return roleInfos
}

func (this *RoleObj)IsIncludeRoleInfoName(name string)( bool){
	include := app.DB.Table("role_info_dyns").Where("user_name = ?",name).First(&dyn.ResourceInfoDyn{}).RecordNotFound()
	return  include
}

func (this *RoleObj)IsIncludeRoleInfoRoleId(roleId string) bool {
	include := app.DB.Table("role_info_dyns").Where("role_id = ?",roleId).First(&dyn.ResourceInfoDyn{}).RecordNotFound()
	return include
}

func DeleteRoleInfoForTest(roleId string) error {
	return app.DB.Table("role_info_dyns").Where("role_id = ?",roleId).Delete(&dyn.RoleInfoDyn{}).Error
}