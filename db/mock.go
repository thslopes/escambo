package db

// Article fixture data
var articles = []*Article{
	{ID: "1", UserID: 100, Title: "Hi", Slug: "hi"},
	{ID: "2", UserID: 200, Title: "sup", Slug: "sup"},
	{ID: "3", UserID: 300, Title: "alo", Slug: "alo"},
	{ID: "4", UserID: 400, Title: "bonjour", Slug: "bonjour"},
	{ID: "5", UserID: 500, Title: "whats up", Slug: "whats-up"},
}

func Mock() {
	dbClient := MongoDBClient{}
	err := dbClient.Connect()
	if err != nil {
		panic(err)
	}

	defer func() {
		err := dbClient.Disconnect()
		if err != nil {
			panic(err)
		}
	}()

	for _, article := range articles {
		err := dbClient.SaveArticle(article)
		if err != nil {
			panic(err)
		}
	}

}
