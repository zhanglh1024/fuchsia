package actions

import (
	"fmt"
	"testing"
)

func TestUpgradeHeroLvOperate(t *testing.T) {
	info,err := UpgradeHeroLvOperate("22",1)
	if err!=nil{
		t.Error(err)
	}
	fmt.Println(info)
}
