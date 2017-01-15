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

const (
	COURSE_TABLE = "http://elearning.ustb.edu.cn/choose_courses/choosecourse/commonChooseCourse_courseList_loadTermCourses.action"
	// listXnxq:2016-2017-2
	// uid:xxxxxxxx
	INNVOATION_SCORE = "http://elearning.ustb.edu.cn/choose_courses/information/singleStuInfo_singleStuInfo_loadSingleStuCxxfPage.action"
	// uid:xxxxxxxx
	ALL_COURSE_SCORE = "http://elearning.ustb.edu.cn/choose_courses/information/singleStuInfo_singleStuInfo_loadSingleStuScorePage.action"
	// uid:xxxxxxxx
)

func getCourse(jsessionid string)  {

	v := url.Values{"listXnxq": {LIST_XNXQ}, "uid": {UID}}
	body := ioutil.NopCloser(strings.NewReader(v.Encode()))

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, COURSE_TABLE, body)

	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "JSESSIONID="+jsessionid)

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(data))

}

func getInnvoationScore(jsessionid string)  {

	v := url.Values{"uid": {UID}}
	body := ioutil.NopCloser(strings.NewReader(v.Encode()))

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, INNVOATION_SCORE, body)

	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "JSESSIONID="+jsessionid)

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	str := strings.NewReader(string(data))

	doc, _ := goquery.NewDocumentFromReader(str)

	res := doc.Find("td").Text()

	// fmt.Println(string(data))

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
	tagCookies := strings.Split(strings.Split(res,";")[1],"=")[1]

	fmt.Println(tagCookies)
	fmt.Println(string(data))

	getInnvoationScore(tagCookies)
}
