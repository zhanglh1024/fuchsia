package actions

import (
	"fmt"
	"testing"
)

func TestLoginOperate(t *testing.T) {
	resp, err := LoginOperate("12", "11111")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}
