package main

import (
	"github.com/agirenko/newsscraper/scraper"
)
const url = "http://news.google.com/news?hl=de&pz=1&ned=us&q=lottery+news"
func main() {
	scraper.GetNews(url);
}