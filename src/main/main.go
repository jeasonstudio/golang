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

func setSQL(key string, value string) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/stu_phone_num?charset=utf8")
	if err != nil {
		log.Println(err)
	}

	// 插入一条新数据
	result, err := db.Exec("INSERT INTO `stu`(" + key + ") VALUES(" + value + ")")
	if err != nil {
		fmt.Println("insert data failed:", err.Error())
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("fetch last insert id failed:", err.Error())
		return
	}
	fmt.Println("insert new record", id)
	//在这里进行一些数据库操作

	defer db.Close()
}

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
	id := strings.Join(r.Form["id"], "")
	username := strings.Join(r.Form["username"], "")
	name := strings.Join(r.Form["name"], "")
	unitInfo := strings.Join(r.Form["unit_info"], "")
	mobile := strings.Join(r.Form["mobile"], "")
	fmt.Println(id)
	fmt.Println(username)
	fmt.Println(name)
	fmt.Println(unitInfo)
	fmt.Println(mobile)

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/stu_phone_num")
	if err != nil {
		log.Println(err)
	}

	// 插入一条新数据
	result, err := db.Exec("INSERT INTO `stu`(`id`,`username`,`name`,`unit_info`,`mobile`) VALUES(" + id + "," + username + "," + name + "," + unitInfo + "," + mobile + ")")
	if err != nil {
		fmt.Println("insert data failed:", err.Error())
		return
	}
	// id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("fetch last insert id failed:", err.Error())
		return
	}
	fmt.Println("insert new record", result)
	//在这里进行一些数据库操作

	defer db.Close()
}

func main() {
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
