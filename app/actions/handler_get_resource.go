package actions

import (
	"beginner-server/app"
	"beginner-server/app/domain/resource/resourcesvc"
	"errors"
	"github.com/cxr29/log"
	"net/http"
	"time"
)

//GetResource路由的请求参数字段
type GetResourceParam struct {
	UserId 		    string 			 //请求者的名字
	ResourceType    app.ResourceType //资源类型
}

//GetResource路由的响应字段
type GetResourceResp struct {
	Code     app.CodeErrorType   //返回码 1.解析数据出错 2.数据库处理出错 3.间隔时间少于60不能收集资源4.该玩家id不存在
	Resource []app.ResourceInfo //玩家当前拥有资源，数据结构与登录下发的相同
}

//GetResource路由的Handler
//	GetResource的功能是，如果传入正确的用户id和资源好，收集资源并返回用户资源信息
//	若出错则不返回资源信息，返回对应的错误码
func GetResourceHandler(w http.ResponseWriter, r *http.Request) {
	param := GetResourceParam{}
	resp := GetResourceResp{
		Code: 0,
		Resource: []app.ResourceInfo{},
	}
	err := parseRequest(r, &param)
	if err != nil {
		resp.Code = app.ParseRequestError

	}
	resp, err  = GetResourceOperate(param.UserId, param.ResourceType)
	if err != nil {
		log.Error(err)
	}
	writeResponse(w, resp)
}

func GetResourceOperate(roleId string, resourceType app.ResourceType)(resp GetResourceResp,err error){


	Exit := resourcesvc.IsExitRoleIdInTable(roleId)
	if Exit{
		resp.Code = app.RoleInfoIsNotExit
		err = errors.New("数据表中不存在该玩家信息")
		return
	}

	resourceTime,err := resourcesvc.GetResourceTime(roleId,resourceType)
	if err != nil{
		resp.Code = app.SearchDataError
		err = errors.New("查找收集资源时间出错")
		return
	}

	if ((time.Now().Unix() -resourceTime.RefreshTime.Unix()) < 60) && resourceTime.First != 0{
		resp.Code = app.IntervalLessThanOneMinute
		err = errors.New("间隔时间小于60秒不能收获")
		return
	}

	err = resourcesvc.UpdateResourceRefresh(roleId,resourceType,1)
	if err != nil{
		resp.Code = app.UpdateDataError
		err = errors.New("更新收集资源信息出错")
		return
	}

	err = resourcesvc.UpdateResourceByRoleId(roleId, resourceType)
	if err != nil {
		resp.Code = app.UpdateDataError
		err = errors.New("更新数据出错")
		return
	}

	resourceInfo, err := resourcesvc.GetResourceInfoByRoleId(roleId)
	if err != nil {
		resp.Code = app.SearchDataError
		err = errors.New("查找资源信息出错")
		return
	}

	resp.Resource = resourceInfo
	return
}

