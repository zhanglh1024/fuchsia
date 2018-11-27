package roleobj

import (
	"fmt"
	"log"
	"testing"
)

func TestInsertRoleInfo(t *testing.T)  {
	this := RoleObj{}
	flag := this.InsertRoleInfo("1236", "王五", "888796")
	if flag==0 {
		fmt.Println("插入成功")
	}

	err := DeleteRoleInfoForTest("1236")
	if err != nil{
		t.Error(err)
	}
}

func TestGetAllRoleInfo(t *testing.T){
	this := RoleObj{}
	roleInfo := this.GetAllRoleInfo()
	fmt.Println("测试输出", roleInfo)
}

func TestGetRoleInfoByRoleId(t *testing.T) {
	this := RoleObj{}

	flag := this.InsertRoleInfo("1236", "王五", "888796")
	if flag==0 {
		fmt.Println("插入成功")
	}

	roleInfo, err := this.GetRoleInfoByRoleId("1236")
	if err != nil {
		log.Fatal("err")
	}
	fmt.Println(roleInfo)

	err = DeleteRoleInfoForTest("1236")
	if err != nil{
		t.Error(err)
	}
}

func TestRoleObj_IsIncludeRoleInfoName(t *testing.T) {
	this := RoleObj{}
	flag := this.InsertRoleInfo("1236", "王五", "888796")
	if flag==0 {
		fmt.Println("插入成功")
	}
	num := this.IsIncludeRoleInfoName("王五")
	fmt.Println(num)
	err := DeleteRoleInfoForTest("1236")
	if err != nil{
		t.Error(err)
	}
}

func TestRoleObj_IsIncludeRoleInfoRoleId(t *testing.T) {
	this := RoleObj{}
	flag := this.InsertRoleInfo("1236", "王五", "888796")
	if flag==0 {
		fmt.Println("插入成功")
	}
	num := this.IsIncludeRoleInfoRoleId("1236")
	fmt.Println(num)
	err := DeleteRoleInfoForTest("1236")
	if err != nil{
		t.Error(err)
	}
}

func TestDeleteRoleInfoForTest(t *testing.T) {
	this := RoleObj{}
	flag := this.InsertRoleInfo("1236", "王五", "888796")
	if flag==0 {
		fmt.Println("插入成功")
	}

	err := DeleteRoleInfoForTest("1236")
	if err != nil{
		t.Error(err)
	}
}