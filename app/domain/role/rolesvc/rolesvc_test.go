package rolesvc

import (
	"fmt"
	"testing"
)

//func TestRegister(t *testing.T) {
//	//app.Init()
//	code := Register("10","Tyrion","11111")
//	if code != 0{
//		t.Errorf("注册失败:%d",code)
//	}
//}

func TestLoginVerification(t *testing.T) {
	Correct, err := LoginVerification("14","11111")
	if !Correct {
		t.Errorf("login Error:%s",err)
	}
	fmt.Println("账号验证成功")
}

func TestGetRoleInfo(t *testing.T) {
	info,err:=GetRoleInfo("23")
	fmt.Println(info)
	if err != nil{
		t.Error(err)
	}
}

func TestInitResourceFresh(t *testing.T) {
	err:=InitResourceFresh("22")
	if err != nil {
		t.Errorf("initErr:%s", err)
	}
}

func TestIsExitRoleInfo(t *testing.T) {
	blog := IsExitRoleInfo("23")
	log.Info(blog)
}
