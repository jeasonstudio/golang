package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getAll(tagCookie string) {
	thisCookie := "JSESSIONID=" + tagCookie
	tagLoginURL := "http://zhiyuan.ustb.edu.cn/app.VPClient/index.jsp?m=vpclient&c=index&a=showIndex"

	client := &http.Client{}
	req, _ := http.NewRequest("GET", tagLoginURL, nil)

	req.Header.Set("Cookie", thisCookie)

	resp, err := client.Do(req) //发送
	defer resp.Body.Close()     //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(data), err)
	if err != nil {
		return
	}

	str := strings.NewReader(string(data))

	doc, _ := goquery.NewDocumentFromReader(str)

	res := doc.Find(".grxx_text .jj li span a").Text()

	fmt.Println("我已获得志愿工时：" + res + "小时")

}

func main() {
	tagLoginURL := "https://www.zhihu.com/login/phone_num"

	v := url.Values{"phone_num": {"13220101996"}, "password": {"Zjt13832913646"}, "lastUrl": {""}}

	resp, err := http.PostForm(tagLoginURL, v)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {

		fmt.Println(resp.Header["Set-Cookie"])

		// tagCookie := strings.Split(strings.Split(resp.Header["Set-Cookie"][0], ";")[0], "=")[1]

		// fmt.Println(tagCookie)
		// fmt.Println(resp.StatusCode)

		// getAll(tagCookie)
	}

}
