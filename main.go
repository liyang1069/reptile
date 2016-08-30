package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"reflect"
	"regexp"
	// "strings"
)

var (
	desUrl          = "http://www.qiushibaike.com/8hr/page/" //http://www.qiushibaike.com/8hr/page/2/
	ptnIndexItem    = regexp.MustCompile(`<div.*?author.*?>.*?<img.*?>.*?<h2>.*?</h2>.*?<div.*?content">(.*?)</div>(.*?)<i.*?class="number">(.*?)</i>.*?class="number">(.*?)</i>`)
	testItem        = regexp.MustCompile("<div class=\"content\">(.|\n)*?</div>(.|\n)*?<i class=\"number\">.*</i>(.|\n)*?<i class=\"number\">.*</i>")
	testRe          = regexp.MustCompile("a.*")
	ptnContentRough = regexp.MustCompile(`(?s).*<div class="artcontent">(.*)<div id="zhanwei">.*`)
	ptnBrTag        = regexp.MustCompile(`<br>`)
	ptnHTMLTag      = regexp.MustCompile(`(?s)</?.*?>`)
	ptnSpace        = regexp.MustCompile(`(^\s+)|( )`)
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

type IndexItem struct {
	url   string
	title string
}

func findIndex(content string) (index []IndexItem, err error) {
	fmt.Println("html_length:", len(content))
	matches := ptnIndexItem.FindAllStringSubmatch(content, 10000)
	fmt.Println("html_length:", len(matches))
	// test_string := "<div class=\"content\">\n你们碰到公司奇葩的规定是什么我不知道，但是我们公司：下午聚餐，不去罚款！！！\n</div><div class=\"content\">\n你们碰到公司奇葩的规定是什么我不知道，但是我们公司：下午聚餐，不去罚款！！！\n</div>"
	test_matches := testItem.FindAllStringSubmatch(content, 10000)
	fmt.Println("html_length:", len(test_matches))
	m := testRe.FindAllStringSubmatch("asdabcnbmabcff", 10000)
	fmt.Println("html_length:", len(m))
	index = make([]IndexItem, len(test_matches))
	for i, item := range test_matches {
		if i == 0 {
			fmt.Println(item[0], "aaaaaaaaaaaaaaaaaaaaaaaaaaa\n", item[1], "bbbbbbbbbbbbbbbbbbbbbb\n", item[2], "cccccccccccccccccccc\n")
		}
		index[i] = IndexItem{desUrl + item[1], item[2]}
	}
	return
}

func readContent(url string) (content string) {
	raw, statusCode := Get(url)
	if statusCode != 200 {
		fmt.Print("Fail to get the raw data from", url, "\n")
		return
	}

	match := ptnContentRough.FindStringSubmatch(raw)
	if match != nil {
		content = match[1]
	} else {
		return
	}

	content = ptnBrTag.ReplaceAllString(content, "\r\n")
	content = ptnHTMLTag.ReplaceAllString(content, "")
	content = ptnSpace.ReplaceAllString(content, "")
	return
}

func main() {
	fmt.Println(`Get index ...`)
	// str := "i am like cat"
	// fmt.Println(str)
	// sp := make([]string, 5)
	// j := 0
	// for _, ss := range str {
	// 	if ss != " " {
	// 		// if strings.EqualFold(ss, " ") {
	// 		sp[j] += ss
	// 	} else {
	// 		j++
	// 	}
	// }
	// fmt.Println(sp.join(" "))
	s, statusCode := Get(desUrl)
	if statusCode != 200 {
		fmt.Println("!!!!error!!!!  statusCode=%d", statusCode)
		return
	}
	index, _ := findIndex(s)
	ioutil.WriteFile("fileName.txt", []byte(s), 0644)

	fmt.Println(`Get contents and write to file ...`)
	for _, item := range index {
		fmt.Printf("Get content %s from %s and write to file.\n", item.title, item.url)
		// fileName := fmt.Sprintf("%s.txt", item.title)
		// content := readContent(item.url)
		// ioutil.WriteFile(fileName, []byte(content), 0644)
		// fmt.Printf("Finish writing to %s.\n", fileName)
	}
}
