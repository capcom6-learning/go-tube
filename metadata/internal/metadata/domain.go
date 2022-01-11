package metadata

import "go.mongodb.org/mongo-driver/bson/primitive"

type Metadata struct {
	ID   primitive.ObjectID `bson:"_id" json:"id"`
	Path string             `bson:"videoPath" json:"videoPath"`
	Name string             `bson:"name" json:"name"`
}
