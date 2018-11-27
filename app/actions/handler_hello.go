package actions

import (
	"fmt"
	"net/http"
)

const (
	//打招呼模板
	greetingFmt = "Aloha %s!"
	//默认名字
	defaultName = "Stranger"
)

//Hello路由的请求参数字段
type HelloParam struct {
	MyName string //请求者的名字
}

//Hello路由的响应字段
type HelloResp struct {
	Code     int    //返回码
	Greeting string //问候语
}

//Hello路由的Handler
//	Hello的功能是，如果传入的MyName不为空字符传，则返回<问候语>+<请求者名字>，
//	若MyName为空字符串，则返回<问候语>+"Stranger"
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	param := HelloParam{}
	resp := HelloResp{
		Code: 0,
	}
	err := parseRequest(r, &param)
	if err != nil {
		resp.Code = 1

	}
	resp.Greeting = SayHello(param.MyName)
	writeResponse(w, resp)
}

//输入名字，返回问候语
// 这里之所以单独写一个函数是为了方便测试时调用。
// Handler里负责处理请求相关的字段，格式化返回，
// 而SayHello我们将其称之为业务函数，不关心协议（是http短连接还是websocket长连接，用json还是msgpack），
// 只关注具体业务逻辑
func SayHello(myName string) (greeting string) {
	if myName == "" {
		myName = defaultName
	}
	return fmt.Sprintf(greetingFmt, myName)
}
