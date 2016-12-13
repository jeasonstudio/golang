package main

import (
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/jinzhu/gorm"
	"github.com/nladuo/ustb_book_reminder/model"
	"strconv"
	"time"
)

//mysql配置
const (
	DB_USER   = "root"
	DB_PASSWD = "root"
	DB_HOST   = "localhost"
	DB_PORT   = "3306"
	DBNAME    = "ustb_book_reminder"
)

func main() {
	db, err := gorm.Open("mysql", DB_USER+":"+DB_PASSWD+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DBNAME+"?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}

	users := []model.User{}
	db.Find(&users)
	now := time.Now()

	for _, user := range users {
		fmt.Println(user.Username)
		books := model.LoginAndGetBooks(user)
		need_alert := false
		alert_msg := ""
		for _, book := range books {
			left_day := float64(book.ReturnDate.Unix()-now.Unix()) / 60 / 60 / 24
			if left_day < 2.0 {
				need_alert = true
				//每天早上七点半提醒
				if left_day >= 0 {
					left_day_str := strconv.FormatInt(int64(left_day), 10)
					alert_msg += "《" + book.Name + "》[" + book.Position + "]" + " 剩余 " + left_day_str + " 天到期\n"
				} else {
					left_day_str := strconv.FormatInt(int64(-left_day)+1, 10)
					alert_msg += "《" + book.Name + "》[" + book.Position + "]" + " 已过期 " + left_day_str + " 天\n"
				}
			}
			fmt.Println(book.Name, left_day)
		}

		if need_alert {
			fmt.Println(alert_msg)
			email := model.NewEmail(user.Mail, alert_msg)
			err := model.SendEmail(email)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		fmt.Println("")
	}
}
