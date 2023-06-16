package repo

import (
	"api/core"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MongoTest struct {
	c *mongo.Client
}

func NewMongoTest(c *mongo.Client) *MongoTest {
	return &MongoTest{c: c}
}

func (m *MongoTest) Create(d *core.Data) error {
	collection := m.c.Database("testing").Collection("test")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, bson.D{
		{"c1", d.C1},
		{"c2", d.C2},
		{"c3", d.C3},
	})

	return err
}
