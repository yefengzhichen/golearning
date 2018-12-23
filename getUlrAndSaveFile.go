package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	cookie    string = "user_trace_token=20181027095150-ea0a2686-702d-495c-bb1e-b5f2754a713e; _ga=GA1.2.1816089271.1540605339; LGUID=20181027095555-70ce349c-d98b-11e8-822b-5254005c3644; JSESSIONID=ABAAABAAAGFABEF997290EB6B43F5201A49ACC81F1CAF2B; _gid=GA1.2.1090287615.1545466788; Hm_lvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1545466789; TG-TRACK-CODE=index_navigation; LGSID=20181222165116-be45ad74-05c6-11e9-88d3-525400f775ce; PRE_UTM=; PRE_HOST=; PRE_SITE=https%3A%2F%2Fwww.lagou.com%2F; PRE_LAND=https%3A%2F%2Fwww.lagou.com%2Fzhaopin%2Fgo%2F%3FlabelWords%3Dlabel; index_location_city=%E6%B7%B1%E5%9C%B3; SEARCH_ID=570479f2fcb34ba9b20b5e613f93374b; _gat=1; Hm_lpvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1545469731; LGRID=20181222170850-323a1866-05c9-11e9-9f66-5254005c3644"
	userAgent string = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.26 Safari/537.36 Core/1.63.6788.400 QQBrowser/10.3.2714.400"
)

func main() {
	//url := "https://www.baidu.com/"
	baseUrl := "https://www.lagou.com/"
	lang := "Go"
	city := "成都"
	url := baseUrl + "jobs/list_" + lang + "?"
	url += "city=" + city

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	//request.Header.Add("Cookie", cookie)
	request.Header.Add("User-Agent", userAgent)
	rsp, err := client.Do(request)

	if err != nil {
		fmt.Println("error read", err)
		return
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	fmt.Println(string(body))
	fout, err := os.Create("index.html")
	if err != nil {
		fmt.Println("create file error", err)
		return
	}
	defer fout.Close()
	fout.Write(body)
	fout.Sync()
}
