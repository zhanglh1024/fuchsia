package resourcesvc

import (
	"github.com/cxr29/log"
	"beginner-server/app"
	"beginner-server/app/domain/hero/heroobj"
	"beginner-server/app/domain/resource/resourceobj"
	"beginner-server/app/model/dyn"
)


//根据用户Id查找用户所有的资源信息
func GetResourceInfoByRoleId(roleId string)  ([]app.ResourceInfo, error){

	resouInfos := make([]app.ResourceInfo, 3)

	resourceInfo, err := resourceobj.GetResourceInfoByRoleId(roleId)
	if err != nil {
		return resouInfos, err
	}

	for k,value := range app.ResourceTypeList{
		resouInfos[k].Type = value
		switch value {
		case app.FOOD:
			resouInfos[k].Count = int(resourceInfo.Food)
		case app.SOLDIER:
			resouInfos[k].Count = int(resourceInfo.Soldier)
		default:
			resouInfos[k].Count = int(resourceInfo.Coin)
		}

	}

	return resouInfos, nil

}

//根据用户id来获取用户资源
func GetRoleResourceInfoByRoleId(roleId string)(dyn.ResourceInfoDyn, error) {
	info, err := resourceobj.GetResourceInfoByRoleId(roleId)
	return info, err
}

//更新用户资源信息，收集资源handler调用
func UpdateResourceByRoleId(roleId string, resourceType app.ResourceType) error  {
	heros ,err := heroobj.GetHeroInfoByRoleId(roleId)
	if err != nil {
		log.Errorf("获取英雄数据出错:%s", err)
		return err
	}
	wise,loyal,diligent := 0, 0, 0
	for _, hero := range heros{
		wise += hero.Wise
		loyal += hero.Loyal
		diligent += hero.Diligent
	}

	resourceInfo, err := resourceobj.GetResourceInfoByRoleId(roleId)
	if err != nil {
		log.Errorf("获取资源数据出错：%s", err)
		return err
	}

	switch resourceType {
	case app.FOOD:
		resourceInfo.Food = int64(((diligent-1)/3)*10 + 100) + resourceInfo.Food
	case app.SOLDIER:
		resourceInfo.Soldier = int64(((loyal-1)/3)*10 + 100) + resourceInfo.Soldier
	case app.GOLD:
		resourceInfo.Coin = int64(((wise-1)/3)*10 + 100) + resourceInfo.Coin
	}

	err= resourceobj.UpdateResource(roleId,resourceInfo)
	return err

}

//更新资源信息不设置成默认值的借口
func UpgradeHeroLvCostResource(roleId string, resourceInfo dyn.ResourceInfoDyn) (error) {
	return resourceobj.UpdateResource(roleId, resourceInfo)
}

//func UpdateSoldierForFight(cost int, resourceDyn dyn.ResourceInfoDyn) error {
//	if cost > int(resourceDyn.Soldier){
//		resourceDyn.Soldier = 0
//	}else{
//		resourceDyn.Soldier -= int64(cost)
//	}
//	return resourceobj.UpdateSoldierNum(resourceDyn.RoleId, resourceDyn)
//}

//更新资源信息需要设置成默认值0调用这个更新接口
func UpdateMatchBattleData(cost int,RoleId string, soldierCount int) error {
	info := dyn.ResourceInfoDyn{
		RoleId:RoleId,
	}
	if cost >= soldierCount{
		soldierCount = 0
	}else{
		soldierCount -= cost
	}
	info.Soldier = int64(soldierCount)
	log.Info(info)
	return resourceobj.UpdateSoldierNum(RoleId, info)
}






