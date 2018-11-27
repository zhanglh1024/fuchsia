package heroobj

import (
	"beginner-server/app"
	"beginner-server/app/model/dyn"
)




func InsertHeroValue(hero dyn.HeroInfoDyn)error{
	err := app.DB.Save(&hero).Error
	return err
}


func GetHeroInfoByRoleId(roleId string)([]dyn.HeroInfoDyn, error){
	HeroInfo := []dyn.HeroInfoDyn{}
	err := app.DB.Where("role_id = ?", roleId).Find(&HeroInfo).Error
	return HeroInfo, err
}

func GetHeroInfoByRoleIdAndHeroNo(roleId string, heroNo int) (dyn.HeroInfoDyn, error) {
	info := dyn.HeroInfoDyn{}
	err := app.DB.Table("hero_info_dyns").Where("no = ? and role_id = ?",heroNo,roleId).Find(&info).Error
	return info ,err

}

func UpdateHeroInfo(heroInfo dyn.HeroInfoDyn) error {
	return app.DB.Table("hero_info_dyns").Where("no = ? and role_id = ?", heroInfo.No, heroInfo.RoleId).Update(&heroInfo).Error
}

func DeleteHeroInfo(roleId string, heroNo int) error {
	return app.DB.Table("hero_info_dyns").Where("no = ? and role_id = ?", heroNo,roleId).Delete(&dyn.HeroInfoDyn{}).Error
}