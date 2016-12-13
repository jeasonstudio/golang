package model

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/levigross/grequests"
	"strings"
	"time"
)

type Book struct {
	Name       string
	Position   string
	ReturnDate time.Time
}

func LoginAndGetBooks(user User) []Book {
	retry_times := 0
RETRY:
	s := grequests.NewSession(nil)
	s.Get("http://lib.ustb.edu.cn:8080/reader/login.php", nil)
	//获取验证码
	s.Get("http://lib.ustb.edu.cn:8080/reader/captcha.php?code=1", nil)
	option := &grequests.RequestOptions{
		Params: map[string]string{"number": user.Username,
			"passwd": user.Password, "captcha": "1",
			"select": "cert_no", "returnUrl": ""},
	}
	s.Post("http://lib.ustb.edu.cn:8080/reader/redr_verify.php", option)
	resp, _ := s.Get("http://lib.ustb.edu.cn:8080/reader/book_lst.php", nil)
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))

	books := []Book{}
	var is_null bool
	var is_find bool
	doc.Find("tr ").Each(func(i int, s *goquery.Selection) {
		is_find = true
		is_null = false
		var book Book
		s.Find(".whitetext").Each(func(i2 int, s2 *goquery.Selection) {
			switch i2 {
			case 1:
				book.Name = strings.Trim(strings.Split(s2.Text(), "/")[0], " ")
				if len(book.Name) == 0 {
					is_null = true
				}
			case 3:
				book.ReturnDate, _ = time.Parse("2006-01-02", strings.Trim(s2.Text(), " "))
			case 5:
				book.Position = strings.Trim(s2.Text(), " ")
				if !is_null {
					books = append(books, book)
				}
			}
		})
	})
	if !is_find && retry_times <= 3 {
		retry_times++
		fmt.Println("retry_times = ", retry_times)
		goto RETRY
	}

	return books
}
