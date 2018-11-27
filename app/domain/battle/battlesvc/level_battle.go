package battlesvc

import (
	"beginner-server/app/domain/battle/battleobj"
	"beginner-server/app/domain/hero/heroobj"
	"beginner-server/app/model"
	"beginner-server/app/model/dyn"
	"errors"
	"fmt"
)


//通关处理
func PassEditLevelInfo(levelInfo dyn.LevelBattle) (error, int) {
	if levelInfo.RoundId < 5{
		fmt.Println(levelInfo.RoundId,model.ProfConfig.ConfLever)
		info := model.ProfConfig.ConfLever[levelInfo.RoundId]
		hero := model.ProfConfig.ConfHero[levelInfo.RoundId]
		levelInfo.RoundId = int(info.No)
		levelInfo.SoldierLeft = info.SoldierNum
		levelInfo.HeroFight = info.HeroFight
		err := battleobj.UpdateLevelBattleValue(levelInfo)
		if err != nil{
			return errors.New("下一关，关卡配置失败"), 0
		}

		//通关获取新英雄
		heroInfo := dyn.HeroInfoDyn{}
		heroInfo.RoleId = levelInfo.RoleId
		heroInfo.No = int(hero.HeroNum)
		heroInfo.Lv = 1
		heroInfo.Wise = hero.InitWise
		heroInfo.Loyal = hero.InitLoyalty
		heroInfo.Diligent = hero.InitDiligent
		heroInfo.Heroic = hero.InitHeroic
		err = heroobj.InsertHeroValue(heroInfo)

		if err!= nil{
			return errors.New("获取新英雄失败"), 0
		}

		return nil, 0
	}

	if levelInfo.RoundId == 5{
		levelInfo.RoundId = 6
		err := battleobj.UpdateLevelBattleValue(levelInfo)
		if err != nil{
			return errors.New("玩家通关设置失败"), 0
		}
		return nil, 1
	}

	return errors.New("关卡数出现错误"), 0
}

//未通关更新关卡资源信息
func EditCurrentLevel(cost int,levelInfo dyn.LevelBattle) error {
	levelInfo.SoldierLeft -= cost
	return battleobj.UpdateLevelBattleValue(levelInfo)
}


//获取用户当前光卡信息
func GetCurrentLevelInfo(roleId string) (dyn.LevelBattle, error) {
	info ,err := battleobj.GetLevelBattleValue(roleId)
	return info, err
}