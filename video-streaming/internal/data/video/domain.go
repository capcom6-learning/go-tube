package video

import "go.mongodb.org/mongo-driver/bson/primitive"

type Video struct {
	ID   primitive.ObjectID `bson:"_id"`
	Path string             `bson:"videoPath"`
}
