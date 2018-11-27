package actions

import (
	"encoding/json"
	"errors"
	"net/http"
)

//解析请求body中的json，写入到结构体中, param必须是结构体的指针
func parseRequest(r *http.Request, param interface{}) error {
	if r.Body == nil {
		return errors.New("request body is empty")
	}
	err := json.NewDecoder(r.Body).Decode(param)
	if err != nil {
		return err
	}
	return nil
}

//序列化响应字段为json，并输出到http响应中，resp必须是结构体
func writeResponse(w http.ResponseWriter, resp interface{}) error {
	b, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}
