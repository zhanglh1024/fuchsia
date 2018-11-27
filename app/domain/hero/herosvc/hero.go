package herosvc

import (
	"beginner-server/app"
	"beginner-server/app/domain/hero/heroobj"
	"beginner-server/app/model/dyn"
	"errors"
)

//获取用户名下所有英雄信息
func GetLoginHeroInfo(roleId string) ([]app.HeroInfo, error) {
	heroInfos, err := heroobj.GetHeroInfoByRoleId(roleId)
	if err != nil {
		errors.New("查询出错")
		return nil,err
	}
	infos := make([]app.HeroInfo,0)
	info:=app.HeroInfo{}
	for _,herInfo := range heroInfos{
		info.No = herInfo.No
		info.Lv = int(herInfo.Lv)
		infos = append(infos,info)
	}

	return infos, nil
}

//根据用户编号和英雄编号获取英雄信息
func GetHeroInfoByRoleIdAndHeroNo(roleId string, heroNo int) (dyn.HeroInfoDyn, error){
	info, err := heroobj.GetHeroInfoByRoleIdAndHeroNo(roleId, heroNo)
	return info, err
}

//根据用户编号和英雄编号获取英雄战斗力
func GetHeroFightByRoleIdAndHeroNo(roleId string, heroNo int) (int,error) {
	info, err := heroobj.GetHeroInfoByRoleIdAndHeroNo(roleId, heroNo)
	count := info.Heroic+info.Diligent+info.Loyal+info.Wise
	return count, err
}

//计算英雄战斗力
func CalculateHeroCombatPower(soldierNum, heroFight int) int {
	return soldierNum + heroFight*5
}


//计算输赢玩家消耗的士兵数
func ConsumingSoldiers(soldierNum, heroFight, roundId,loseSoldier,loseHeroFight int) (winCostSoldier,loseCostSoldier int)  {
	winCostSoldier = (soldierNum*5)/heroFight + 5*roundId
	loseCostSoldier = (loseSoldier*8)/loseHeroFight + 5*roundId
	return
}

//升级英雄等级
func UpgradeHeroLv(heroInfo dyn.HeroInfoDyn) (error) {
	return heroobj.UpdateHeroInfo(heroInfo)
}