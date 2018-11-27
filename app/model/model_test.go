package model

import (
	"fmt"
	"testing"
)

func TestGetLink(t *testing.T){
	orm := GetLink()
	fmt.Print(orm)
}