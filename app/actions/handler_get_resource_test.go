package actions

import (
	"fmt"
	"testing"
)

func TestGetResourceOperate(t *testing.T) {
	info, err := GetResourceOperate("22", 3)
	if err != nil {
		t.Errorf("收集资源出错:%s", err)
	}
	fmt.Println(info)
}