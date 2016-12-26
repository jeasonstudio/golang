package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func getAll(tagCookie string) {
	thisCookie := "JSESSIONID=" + tagCookie
	tagLoginURL := "http://elearning.ustb.edu.cn/choose_courses/loginsucc.action"

	client := &http.Client{}

	req, err := http.NewRequest("GET", tagLoginURL, strings.NewReader(""))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", thisCookie)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(resp.Body)

}

func main() {
	tagLoginURL := "http://elearning.ustb.edu.cn/choose_courses/j_spring_security_check"

	v := url.Values{}
	v.Set("j_username", "41524122,undergraduate")
	v.Set("j_password", "07060016")
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{}
	req, _ := http.NewRequest("POST", tagLoginURL, body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去，被坑了两小时
	// fmt.Printf("%+v\n", req)                                                         //看下发送的结构
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.95 Safari/537.36")
	req.Header.Set("Cookie", "JSESSIONID=C010D246699DC59E6C29F0C2932244B3; _ga=GA1.3.1043019656.1481126573; JSESSIONID=00AA0A5E7C5C4B3FBECF9E9350594A1A")

	resp, err := client.Do(req) //发送
	defer resp.Body.Close()     //一定要关闭resp.Body
	// data, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(data), err)
	if err != nil {
		return
	}
	tagCookie := strings.Split(strings.Split(resp.Header["Set-Cookie"][0], ";")[0], "=")[1]

	// fmt.Println(resp.Header["Set-Cookie"], err)
	fmt.Println(tagCookie)

	getAll(tagCookie)
	// fmt.Println(resp.Header["Set-Cookie"][0])

}
