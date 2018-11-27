package actions

import(
	"beginner-server/app"
	"beginner-server/app/domain/role/rolesvc"
	"net/http"
)


//注册路由请求参数字段
type RegisterParam struct {
	UserId 		string		//用户ID，全服唯一
	Name   		string		//用户ID，全服唯一
	Password	string		//密码，不可直接将明文存储至数据库
}

//注册路由返回参数字段
type RegisterResp struct{
	Code 		app.CodeErrorType			//状态码， 0为正常，1. 解析数据出错  5.UserId已经存在， 6.Name已经存在 ，4.插入数据出错
}


//注册路由的Handler
//	注册的功能是，如果传入正确的用户名和密码则注册用户
//	若出错则返回相应的出错码
func RegisterHandler(w http.ResponseWriter, r *http.Request){
	param := RegisterParam{}
	resp := RegisterResp{
		Code: 0,
	}
	err := parseRequest(r, &param)
	if err != nil {
		resp.Code = app.ParseRequestError

	}
	resp.Code = RegisterRole(param.UserId, param.Name, param.Password)
	writeResponse(w, resp)
}

func RegisterRole(userId, name, password string)app.CodeErrorType	{
	return	rolesvc.Register(userId,name,password)

}
