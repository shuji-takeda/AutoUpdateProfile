package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	displayContentNum = 5
	baseReadme        = `# 🐬 Welcome to my Go playground 🐬
	🌈Thanks visit here🌈
	This golang trial is auto update your ReadMe every 1hour.
	1st.The Blog which I posted Blog in Qiita get by API.
	2nd.Update your ReadMe.md by github workFlow.

	## Recent posts - Blog 📜 
	`
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
	for _, item := range data {
		context += string("\n" + "🌵 " + "[" + item.Title + "]" + "(" + item.Url + ")" + "\n\n")
	}
	err = os.WriteFile("README.md", []byte(context), 0666)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(context)
}
