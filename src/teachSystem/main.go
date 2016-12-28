package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
	tagLoginURL := "http://teach.ustb.edu.cn/"

	client := &http.Client{}
	req, _ := http.NewRequest("GET", tagLoginURL, nil)

	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,it;q=0.6,fr;q=0.4,en;q=0.2,zh-TW;q=0.2")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")

	resp, err := client.Do(req) //发送
	defer resp.Body.Close()     //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data), err)
	if err != nil {
		return
	}

}
