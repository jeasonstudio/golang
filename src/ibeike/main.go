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
	tagLoginURL := "http://pt.ibeike.com/takelogin.php"

	v := url.Values{}
	v.Set("username", "41524119")
	v.Set("password", "1024001x")
	v.Set("submitbutton", "提 交")
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{}
	req, _ := http.NewRequest("POST", tagLoginURL, body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去，被坑了两小时
	// fmt.Printf("%+v\n", req)                                                         //看下发送的结构
	// req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.95 Safari/537.36")
	// req.Header.Set("Content-Length", "42")
	req.Header.Set("Cookie", "__utmt=1; __utma=213053004.1368731933.1482642887.1482642887.1482642887.1; __utmb=213053004.1.10.1482642887; __utmc=213053004; __utmz=213053004.1482642887.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); PHPSESSID=shcim9cvious1l66i6r8kcb8f6; cJcK_b943_saltkey=EG88CNWg; cJcK_b943_lastvisit=1482639299; cJcK_b943_sid=RwWroW; cJcK_b943_lip=10.18.32.194%2C1482551952; cJcK_b943_creditnotice=0D1D2D0D0D0D0D0D0D216473; cJcK_b943_creditbase=0D70D236D0D0D0D0D0D1067; cJcK_b943_creditrule=%E6%AF%8F%E5%A4%A9%E7%99%BB%E5%BD%95; cJcK_b943_lastact=1482643207%09uc.php%09")

	resp, err := client.Do(req) //发送
	defer resp.Body.Close()     //一定要关闭resp.Body
	// data, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(data), err)
	if err != nil {
		return
	}
	// tagCookie := strings.Split(strings.Split(resp.Header["Set-Cookie"][0], ";")[0], "=")[1]

	fmt.Println(resp.Header, err)
	fmt.Println(resp.StatusCode)

	// getAll(tagCookie)
	// fmt.Println(resp.Header["Set-Cookie"][0])

}
