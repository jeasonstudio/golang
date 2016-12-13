package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func write(str string) {
	userFile := "test.txt"
	fout, err := os.Create(userFile)
	defer fout.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	fout.WriteString(str)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	r.ParseForm()                                                  //解析参数，默认是不会解析的
	// fmt.Println(r.Form)                                            //这些信息是输出到服务器端的打印信息
	// fmt.Println("path", r.URL.Path)
	// fmt.Println("scheme", r.URL.Scheme)
	// fmt.Println(r.Form["url_long"])

	// for k, v := range r.Form {
	// 	fmt.Println("key:", k)
	// 	fmt.Println("val:", strings.Join(v, ""))
	// 	// write(strings.Join(v, "")
	// 	// setSQL(k, strings.Join(v, ""))
	// }
	// fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的

	// l, err := url.ParseQuery(urlStr)

	id := strings.Join(r.Form["id"], "")
	username := strings.Join(r.Form["username"], "")
	name := (strings.Join(r.Form["name"], ""))
	unitInfo := (strings.Join(r.Form["unit_info"], ""))
	mobile := strings.Join(r.Form["mobile"], "")
	fmt.Println(id)
	fmt.Println(username)
	fmt.Println(name)
	fmt.Println(unitInfo)
	fmt.Println(mobile)

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/stu_phone_num?charset=utf8")
	checkErr(err)

	stmt, err := db.Prepare("INSERT stu SET s_id=?,username=?,name=?,unit_info=?,mobile=?")
	checkErr(err)

	res, err := stmt.Exec(id, username, name, unitInfo, mobile)
	checkErr(err)

	fmt.Println(res)

	defer db.Close()
}

func main() {
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
