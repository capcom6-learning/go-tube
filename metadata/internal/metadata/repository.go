package metadata

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MongoClientTimeout = 5
)

type MetadataRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func makeContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), MongoClientTimeout*time.Second)
}

func NewMetadataRepository(conn string, database string) (*MetadataRepository, error) {
	ctx, cancelFunc := makeContext()
	defer cancelFunc()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conn))
	if err != nil {
		return nil, err
	}

	return &MetadataRepository{
		client:     client,
		database:   client.Database(database),
		collection: client.Database(database).Collection("videos"),
	}, nil
}

func (r *MetadataRepository) SelectAll() ([]Metadata, error) {
	ctx, cancelFunc := makeContext()
	defer cancelFunc()

	cur, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var res []Metadata
	if err := cur.All(ctx, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (r *MetadataRepository) GetById(id string) (*Metadata, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx, cancelFunc := makeContext()
	defer cancelFunc()

	var metadata Metadata
	if err := r.collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&metadata); err != nil {
		return nil, err
	}

	return &metadata, nil
}

func (r *MetadataRepository) Create(metadata *Metadata) (*Metadata, error) {
	metadata.ID = primitive.NewObjectID()

	ctx, cancelFunc := makeContext()
	defer cancelFunc()

	_, err := r.collection.InsertOne(ctx, metadata)
	if err != nil {
		return nil, err
	}

	return metadata, nil
}

func (r *MetadataRepository) Disconnect() error {
	ctx, cancelFunc := makeContext()
	defer cancelFunc()

	return r.client.Disconnect(ctx)
}
