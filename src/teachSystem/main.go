package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)
const (
	LIST_XNXQ = "2016-2017-1"
	UID = "41524122"
)

func getCourse(jsessionid string)  {
	getCourseUrl := "http://elearning.ustb.edu.cn/choose_courses/choosecourse/commonChooseCourse_courseList_loadTermCourses.action"

	v := url.Values{"listXnxq": {LIST_XNXQ}, "uid": {UID}}
	body := ioutil.NopCloser(strings.NewReader(v.Encode()))

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, getCourseUrl, body)

	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", jsessionid)

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(data))

}

func main() {
	tagLoginURL := "http://elearning.ustb.edu.cn/choose_courses/j_spring_security_check"

	v := url.Values{"j_username": {"41524122,undergraduate"}, "j_password": {"07060016"}}
	body := ioutil.NopCloser(strings.NewReader(v.Encode()))

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, tagLoginURL, body)

	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	res := fmt.Sprintf("%s",resp.Request.URL)
	tagCookies := strings.Split(res,";")[1]

	fmt.Println(tagCookies)
	fmt.Println(string(data))

	getCourse(tagCookies)
}
