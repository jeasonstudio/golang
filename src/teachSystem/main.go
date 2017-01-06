package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// func main()  {
//     tagLoginURL := "http://elearning.ustb.edu.cn/choose_courses/choosecourse/normalChooseCourse_normalRequired_loadPreNormalAccordByKchRequiredCourses.action?kch=2050414&_dc=1483270853482&limit=5000&start=0&uid=41524122"

// 	client := &http.Client{}
// 	req, _ := http.NewRequest("GET", tagLoginURL, nil)

// 	// req.Header.Set("Cookie", thisCookie)

// 	resp, err := client.Do(req) //发送
// 	defer resp.Body.Close()     //一定要关闭resp.Body
// 	data, _ := ioutil.ReadAll(resp.Body)
// 	fmt.Println(string(data), err)
// 	if err != nil {
// 		return
// 	}
// }

// func defaultCheckRedirect(req *Request, via []*Request) error {
// 	if len(via) >= 10 {
// 		return errors.New("stopped after 10 redirects")
// 	}
// 	return nil
// }

func typeSwitch(tst interface{}) {
    switch v := tst.(type) {
        case string:
           fmt.Println("Stringer:", v)
		case  []interface{}:
			fmt.Println("in")
        default:
           fmt.Println("Unknown")
    }
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

	// tagCookie := strings.Split(resp.Request.URL,";")[1]

	// aaaa := funcName(resp.Request.URL)
	// var interCookie interface{} = resp.Request.URL
    // j := interCookie.(string)

	// fmt.Printf("%s",resp.Request.URL)

	// j := resp.Request.URL.(string)
	// var idx interface {}
	var tag string
	switch v := resp.Request.URL.(type) {
		default:
			tag = string(v)
		case int, int8, int16, int32, int64:
		case float32:
		case float64:
	}

	fmt.Println(tag)
	fmt.Println(string(data))

}
