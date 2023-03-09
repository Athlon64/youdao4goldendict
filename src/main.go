package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

const url = "https://dict.youdao.com/search?q="

var word = os.Args[1]
var queryUrl = url + word

func Scrape() {
	// Request the HTML page.
	client := &http.Client{}

	req, err := http.NewRequest("GET", queryUrl, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".c-topbar-wrapper").Remove()
	doc.Find("#topImgAd").Remove()
	doc.Find(".ads").Remove()
	doc.Find("#c_footer").Remove()
	doc.Find("#results-contents").AppendHtml("<a href=" + queryUrl + ">在浏览器中查看</a>")
	fmt.Println(doc.Html())
}

func main() {
	Scrape()
}
