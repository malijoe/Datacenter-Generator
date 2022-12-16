package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectTimeout  = 30 * time.Second
	maxConnIdleTime = 3 * time.Minute
	minPoolSize     = 20
	maxPoolSize     = 300
)

type Config struct {
	URI      string
	User     string
	Password string
	DbName   string
}

func NewMongoDBConn(ctx context.Context, cfg *Config) (*mongo.Client, error) {
	opts := options.Client().
		ApplyURI(cfg.URI).
		SetAuth(options.Credential{Username: cfg.User, Password: cfg.Password}).
		SetConnectTimeout(connectTimeout).
		SetMaxConnIdleTime(maxConnIdleTime).
		SetMinPoolSize(minPoolSize).
		SetMaxPoolSize(maxPoolSize)

	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, err
	}

	if err := client.Connect(ctx); err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}

func ensureMongoIndices(database *mongo.Database) {
	ensureIndex(database, "datacenter", "site", true, nil)
}

func ensureIndex(database *mongo.Database, collectionName string, field string, unique bool, partialFilterExpression any) bool {
	createIndexOpts := &options.IndexOptions{Unique: &unique}
	if partialFilterExpression != nil {
		createIndexOpts.SetPartialFilterExpression(partialFilterExpression)
	}

	mod := mongo.IndexModel{
		Keys:    bson.D{{Key: field, Value: 1}},
		Options: createIndexOpts,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := database.Collection(collectionName)
	_, err := collection.Indexes().CreateOne(ctx, mod)
	if err != nil {
		return false
	}

	return true
}

func ensureCompoundIndex(database *mongo.Database, collectionName string) bool {
	collection := database.Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	compoundIndices := compoundIndices()

	compoundIndex, ok := compoundIndices[collectionName]
	if !ok {
		return false
	}

	_, err := collection.Indexes().CreateMany(ctx, compoundIndex)
	if err != nil {
		return false
	}

	return true
}

func compoundIndices() map[string][]mongo.IndexModel {
	compoundIndices := map[string][]mongo.IndexModel{
		"pod": {
			{
				Keys: bson.D{
					{Key: "datacenterId", Value: 1},
					{Key: "name", Value: 1},
				},
				Options: options.Index().SetUnique(true),
			},
		},
	}

	return compoundIndices
}
