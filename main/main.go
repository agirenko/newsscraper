package main

import (
	"fmt"
	"github.com/agirenko/newsscraper/scraper"
)
const URL = "http://news.google.com/news?hl=de&pz=1&ned=us&q=lottery+news"
func main() {
	articles := scraper.GetNews(URL);
	fmt.Println(len(articles))
}