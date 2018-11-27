package model

import (
	"beginner-server/app/model/stat"
	"github.com/cxr29/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


type ProfileConfig struct{
	ConfLever   	[]stat.ConfLever
	ConfHero	[]stat.ConfHero
}


var ProfConfig =  ProfileConfig{}

func GetLink() *gorm.DB {
	db, err:= gorm.Open("mysql", "root:root@tcp(localhost:3306)/beginner_databases?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		log.Errorf("连接数据库失败:%v",err)
		panic("连接数据库失败")
	}
	return db
}
