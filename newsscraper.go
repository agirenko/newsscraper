package main

import (
	"fmt"
	"github.com/agirenko/newsscraper/scraper"
	"github.com/agirenko/newsscraper/mongodbnewsapi"
)
const URL = "http://news.google.com/news?hl=de&pz=1&ned=us&q=lottery+news"
func main() {
	articles := scraper.GetNews(URL);
	mongodbnewsapi.SaveArticlesToMongoDb(articles)
	fmt.Println(len(articles))
}