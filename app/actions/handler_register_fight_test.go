package actions

import (
	"fmt"
	"testing"
)

func TestRegisterFightOperate(t *testing.T) {
	info,err := RegisterFightOperate("10",1,30)
	if err != nil{
		t.Error(err)
	}
	fmt.Println(info)
}
