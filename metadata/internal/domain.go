package internal

import "go.mongodb.org/mongo-driver/bson/primitive"

type Metadata struct {
	ID   primitive.ObjectID `bson:"_id" json:"id"`
	Path string             `bson:"videoPath" json:"videoPath"`
}
