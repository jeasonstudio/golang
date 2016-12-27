package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getContentMoney(tagCookie string) {
	thisCookie := "_ga=GA1.3.1043019656.1481126573; iPlanetDirectoryPro=" + tagCookie
	tagLoginURL := "http://e.ustb.edu.cn/index.portal?.pn=p378_p381"

	client := &http.Client{}

	req, err := http.NewRequest("GET", tagLoginURL, nil)
	if err != nil {
		return
	}
	req.Header.Set("Cookie", thisCookie)
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Encoding", "gzip, deflate, sdch")

	// fmt.Printf("%+v\n", req) //看下发送的结构

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if resp.StatusCode == 200 {

		body, _ := ioutil.ReadAll(resp.Body)
		str := strings.NewReader(string(body))

		doc, _ := goquery.NewDocumentFromReader(str)

		fmt.Printf("\n校园卡消费明细：   \n")

		doc.Find(".portletContent table tbody tr td table tbody").Each(func(i int, s *goquery.Selection) {
			s.Find("tr").Each(func(i int, b *goquery.Selection) {
				band := b.Find("td").Text()
				if i == 0 {
					band = "消费时间         地点 消费金额 剩余金额"
					fmt.Printf("序号     %s \n", band)
				} else {
					// m := strings.Index(band, ":")
					// len := strings.Count(band, "")
					// date := strings.SplitN(band, "", len-5)
					// // raceMoney := strings.SplitN(band, "", len-5)
					// fmt.Println(date)
					fmt.Printf("%-2d:  %s 元 \n", i, band)
				}
			})
		})
	}
}

func getCardMoney(tagCookie string) {
	thisCookie := "_ga=GA1.3.1043019656.1481126573; iPlanetDirectoryPro=" + tagCookie
	tagLoginURL := "http://e.ustb.edu.cn/index.portal"

	client := &http.Client{}

	req, err := http.NewRequest("GET", tagLoginURL, nil)
	if err != nil {
		return
	}
	req.Header.Set("Cookie", thisCookie)
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Encoding", "gzip, deflate, sdch")

	// fmt.Printf("%+v\n", req) //看下发送的结构

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if resp.StatusCode == 200 {

		body, _ := ioutil.ReadAll(resp.Body)

		src := string(body)
		//将HTML标签全转换成小写
		re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
		src = re.ReplaceAllStringFunc(src, strings.ToLower)
		//  提取table 标签
		re, _ = regexp.Compile("\\<!doc[\\S\\s]+?\\<font color='red'>")
		src = re.ReplaceAllString(src, "<font color='red'>")
		re, _ = regexp.Compile("</font\\>[\\S\\s]+?\\</html\\>")
		src = re.ReplaceAllString(src, "</font>")

		resu := strings.Split(strings.Split(src, "<")[1], ">")[1]
		fmt.Println("校园卡余额为：" + resu + "元")

		// doc, _ := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
		// fmt.Println(doc.Find("font"))

	}

	if err != nil {
		return
	}
}

func main() {
	tagLoginURL := "http://e.ustb.edu.cn/userPasswordValidate.portal"

	v := url.Values{"Login.Token1": {"41524122"}, "Login.Token2": {"060016"}}

	resp, err := http.PostForm(tagLoginURL, v)
	if err != nil {
		return
	}

	defer resp.Body.Close() //一定要关闭resp.Body

	if err != nil {
		return
	}

	tagCookie := strings.Split(strings.Split(resp.Header["Set-Cookie"][0], ";")[0], "=")[1]

	getCardMoney(tagCookie)
	getContentMoney(tagCookie)
}
