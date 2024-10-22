package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"runtime"
	"time"

	_ "github.com/Go-SQL-Driver/MySQL"
	simplejson "github.com/bitly/go-simplejson"
)

//chan中存入string类型的href属性,缓冲200
var urlChannel = make(chan string, 2000)

var userAgent = [...]string{"Mozilla/5.0 (compatible, MSIE 10.0, Windows NT, DigExt)",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, 360SE)",
	"Mozilla/4.0 (compatible, MSIE 8.0, Windows NT 6.0, Trident/4.0)",
	"Mozilla/5.0 (compatible, MSIE 9.0, Windows NT 6.1, Trident/5.0,",
	"Opera/9.80 (Windows NT 6.1, U, en) Presto/2.8.131 Version/11.11",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, TencentTraveler 4.0)",
	"Mozilla/5.0 (Windows, U, Windows NT 6.1, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Macintosh, Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
	"Mozilla/5.0 (Macintosh, U, Intel Mac OS X 10_6_8, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Linux, U, Android 3.0, en-us, Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
	"Mozilla/5.0 (iPad, U, CPU OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, Trident/4.0, SE 2.X MetaSr 1.0, SE 2.X MetaSr 1.0, .NET CLR 2.0.50727, SE 2.X MetaSr 1.0)",
	"Mozilla/5.0 (iPhone, U, CPU iPhone OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"MQQBrowser/26 Mozilla/5.0 (Linux, U, Android 2.3.7, zh-cn, MB200 Build/GRJ22, CyanogenMod-7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1"}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func GetRandomUserAgent() string {
	return userAgent[r.Intn(len(userAgent))]
}

func Spy(tagURL string) bool {

	// 程序异常终止函数
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		log.Println("[E]", r)
	// 	}
	// }()

	userAgent := GetRandomUserAgent()

	client := &http.Client{}
	req, _ := http.NewRequest("GET", tagURL, nil)

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Referer", "https://www.zhihu.com/people/excited-vczh/followers")
	req.Header.Set("Host", "www.zhihu.com")
	req.Header.Set("authorization", "oauth c3cef7c66a1843f8b3a9e6a1e3160e20")

	resp, err := client.Do(req) //发送
	if err != nil {
		return false
	}
	defer resp.Body.Close() //一定要关闭resp.Body
	bodyData, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(bodyData))
	if resp.StatusCode != 200 {
		fmt.Println("HTTP GET Error!")
		return false
	}
	myjson, _ := simplejson.NewJson(bodyData)

	peopleData, _ := myjson.Get("data").Array()

	for i := 0; i < len(peopleData); i++ {
		// fmt.Println(myjson.Get("data").GetIndex(i).Get("id").String())
		// fmt.Println(i)
		urlToken, _ := myjson.Get("data").GetIndex(i).Get("url_token").String()
		follows, _ := myjson.Get("data").GetIndex(i).Get("follower_count").Int()

		if follows > 4 {
			href := "https://www.zhihu.com/api/v4/members/" + urlToken + "/followees?include=data[*].answer_count,articles_count,follower_count,is_followed,is_following,badge[?(type=best_answerer)].topics&offset=0&limit=500"
			urlChannel <- href

			iName, _ := myjson.Get("data").GetIndex(i).Get("name").String()
			isFollowed, _ := myjson.Get("data").GetIndex(i).Get("is_followed").Bool()
			avatarURLTemplate, _ := myjson.Get("data").GetIndex(i).Get("avatar_url_template").String()
			url, _ := myjson.Get("data").GetIndex(i).Get("url").String()
			userType, _ := myjson.Get("data").GetIndex(i).Get("user_type").String()
			answerCount, _ := myjson.Get("data").GetIndex(i).Get("name").Int()
			urlToken, _ := myjson.Get("data").GetIndex(i).Get("url_token").String()
			isAdvertiser, _ := myjson.Get("data").GetIndex(i).Get("is_advertiser").Bool()
			avatarURL, _ := myjson.Get("data").GetIndex(i).Get("avatar_url").String()
			isFollowing, _ := myjson.Get("data").GetIndex(i).Get("is_following").Bool()
			isOrg, _ := myjson.Get("data").GetIndex(i).Get("is_org").Bool()
			headline, _ := myjson.Get("data").GetIndex(i).Get("headline").String()
			followerCount, _ := myjson.Get("data").GetIndex(i).Get("follower_count").Int()
			ttype, _ := myjson.Get("data").GetIndex(i).Get("type").String()
			id, _ := myjson.Get("data").GetIndex(i).Get("id").String()
			articlesCount, _ := myjson.Get("data").GetIndex(i).Get("articles_count").Int()

			insertDB(iName, isFollowed, avatarURLTemplate, url, userType, answerCount, urlToken, isAdvertiser, avatarURL, isFollowing, isOrg, headline, followerCount, ttype, id, articlesCount)
		}

	}

	return true
}

func insertDB(name string, is_followed bool, avatar_url_template string, url string, user_type string, answer_count int, url_token string, is_advertiser bool, avatar_url string, is_following bool, is_org bool, headline string, follower_count int, the_type string, id string, articles_count int) {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/zhihu_user?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}

	stmt, err := db.Prepare("INSERT user SET name=?,is_followed=?,avatar_url_template=?,url=?,user_type=?,answer_count=?,url_token=?,is_advertiser=?,avatar_url=?,is_following=?,is_org=?,headline=?,follower_count=?,the_type=?,id=?,articles_count=?")
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := stmt.Exec(name, is_followed, avatar_url_template, url, user_type, answer_count, url_token, is_advertiser, avatar_url, is_following, is_org, headline, follower_count, the_type, id, articles_count)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	// fmt.Println(res)
	if res == nil {
		fmt.Println("An Sql Error!")
	}

}

func main() {

	// 种子：轮子哥(vczh)
	go Spy("https://www.zhihu.com/api/v4/members/excited-vczh/followees?include=data[*].answer_count,articles_count,follower_count,is_followed,is_following,badge[?(type=best_answerer)].topics&offset=0&limit=500")

	for url := range urlChannel {
		// 获取当前运行时的一些相关参数等
		fmt.Println("routines num = ", runtime.NumGoroutine(), "chan len = ", len(urlChannel))
		time.Sleep(1 * time.Second)
		// 递归获取
		go Spy(url)
	}
}
