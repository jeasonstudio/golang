package teachSystem

func main()  {
    tagLoginURL := "http://elearning.ustb.edu.cn/choose_courses/choosecourse/normalChooseCourse_normalRequired_loadPreNormalAccordByKchRequiredCourses.action?kch=2050414&_dc=1483270853482&limit=5000&start=0&uid=41524122"

	client := &http.Client{}
	req, _ := http.NewRequest("GET", tagLoginURL, nil)

	// req.Header.Set("Cookie", thisCookie)

	resp, err := client.Do(req) //发送
	defer resp.Body.Close()     //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data), err)
	if err != nil {
		return
	}
}