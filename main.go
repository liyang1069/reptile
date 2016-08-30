package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	desUrl  = "http://www.qiushibaike.com/hot/page/" //http://www.qiushibaike.com/8hr/page/2/
	ptnItem = regexp.MustCompile("<div class=\"content\">(.|\n)*?</div>(.|\n)*?<i class=\"number\">.*</i>(.|\n)*?<i class=\"number\">.*</i>")
)

//获取html body and status code
func Get(url string) (content string, statusCode int) {
	resp, err1 := http.Get(url)
	// fmt.Println("resp:", resp)
	// fmt.Println("err1:", err1)
	if err1 != nil {
		statusCode = -100
		return
	}
	defer resp.Body.Close()
	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		statusCode = -200
		return
	}
	statusCode = resp.StatusCode
	content = string(data)
	return
}

type JokeItem struct {
	content string
	vote    int
	comment int
}

// func findIndex(content string) (jokes []JokeItem, err error) {
func findIndex(content string) {
	ptn_matches := ptnItem.FindAllStringSubmatch(content, 10000)
	// fmt.Println("html_length:", len(ptn_matches))
	jokes := make([]JokeItem, len(ptn_matches))
	for i, item := range ptn_matches {
		joke_str := item[0]
		joke_str = strings.Replace(joke_str, "<div class=\"content\">", "", -1)
		joke_str = strings.Replace(joke_str, "\n", "", -1)
		joke_str = strings.Replace(joke_str, "<br>", "\n", -1)
		array := strings.Split(joke_str, "<i class=\"number\">")
		content := strings.Split(array[0], "</div>")[0]
		vote, err_vote := strconv.Atoi(strings.Split(array[1], "</i>")[0])
		comment, err_com := strconv.Atoi(strings.Split(array[2], "</i>")[0])
		if err_vote != nil {
			vote = 0
		}
		if err_com != nil {
			comment = 0
		}
		if vote > 5000 {
			fmt.Println(content)
			fmt.Println("=========================")
		}
		jokes[i] = JokeItem{content, vote, comment}
	}
	return
}

func get_jokes(url string) {
	s, statusCode := Get(url)
	if statusCode != 200 {
		fmt.Println("!!!!error!!!!  statusCode=%d  url=%s", statusCode, url)
		return
	}
	findIndex(s)
	// index, _ := findIndex(s)
	// len(index)
	// _ == nil
}

func main() {
	fmt.Println(`Get index ...`)
	// index, _ := findIndex(s)
	// ioutil.WriteFile("fileName.txt", []byte(s), 0644)

	// fmt.Println(len(index))
	// fmt.Println(`Get contents and write to file ...`)
	for i := 1; i <= 35; i++ {
		go get_jokes(desUrl + strconv.Itoa(i) + "/")
	}
	time.Sleep(1000 * time.Millisecond)
}
