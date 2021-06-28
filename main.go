package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

const (
	displayContentNum = 5
	baseReadme        = `# Welcome to my playground🐬🐳

	
	## Recent posts - Blog 📜
	`
)

var (
	expTitle = regexp.MustCompile("<item><title>.*</title>")
	expLink  = regexp.MustCompile("<link>https://qiita.com/takeshu17.*</link>")
)

func main() {
	response, err := http.Get("https://shuji-takeda.github.io/tech-blog/index.xml")
	if err != nil {
		panic(err)
	}
	defer func() { _ = response.Body.Close() }()
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	feed := string(bytes)

	titles := expTitle.FindAllString(feed, displayContentNum)
	links := expLink.FindAllString(feed, displayContentNum)

	readmeStr := baseReadme
	for i := 0; i < displayContentNum; i++ {
		t := titles[i]
		t = t[13 : len(t)-8]
		l := links[i]
		l = l[6 : len(l)-7]
		readmeStr += fmt.Sprintf("- [%s](%s)\n", t, l)
	}

	readmeFile, err := os.Create("README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = readmeFile.Close() }()

	data := []byte(readmeStr)
	if _, err = readmeFile.Write(data); err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}

// // package main

// // import (
// // 	"bytes"
// // 	"flag"
// // 	"fmt"
// // 	"os"
// // 	"regexp"
// // 	"time"

// // 	"github.com/ikawaha/feedsnippet/feeder"
// // )

// // const (
// // 	startTag   = `<!--[START github.com/ikawaha/feedsnippet]-->`
// // 	endTag     = `<!--[END github.com/ikawaha/feedsnippet]-->`
// // 	timestamp  = `(?:<!--\[(?:.+?)]-->)?\n`
// // 	timestampP = "<!--[%s]-->\n"
// // )

// // var (
// // 	snippetField = regexp.MustCompile(`(?s)` + regexp.QuoteMeta(startTag) + timestamp + `(.*)` + regexp.QuoteMeta(endTag))
// // )

// // type option struct {
// // 	config  string
// // 	target  string
// // 	verbose bool
// // 	debug   bool
// // 	diff    bool
// // }

// // func main() {
// // 	var opt option
// // 	flag.StringVar(&opt.config, "config", "", "config file")
// // 	flag.StringVar(&opt.target, "file", "", "target file (optional)")
// // 	flag.BoolVar(&opt.diff, "diff", false, "replace snippets only when there is a difference (optional)")
// // 	flag.BoolVar(&opt.verbose, "verbose", false, "print snippets to stdout (optional)")
// // 	flag.BoolVar(&opt.debug, "debug", false, "print raw fees for debug (optional)")
// // 	flag.Parse()
// // 	if err := run(opt); err != nil {
// // 		_, _ = fmt.Fprintln(os.Stderr, err)
// // 		os.Exit(1)
// // 	}
// // }

// // func run(opt option) error {
// // 	config, err := feeder.NewConfigFile(opt.config)
// // 	if err != nil {
// // 		return err
// // 	}
// // 	var b bytes.Buffer
// // 	for _, c := range config {
// // 		f, err := feeder.NewFeeder(feeder.DebugOpt(opt.debug), feeder.FilterOpt(c.Filters()...))
// // 		if err != nil {
// // 			return err
// // 		}
// // 		feeds, err := f.Feeds(c.URLs...)
// // 		if err != nil {
// // 			return fmt.Errorf("feeds error: %w", err)
// // 		}
// // 		tmpl := feeder.DefaultFeedT
// // 		if c.Template != "" {
// // 			tmpl = c.Template
// // 		}
// // 		s, err := feeder.ExecuteTmpl(feeds, tmpl)
// // 		if err != nil {
// // 			return fmt.Errorf("%v\n%s", err, tmpl)
// // 		}
// // 		b.WriteString(s)
// // 	}
// // 	if err := writeSnippet(opt, b.Bytes()); err != nil {
// // 		return err
// // 	}
// // 	return nil
// // }

// // func writeSnippet(opt option, snippet []byte) error {
// // 	if opt.target == "" || opt.verbose {
// // 		fmt.Println(string(snippet))
// // 		if opt.target == "" {
// // 			return nil
// // 		}
// // 	}
// // 	body, err := os.ReadFile(opt.target)
// // 	if err != nil {
// // 		return err
// // 	}
// // 	if opt.diff {
// // 		eq, err := snippetEqual(body, snippet)
// // 		if err != nil {
// // 			return err
// // 		}
// // 		if eq {
// // 			return nil
// // 		}
// // 	}
// // 	// creat tagged snippet
// // 	b := bytes.NewBufferString(startTag)
// // 	b.WriteString(fmt.Sprintf(timestampP, time.Now().Format(time.RFC3339)))
// // 	b.Write(snippet)
// // 	b.WriteString(endTag)
// // 	body = snippetField.ReplaceAll(body, b.Bytes())
// // 	if err := os.WriteFile(opt.target, body, 0666); err != nil {
// // 		return err
// // 	}
// // 	return nil
// // }

// // func snippetEqual(tagged, snippet []byte) (bool, error) {
// // 	matched := snippetField.FindSubmatch(tagged)
// // 	if len(matched) != 2 {
// // 		return false, fmt.Errorf("no snipppet field")
// // 	}
// // 	return bytes.Equal(matched[1], snippet), nil
// // }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"time"
// )

// type Item struct {
// 	Title     string
// 	CreatedAt time.Time
// 	Url       string
// }

// func main() {
// 	resp, err := http.Get("http://qiita.com/api/v2/users/takeshu17/items?page=1&per_page=10")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var data []Item // nil slice
// 	// data := make([]Item, 0) のように要素数0の slice としても良い

// 	if err := json.Unmarshal(body, &data); err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(data)

// 	for _, item := range data {
// 		fmt.Printf("%s %s\n %s\n", item.CreatedAt, item.Title, item.Url)
// 	}
// }
