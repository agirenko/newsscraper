package scraper

import (
	"fmt"
	"net/http"
	"golang.org/x/net/html"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"os"
)

func GetNews(url string) []articleDescription {
	fmt.Println("scraper started")
	var articles []articleDescription
	resp := GetResponse(url)
	defer resp.Body.Close()
	//printBody(resp)
	fmt.Println(resp.Header)
	doc, err := goquery.NewDocumentFromResponse(resp)
	if (err != nil) {
		panic(err)
	}
	//var articleNodes = HtmlDoc.DocumentNode.SelectNodes("//td[@class='esc-layout-article-cell']");
	articleNodes :=doc.Find("td.esc-layout-article-cell")
	var firstNode html.Node = articleNodes.Nodes[0]
	fmt.Println(firstNode)
	fmt.Println(articleNodes)


	return articles
}

func GetResponse(url string) (resp *http.Response) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	return
}

func printBody(resp *http.Response) {
	_, err := io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}


