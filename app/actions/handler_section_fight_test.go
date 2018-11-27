package actions

import (
	"fmt"
	"testing"
)

func TestSectionFightOperate(t *testing.T) {
	resp,err := SectionFightOperate("24",1)
	if err != nil{
		t.Error(err)
	}
	fmt.Println(resp)
}
