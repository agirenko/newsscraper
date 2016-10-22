package scraper

import (
	"fmt"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

type articleDescription struct {
	Url             string
	TitleHtml       string
	TitleText       string
	PublicationDate string
	Source          string
	SnippetHtml     string
	ArticleSnippet  string
}

func GetNews(url string) []articleDescription {
	fmt.Println("scraper started")
	var articles []articleDescription
	resp := GetResponse(url)
	defer resp.Body.Close()
	fmt.Println(resp.Header)
	doc, err := goquery.NewDocumentFromResponse(resp)
	if(err != nil) {
		panic(err)
	}
	fmt.Println(doc.Nodes)

	return articles
}

func GetResponse(url string) (resp *http.Response) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	return
}
