package scraper

import (
	"fmt"
	"net/http"
	"golang.org/x/net/html"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"os"
	"strings"
	"github.com/agirenko/newsscraper/helper"
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
	articleNodes := doc.Find("td.esc-layout-article-cell")
	//firstNode := articleNodes.Nodes[0]
	log.Println(articleNodes.Text())
	log.Println("============================================")
	articleNodes.Each(func(_ int, s *goquery.Selection) {
		url, titleText, titleHtml := getUrlAndTitle(s)
		pubDate := getPublicationDate(s)
		article := articleDescription{Url:url, TitleText:titleText, TitleHtml:titleHtml, PublicationDate:pubDate}
		fmt.Println("-----------------------")
		fmt.Println(article.Url)
		fmt.Println(article.TitleText)
		fmt.Println(article.TitleHtml)
		fmt.Println(article.PublicationDate)
		articles = append(articles, article)
	})
	fmt.Println("-----------------------")
	log.Println("============================================")
	return articles

}

func getUrlAndTitle(s *goquery.Selection) (val string, titleText string, titleHtml string) {
	aLink := s.Find("a").First()
	val, _ = aLink.Attr("href")
	titleText = aLink.Text()
	span := aLink.Find("span").First()
	titleHtml, _ = span.Html()
	return
}

func getPublicationDate(s *goquery.Selection) string {
	dtContent := s.Find("span.al-attribution-timestamp").First().Text()
	if strings.Contains(dtContent, "hour") {
		return helper.CountPublicationDateByHoursAgo(dtContent)
	}
	if strings.Contains(dtContent, "minute") {
		return helper.CountPublicationDateByMinutesAgo(dtContent)
	}
	return dtContent
}

func getSource(s *goquery.Selection) string {
	return ""
}

func getSnippetHtml(s *goquery.Selection) string {
	return ""
}

func getArticleSnippet(s *goquery.Selection) string {
	return ""
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



