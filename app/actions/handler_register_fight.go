package actions

import (
	"github.com/cxr29/log"
	"beginner-server/app"
	"beginner-server/app/domain/battle/battlesvc"
	"beginner-server/app/domain/resource/resourcesvc"
	"errors"
	"net/http"
)

//RegisterFight路由的请求参数字段
type RegisterFightParam struct {
	UserId  string //用户ID
	HeroNo  int    //出战英雄编号
	Soldier int    //出站士兵数
}

//RegisterFight路由的响应字段
type RegisterFightResp struct {
	Code     app.CodeErrorType	                //返回码 1解析出错，2数据库处理出错，3用户已经注册过了
	Resource []app.ResourceInfo //玩家当前拥有资源
}

//RegisterFight路由的Handler
//	RegisterFight的功能是，如果传入的MyName不为空字符传，则返回<问候语>+<请求者名字>，
//	若MyName为空字符串，则返回<问候语>+"Stranger"
func RegisterFightHandler(w http.ResponseWriter, r *http.Request) {
	param := RegisterFightParam{}
	resp := RegisterFightResp{
		Code: 0,
		Resource: []app.ResourceInfo{},
	}
	err := parseRequest(r, &param)
	if err != nil {
		resp.Code = app.ParseRequestError

	}
	resp,err = RegisterFightOperate(param.UserId,param.HeroNo,param.Soldier)
	if err != nil{
		log.Error(err)
	}
	writeResponse(w, resp)
}

func RegisterFightOperate(roleId string, heroNo, soldierNum int) (resp RegisterFightResp,err error) {
	matchInfo, err := battlesvc.GetMatchBattleInfo(roleId)
	if err != nil{
		resp.Code = app.SearchDataError
		err = errors.New("查询数据失败")
		return
	}

	if matchInfo.Register == 1 {
		resp.Code = app.RoleIdAlreadyRegister
		err = errors.New("用户已经注册登记大厅，不要重复注册")
		return
	}

	resourceInfo,err := resourcesvc.GetRoleResourceInfoByRoleId(roleId)
	if err != nil{
		resp.Code = app.SearchDataError
		err = errors.New("获取用户资源失败")
		return
	}

	if soldierNum > int(resourceInfo.Soldier){
		resp.Code = app.SendTooMuchSoldier
		err = errors.New("派出士兵数量过多")
		return
	}


	err = resourcesvc.UpdateMatchBattleData(soldierNum, resourceInfo.RoleId, int(resourceInfo.Soldier))
	if err != nil{
		resp.Code = app.UpdateDataError
		err = errors.New("更新用户资源失败")
		return
	}

	err = battlesvc.RegisterMatchBattleHall(roleId, heroNo, soldierNum )
	if err != nil{
		resp.Code = app.UpdateDataError
		err = errors.New("注册失败")
		return
	}

	resp.Resource,err = resourcesvc.GetResourceInfoByRoleId(roleId)
	if err != nil{
		resp.Code = app.SearchDataError
		err = errors.New("获取用户资源失败")
		return
	}

	resp.Code = 0

	return
}
