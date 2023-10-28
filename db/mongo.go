package db

import (
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://test:123@localhost:27017"

type MongoDBClient struct {
	Client *mongo.Client
}

func (c *MongoDBClient) Connect() error {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return err
	}

	c.Client = client

	return nil
}

func (c *MongoDBClient) Disconnect() error {
	if err := c.Client.Disconnect(context.TODO()); err != nil {
		return err
	}
	return nil
}

func (c MongoDBClient) GetArticleByID(id string) (*Article, error) {

	collection := c.Client.Database("blog").Collection("articles")

	article := &Article{}
	filter := bson.D{{"_id", id}}
	err := collection.FindOne(context.Background(), filter).Decode(article)
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (c MongoDBClient) SaveArticle(article *Article) error {
	collection := c.Client.Database("blog").Collection("articles")
	article.ID = uuid.New().String()
	_, err := collection.InsertOne(context.Background(), article)
	if err != nil {
		return err
	}

	return nil
}

func (c MongoDBClient) DeleteArticle(id string) error {
	collection := c.Client.Database("blog").Collection("articles")
	_, err := collection.DeleteOne(context.Background(), Article{ID: id})
	if err != nil {
		return err
	}

	return nil
}

func (c MongoDBClient) ListArticles() ([]*Article, error) {
	var articles []*Article

	collection := c.Client.Database("blog").Collection("articles")
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(context.Background(), &articles)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (c MongoDBClient) UpdateArticle(id string, article *Article) error {
	collection := c.Client.Database("blog").Collection("articles")
	filter := bson.D{{"_id", id}}
	_, err := collection.UpdateOne(context.Background(), filter, article)
	if err != nil {
		return err
	}

	return nil
}
