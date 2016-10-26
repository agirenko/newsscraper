package mongodbnewsapi

import (
	"log"
	"gopkg.in/mgo.v2"
	"github.com/agirenko/newsscraper/scraper"
	"time"
	"gopkg.in/mgo.v2/bson"
)

func SaveArticlesToMongoDb(articles []scraper.ArticleDescription) {
	session, err := mgo.Dial("mongodb://dbusername:Abc%21_587@ds019076.mlab.com:19076/agirenko")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	// http://godoc.org/labix.org/v2/mgo#Session.SetMode
	session.SetMode(mgo.Monotonic, true)
	// get collection
	collection := session.DB("agirenko").C("googlenews")
	// insert new articles one by one to collection
	for i := 0; i < len(articles); i++ {
		foundNumber, err := collection.Find(bson.M{"articleUrl": articles[i].ArticleUrl}).Count()
		if err != nil {
			log.Fatal(err)
		}
		if foundNumber == 0 {
			articles[i].CreatedAt = time.Now().UTC()
			articles[i].UpdatedAt = time.Now().UTC()
			err = collection.Insert(&articles[i])
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}


