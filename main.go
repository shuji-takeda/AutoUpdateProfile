package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

const (
	displayContentNum = 5
	baseReadme        = `# 🐬 Welcome to my playground 🐬
	Thanks visit here.
	This trial is auto update your ReadMe every 1hour.
	1st.The Blog which I posted Blog in Qiita get by API.
	2nd.Update your ReadMe.md by github workFlow.
	
	## Recent posts - Blog 📜 
	`
)

var (
	expTitle = regexp.MustCompile("<item><title>.*</title>")
	expLink  = regexp.MustCompile("<link>https://qiita.com/takeshu17.*</link>")
)

type Item struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Url       string    `json:"url"`
}

func main() {
	resp, err := http.Get("http://qiita.com/api/v2/users/takeshu17/items?page=1&per_page=10")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data []Item
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}
	context := baseReadme
	fmt.Println(context)
	for _, item := range data {
		context += string("\n" + "🌵 " + "[" + item.Title + "]" + "(" + item.Url + ")" + "\n\n")
		fmt.Println(context)
	}
	err = os.WriteFile("README.md", []byte(context), 0666)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.ReadFile("README.md")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(file))
	// 	// resp, err := http.Get("http://qiita.com/api/v2/users/takeshu17/items?page=1&per_page=10")
	// 	// if err != nil {
	// 	// 	log.Fatal(err)
	// 	// }
	// 	// defer resp.Body.Close()
	// 	// bytes, err := ioutil.ReadAll(resp.Body)
	// 	// if err != nil {
	// 	// 	log.Fatal(err)
	// 	// }
	// 	// feed := string(bytes)
	// 	// titles := expTitle.FindAllString(feed, displayContentNum)
	// 	// links := expLink.FindAllString(feed, displayContentNum)
	// 	// fmt.Println(titles)
	// 	// readmeStr := baseReadme
	// 	// for i := 0; i < displayContentNum; i++ {
	// 	// 	t := titles[i]
	// 	// 	t = t[13 : len(t)-8]
	// 	// 	l := links[i]
	// 	// 	l = l[6 : len(l)-7]
	// 	// 	readmeStr += fmt.Sprintf("- [%s](%s)\n", t, l)
	// 	// }
	// 	// readmeFile, err := os.Create("README.md")
	// 	// if err != nil {
	// 	// 	log.Fatal(err)
	// 	// }
	// 	// defer func() { _ = readmeFile.Close() }()
	// 	// data := []byte(readmeStr)
	// 	// if _, err = readmeFile.Write(data); err != nil {
	// 	// 	log.Fatal(err)
	// 	// }
	// 	// os.Exit(0)
}
