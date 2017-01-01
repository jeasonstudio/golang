package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func getAll(tagCookie string) {
	thisCookie := "PHPSESSID=" + tagCookie
	tagLoginURL := "http://pt.ibeike.com/browse.php"

	client := &http.Client{}
	req, _ := http.NewRequest("GET", tagLoginURL, nil)

	req.Header.Set("Cookie", thisCookie)

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

	v := url.Values{"username": {"41524119"}, "password": {"1024001x"}, "submitbutton": {"提 交"}}

	resp, err := http.PostForm(tagLoginURL, v)
	if err != nil {
		return
	}
	// data, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(data))

	defer resp.Body.Close()

	tagCookie := strings.Split(strings.Split(resp.Header["Set-Cookie"][0], ";")[0], "=")[1]
	fmt.Println(tagCookie)
	getAll(tagCookie)

}
