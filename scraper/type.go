package scraper

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type ArticleDescription struct {
	Id              bson.ObjectId 	`bson:"_id,omitempty"`
	CreatedAt       time.Time	`bson:"createdAt"`
	UpdatedAt       time.Time	`bson:"updatedAt"`
	ArticleUrl      string		`bson:"articleUrl"`
	PublicationDate string	`bson:"date"`
	Source          string		`bson:"source"`
	TitleText       string		`bson:"articleTitle"`
	TitleHtml       string		`bson:"titleHtml"`
	ArticleSnippet  string		`bson:"articleSnippet"`
	SnippetHtml     string	`bson:"snippetHtml"`
	Location        string		`bson:"location"`
}
