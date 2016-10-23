package scraper

import (
	"fmt"
	"net/http"
	"golang.org/x/net/html"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"os"
	//"path/filepath"
	//"go/doc"
)

func GetNews(url string) []articleDescription {
	fmt.Println("scraper started")
	var articles []articleDescription
	//resp := GetResponse(url)
	//defer resp.Body.Close()
	//printBody(resp)
	//fmt.Println(resp.Header)
	//doc, err := goquery.NewDocumentFromResponse(resp)
	//if (err != nil) {
	//	panic(err)
	//}

	doc := loadDoc("page.html")
	//var articleNodes = HtmlDoc.DocumentNode.SelectNodes("//td[@class='esc-layout-article-cell']");
	articleNodes :=doc.Find("td.esc-layout-article-cell")
	//firstNode := articleNodes.Nodes[0]
	log.Println(articleNodes.Text())
	articleNodes.Each(func(_ int, s *goquery.Selection) {

	})
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

func loadDoc(page string) *goquery.Document {
	var f *os.File
	var e error
	if f, e = os.Open(fmt.Sprintf("D:/temp/testdata/%s", page)); e != nil {
		panic(e.Error())
	}
	defer f.Close()

	var node *html.Node
	if node, e = html.Parse(f); e != nil {
		panic(e.Error())
	}
	return goquery.NewDocumentFromNode(node)
}



