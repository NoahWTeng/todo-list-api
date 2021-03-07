package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"sync"
	"time"
)

func NewMongodbConnection(cfg *Config) *Handler {
	mongodbHandler := &MongoConnector{
		cfg: cfg,
	}

	err := mongodbHandler.connect()

	if err != nil {
		log.Fatal(err)
	}

	base := &Handler{
		Config:  cfg,
		Methods: mongodbHandler,
	}

	return base

}

func (c *MongoConnector) connect() error {
	var (
		connectOnce sync.Once
		err         error
		client      *mongo.Client
		url         string
	)
	connectOnce.Do(func() {
		url = formatUrl(c.cfg)
		client, err = mongo.NewClient(options.Client().ApplyURI(url))
		if err != nil {
			log.Fatal(err)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.cfg.DialTimeOut))
		defer cancel()
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
			return
		}
	})
	if err != nil {
		return err
	}
	c.client = client
	c.db = c.client.Database(c.cfg.DatabaseName)

	fmt.Printf("Database success connected! %s\n", url)
	return nil
}

func (c *MongoConnector) DB(ctx context.Context) *mongo.Database {
	var rp readpref.ReadPref
	err := c.client.Ping(ctx, &rp)
	if err != nil {
		log.Fatal(err)
	}
	return c.db
}
func (c *MongoConnector) Client() *mongo.Client {
	return c.client
}

func (c *MongoConnector) Config() Config {
	return *c.cfg
}
