package app

import (
	"github.com/cxr29/log"
	"beginner-server/app/model"
	"beginner-server/app/model/stat"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func CloseDb() {
	if DB != nil {
		log.Info("关闭数据库")
		DB.Close()
	}
	log.Info("关闭数据库")
}

func init(){
	var lever  []stat.ConfLever
	var hero  []stat.ConfHero

	DB = model.GetLink()
	err := DB.Find(&lever).Error
	if err != nil {
		log.Errorf("读取关卡配置出错：%s", err)
	}
	model.ProfConfig.ConfLever = lever

	err = DB.Find(&hero).Error
	if err != nil {
		log.Errorf("读取英雄配置出错:%s", err)
	}
	log.Info("连接数据库成功")
	model.ProfConfig.ConfHero = hero

}