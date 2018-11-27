package resourcesvc

import (
	"beginner-server/app"
	"beginner-server/app/domain/resource/resourceobj"
	"beginner-server/app/model/dyn"
)


//获取玩家收集资源时间
func GetResourceTime(roleId string, resourceType app.ResourceType) (dyn.ResourceRefresh, error) {
	resourceRefresh := resourceobj.ResourceRefreshObj{}
	resourceTime, err := resourceRefresh.GetResourceRefreshTime(roleId, resourceType)
	return resourceTime, err
}

//插入玩家收集资源时间
func InsertResourceTime(roleId string, resourceType app.ResourceType, first int) error {
	resourceRefresh := resourceobj.ResourceRefreshObj{}
	return resourceRefresh.InsertResourceRefreshTime(roleId, resourceType, first)
}

//刷新玩家收集资源时间
func UpdateResourceRefresh(roleId string,resourceType app.ResourceType, first int)error  {
	resourceRefresh := resourceobj.ResourceRefreshObj{}
	return resourceRefresh.UpdateResourceRefreshTime(roleId, resourceType, first)
}

func DeleteResourceTime(roleId string) error {
	return resourceobj.DeleteResourceRefreshTime(roleId)
}


//数据表中是否存在改玩家信息
func IsExitRoleIdInTable(roleId string) bool {
	return resourceobj.IsIncludeResourceRefreshInfoRoleId(roleId)
}
