package main

import (
	"fmt"
	"io"
	"net/http"
)

const (
	HTTP_PORT string = "80"
	TEST_PORT string = "9090"
)

func goAjax(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "这是从后台发送的数据")

	res.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	res.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	res.Header().Set("content-type", "application/json")             //返回数据格式是json

	req.ParseForm()
	fmt.Println("收到客户端请求: ", req.Form)
	fmt.Fprintf(res, "{'a':'222'}")
}

func main() {

	http.HandleFunc("/", goAjax)

	err := http.ListenAndServe(":"+TEST_PORT, nil)
	if err != nil {
		fmt.Println("服务失败 /// ", err)
	}
}
