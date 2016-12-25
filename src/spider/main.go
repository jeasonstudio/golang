package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func getAll(tagCookie string) {
	thisCookie := " JSESSIONID=" + tagCookie
	tagLoginURL := "http://elearning.ustb.edu.cn/choose_courses/choosecourse/commonChooseCourse_courseList_loadTermCourses.action"

	fmt.Println(thisCookie)

	v := url.Values{}
	v.Set("listXnxq", "2016-2017-1")
	v.Set("uid", "41524122")
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{}
	req, _ := http.NewRequest("POST", tagLoginURL, body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去，被坑了两小时
	// fmt.Printf("%+v\n", req)                                                         //看下发送的结构
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.95 Safari/537.36")
	req.Header.Set("Cookie", thisCookie)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	// userCookie := $http.Cookie{
	// 	Name:     "JSESSIONID",
	// 	Value:    tagCookie,
	// 	Path:     "/",
	// 	HttpOnly: true,
	// }

	// req.AddCookie(userCookie)

	resp, err := client.Do(req) //发送
	defer resp.Body.Close()     //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data), err)
	if err != nil {
		return
	}

}

func main() {
	tagLoginURL := "http://bj.lianjia.com/api/newhouserecommend"

	v := url.Values{}
	v.Set("type", "1")
	v.Set("query", "http://bj.lianjia.com/ershoufang/rs海淀/")
	// v.Set("lastUrl", "")
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{}
	req, _ := http.NewRequest("GET", tagLoginURL, body)

	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去，被坑了两小时
	// fmt.Printf("%+v\n", req)                                                         //看下发送的结构
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.95 Safari/537.36")
	// req.Header.Set("Content-Length", "42")
	// req.Header.Set("Cookie", "lianjia_uuid=4242529c-385e-4683-b572-e7cc89a9039d; select_city=110000; all-lj=80b391239fd880f59f779618fca39507; _gat=1; _gat_past=1; _gat_global=1; _gat_new_global=1; _gat_dianpu_agent=1; logger_session=0e3961d782a43dd6c22a3692833594e9; _ga=GA1.2.1571583035.1482682480; CNZZDATA1253477573=208171418-1482677452-null%7C1482677452; _smt_uid=585ff06e.4d27e347; CNZZDATA1254525948=1872321740-1482680111-null%7C1482680111; CNZZDATA1255633284=1260721563-1482681473-null%7C1482681473; CNZZDATA1255604082=661565453-1482678657-null%7C1482678657; lianjia_ssid=e2a2989f-85d0-49ed-9b51-4cab4476f14a")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")

	resp, err := client.Do(req) //发送
	defer resp.Body.Close()     //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data), err)
	if err != nil {
		return
	}
	fmt.Println(resp.StatusCode)

}
