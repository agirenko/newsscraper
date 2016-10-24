package mongodbnewsapi
// package contains methods to save news records to MongoDb
//https://godoc.org/gopkg.in/mgo.v2

//private static void SaveArticlesToMongoDb(List<ArticleDescription> articles, string location)
//{
//
//_client_2 = new MongoClient(MongoDb_2_ConnectionString);
//_database_2 = _client_2.GetDatabase(RemoteMongoDb_2_DatabaseName);
//
//var collection_2 = _database_2.GetCollection<BsonDocument>(MongoDbCollectionName);
//
//foreach (ArticleDescription article in articles)
//{
//try
//{
//
//bool doesResultExist_2 = DoesRecordExistInMongoDb(collection_2, article.Url, location);
//Dictionary<string, object> results = null;
//string resultsJson = null;
//BsonDocument resultsBson = null;
//
//if (!doesResultExist_2) // this article is absent in mongodb database. add it
//{
//if (results == null)
//{
//results = MapArticleToResultsDictionary(article, location);
//// convert results to JSON first
//resultsJson = JsonConvert.SerializeObject(results);
//// convert json to BsonDocument
//resultsBson = MongoDB.Bson.Serialization.BsonSerializer.Deserialize<BsonDocument>(resultsJson);
//}
//Task insertTwo = collection_2.InsertOneAsync(resultsBson);
//insertTwo.Wait();
//}
//
//}
//catch (Exception ex)
//{
//int i = 0;
//// just catch for debugging
//}
//}
//}

// Method which was used one time to clear wrong record on parse.com database
//public static async Task DeleteUndefinedRecords(string parseClass)
//{
//    Task<IEnumerable<ParseObject>> query = ParseObject.GetQuery(parseClass).WhereDoesNotExist("Source").Limit(1000).FindAsync();
//    var result = query.Result;
//    foreach (ParseObject po in result)
//    {
//        var deleteTask = po.DeleteAsync();
//        deleteTask.Wait();
//    }
//}

//public static bool DoesRecordExistInMongoDb(IMongoCollection<BsonDocument> collection, string url, string location)
//{
//try
//{
//var builder = Builders<BsonDocument>.Filter;
//var filter = builder.Eq("articleUrl", url) & builder.Eq("location", location);
//return (collection.Count(filter) > 0);
//}
//catch (Exception ex)
//{
//return false;
//}
//}
//

//private static Dictionary<string, object> MapArticleToResultsDictionary(ArticleDescription article, string location)
//{
//var results = new Dictionary<string, object>();
//results["createdAt"] = DateTime.UtcNow;
//results["updatedAt"] = DateTime.UtcNow;
//results["articleUrl"] = article.Url;
//results["date"] = article.PublicationDate;
//results["Source"] = article.Source;
//results["articleTitle"] = article.TitleText;
//results["TitleHtml"] = article.TitleHtml;
//results["snippetHtml"] = article.SnippetHtml;
//results["articleSnippet"] = article.ArticleSnippet;
//results["location"] = location;
//return results;
//}


