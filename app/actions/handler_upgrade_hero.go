package actions

import (
	"github.com/cxr29/log"
	"beginner-server/app"
	"beginner-server/app/domain/hero/herosvc"
	"beginner-server/app/domain/resource/resourcesvc"
	"errors"
	"fmt"
	"net/http"
)

//UpdateHero路由的请求参数字段
type UpdateHeroParam struct {
	UserId string //请求者的ID
	HeroNo int    //英雄编号
}

//UpdateHero路由的响应字段
type UpdateHeroResp struct {
	Code       app.CodeErrorType	    //返回码
	Lv 		   int    //英雄当前等级
	Resource   []app.ResourceInfo//玩家当前拥有资源
}

//UpdateHero路由的Handler
//	UpdateHero的功能是，如果传入正确的用户名和英雄编号则返回升级后登记和玩家拥有的资源信息
//	若出错则返回当前英雄等级，不返回玩家资源信息
func UpdateHeroHandler(w http.ResponseWriter, r *http.Request) {
	param := UpdateHeroParam{}
	resp := UpdateHeroResp{
		Code: 0,
		Lv:   0,
		Resource: []app.ResourceInfo{},
	}
	err := parseRequest(r, &param)
	if err != nil {
		resp.Code = app.ParseRequestError

	}

	resp, err = UpgradeHeroLvOperate(param.UserId, param.HeroNo)
	if err != nil{
		log.Error(err)
	}

	writeResponse(w, resp)
}

func UpgradeHeroLvOperate(roleId string, heroNo int) (resp UpdateHeroResp,err error){
	heroInfo, err := herosvc.GetHeroInfoByRoleIdAndHeroNo(roleId,heroNo)
	if err != nil {
		resp.Code = app.SearchDataError
		resp.Lv = int(heroInfo.Lv)
		err = errors.New("获取英雄信息出错")
		return
	}

	resourceInfo, err := resourcesvc.GetRoleResourceInfoByRoleId(roleId)
	if err != nil {
		resp.Code = app.SearchDataError
		resp.Lv = int(heroInfo.Lv)
		err = errors.New("获取资源失败")
		return
	}

	resourceInfos, err := resourcesvc.GetResourceInfoByRoleId(roleId)
	if err != nil {
		resp.Code = app.SearchDataError
		resp.Lv = int(heroInfo.Lv)
		err = errors.New("获取资源失败")
		return
	}

	//判断是否满足升级条件
	if heroInfo.Lv >= 100 {
		resp.Code = app.UserAlreadyFull
		resp.Lv = int(heroInfo.Lv)
		resp.Resource = resourceInfos
		err = errors.New("用户已经满级了")
		return
	}
	if int(resourceInfo.Food)<int(heroInfo.Lv)*10 || int(resourceInfo.Coin)<int(heroInfo.Lv)*10{
		resp.Code = app.NotEnoughResources
		resp.Lv = int(heroInfo.Lv)
		resp.Resource = resourceInfos
		err = errors.New("用户资源不够")
		return
	}


	//执行资源更新操作
	resourceInfo.Food = resourceInfo.Food - int64(heroInfo.Lv *10)
	resourceInfo.Coin = resourceInfo.Coin - int64(heroInfo.Lv*10)
	fmt.Println(roleId, "-------",resourceInfo)
	err =resourcesvc.UpgradeHeroLvCostResource(roleId, resourceInfo)
	if err != nil {
		resp.Code = app.UpdateDataError
		resp.Lv = int(heroInfo.Lv)
		resp.Resource = resourceInfos
		err = errors.New("更新用户资源出错")
		return
	}

	resourceInfos, err = resourcesvc.GetResourceInfoByRoleId(roleId)
	if err != nil {
		resp.Code = app.UpdateDataError
		resp.Lv = int(heroInfo.Lv)
		err = errors.New("获取资源失败")
		return
	}

	heroInfo.Lv += 1
	heroInfo.Wise += 1
	heroInfo.Diligent += 1
	heroInfo.Loyal += 1
	heroInfo.Heroic += 3

	err = herosvc.UpgradeHeroLv(heroInfo)
	if err != nil {
		resp.Code = app.UpdateDataError
		resp.Lv = int(heroInfo.Lv -1 )
		resp.Resource = resourceInfos
		err = errors.New("更新用户英雄等级出错")
		return
	}

	resp.Code = 0
	resp.Lv = int(heroInfo.Lv)
	resp.Resource = resourceInfos

	return
}
