package history

import "go.mongodb.org/mongo-driver/bson/primitive"

type History struct {
	ID   primitive.ObjectID `bson:"_id"`
	Path string             `bson:"videoPath"`
}
