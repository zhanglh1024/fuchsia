package actions

import (
	"fmt"
	"testing"
)

func TestMatchFightOperate(t *testing.T) {
	info, err := MatchFightOperate("23",2)
	if err != nil{
		t.Error(err)
	}
	fmt.Println(info)
}
