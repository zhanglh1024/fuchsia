package actions

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestSayHello(t *testing.T) {
	//这个第三方包可以用来进行测试的断言，比go自带的testing.T要方便
	ast := assert.New(t)

	//先测试传入字符串为空的情况
	name := ""
	greeting := SayHello(name)
	t.Log("Greeting if name is empty: ", greeting)
	ast.True(strings.Contains(greeting, defaultName))

	name = "Eddie"
	greeting = SayHello(name)
	t.Logf("Greeting if name is %s: %s", name, greeting)

	ast.True(strings.Contains(greeting, name))
}
