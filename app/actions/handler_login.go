package actions

import (
	"github.com/cxr29/log"
	"beginner-server/app"
	"beginner-server/app/domain/hero/herosvc"
	"beginner-server/app/domain/resource/resourcesvc"
	"beginner-server/app/domain/role/rolesvc"
	"errors"
	"net/http"
)

//登录路由请求参数字段
type LoginParam struct {
	UserId 		string		//用户ID，全服唯一
	Password	string		//密码，不可直接将明文存储至数据库
}

//登录路由返回参数字段
type LoginResp struct{
	Code 		app.CodeErrorType						   //状态码
	Resource    []app.ResourceInfo		    //资源数据列表
	Hero		[]app.HeroInfo   	    	//英雄数据列表
}



//login路由的Handler
//	login的功能是，如果传入正确的用户名和密码，则返回资源的用户资源信息和名下英雄信息
//	若出错则返回错误码，资源和用户英雄信息为空
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	param := LoginParam{}
	resp := LoginResp{
		Code: 0,
		Resource: []app.ResourceInfo{},
		Hero: []app.HeroInfo{},
	}
	err := parseRequest(r, &param)
	if err != nil {
		resp.Code = app.ParseRequestError
	}
	resp, err = LoginOperate(param.UserId, param.Password)
	if err != nil{
		log.Error(err)
	}
	writeResponse(w, resp)
}


func LoginOperate(roleId, passWord string)(loginResp LoginResp, err error){
	if roleId == "" || passWord == ""{
		loginResp.Code = app.RoleIdOrPassWordNil
		err = errors.New("账号或者密码为空")
		return
	}

	exit := rolesvc.IsExitRoleInfo(roleId)
	if exit{
		loginResp.Code = app.RoleInfoIsNotExit
		err = errors.New("该账号不存在")
		return
	}

	correct, err := rolesvc.LoginVerification(roleId, passWord)
	if err != nil || correct==false{
		loginResp.Code = app.RoleIdOrPassWordError
		err = errors.New("用户密码错误")
		return
	}

	heroInfo, err := herosvc.GetLoginHeroInfo(roleId)
	if err != nil {
		loginResp.Code = app.SearchDataError
		err = errors.New("查找用户英雄信息出错")
		return
	}
	loginResp.Hero = heroInfo

	resourceInfo, err := resourcesvc.GetResourceInfoByRoleId(roleId)
	if err != nil {
		loginResp.Code = app.SearchDataError
		err = errors.New("查找用户资源信息出错")
		return
	}
	loginResp.Resource = resourceInfo
	loginResp.Code = 0

	return 
}


