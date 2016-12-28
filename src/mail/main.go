package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/smtp"
	"strings"
	"time"
)

var all = [30]string{
	"你若继续做过去的事，你将一直是过去的你。以自我为中心的人所拆毁的，以他人为中心的人可以重建。忍耐不仅是忍受困难，更是将忍耐化为能力。天够黑的时候，才能看见星星。",
	"很多事情努力了未必有结果，但是不努力却什么改变也没有。",
	"留一片空白，随时浓墨重彩。",
	"你认为自己行就一定行，每天要不断向自己重复。",
	"你是唯一的，你是非常独特的，你就是你生命中的第一名。",
	"只要充分相信自己，没有什么困难可以足够持久。",
	"不要沉沦，在任何环境中你都可以选择奋起。",
	"一个人只有投身于伟大的时代洪流中，他的生命才会闪耀出光彩。",
	"要培养各方面的能力，包括承受悲惨命运的能力。",
	"只要你不认输，就有机会。",
	"要战胜恐惧，而不是退缩。",
	"失败者任其失败，成功者创造成功。",
	"胜利，是属于最坚韧的人。",
	"只有收获，才能检验耕耘的意义；只有贡献，方可衡量人生的价值。",
	"我自信，故我成功；我行，我一定能行。",
	"人活着要呼吸。呼者，出一口气；吸者，争一口气。",
	"相信所有的汗水与眼泪，最后会化成一篇山花烂漫。",
	"一百次心动不如一次行动。",
	"一时的忍耐是为了更广阔的自由，一时的纪律约束是为了更大的成功。",
	"追求让人充实，分享让人快乐。",
	"发奋图强的励志的句无论做什么，记得是为自己而做，那就毫无怨言。",
	"人生短短数十载，最要紧是证明自己，不是讨好他人。",
	"命运不是一个机遇的问题，而是一个选择问题；它不是我们要等待的东西，而是我们要实现的东西。",
	"任何业绩的质变，都来自于量变的积累。",
	"每一个发奋努力的背后，必有加倍的赏赐。",
	"不要说你不会做！你是个人你就会做。",
	"天赐我一双翅膀，就应该展翅翱翔，满天乌云又能怎样，穿越过就是阳光。",
	"我们的目的是什么？是胜利！不惜一切代价争取胜利。",
	"天助自助者，你要你就能。",
	"相信自己，你能作茧自缚，就能破茧成蝶。",
}

func SendToMAIL(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func sendReady(innerHTML string) {
	user := "mailbyjeason@jeasonstudio.cn"
	password := "Admin12345"
	host := "smtp.exmail.qq.com:25"
	to := "748807384@qq.com;me@jeasonstudio.cn"

	subject := "Jeason Studio"

	body := innerHTML
	fmt.Println("send email")
	err := SendToMAIL(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}

}

func getCent() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(r.Intn(30))
	// }
	fmt.Println(all[r.Intn(30)])
	tagInner := all[r.Intn(30)]

	tagH := "<html><body><h2>" + tagInner + "加油啊！💪</h2><p>————来自爱你的彤哥</p></body></html>"

	sendReady(tagH)
}

func main() {
	// thisCookie := "JSESSIONID=" + tagCookie
	tagLoginURL := "http://jeasonstudio.github.io/"

	client := &http.Client{}
	req, _ := http.NewRequest("GET", tagLoginURL, nil)

	// req.Header.Set("Cookie", thisCookie)

	resp, err := client.Do(req) //发送
	defer resp.Body.Close()     //一定要关闭resp.Body
	// data, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(data), err)
	if err != nil {
		return
	}
	// sendReady(string(data))

	ticker := time.NewTicker(time.Minute * 30)
	for _ = range ticker.C {
		eMailTime := time.Now().Format("15:04")
		fmt.Println(eMailTime)
		// fmt.Println(eMailTime > "12:03")

		// getCent()

		if eMailTime > "19:44" && eMailTime < "20:17" {
			fmt.Println("true")
			getCent()
		}
	}

}
