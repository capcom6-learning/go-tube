package history

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MongoClientTimeout = 5
)

type HistoryRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func makeContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), MongoClientTimeout*time.Second)
}

func NewHistoryRepository(conn string, database string) (*HistoryRepository, error) {
	ctx, cancelFunc := makeContext()
	defer cancelFunc()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conn))
	if err != nil {
		return nil, err
	}

	return &HistoryRepository{
		client:     client,
		database:   client.Database(database),
		collection: client.Database(database).Collection("videos"),
	}, nil
}

func (r *HistoryRepository) Insert(path string) error {
	ctx, cancelFunc := makeContext()
	defer cancelFunc()

	record := History{
		ID:      primitive.NewObjectID(),
		Path:    path,
		Watched: time.Now(),
	}

	_, err := r.collection.InsertOne(ctx, record)

	return err
}

func (r *HistoryRepository) Disconnect() error {
	ctx, cancelFunc := makeContext()
	defer cancelFunc()

	return r.client.Disconnect(ctx)
}
