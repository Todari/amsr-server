package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	//필수
	Id         primitive.ObjectID `bson:"_id"`
	Round      string             `bson:"round" validate:"required"`
	Privacy    bool               `bson:"privacy" validate:"required"`
	Name       string             `bson:"name" validate:"required"`
	Gender     bool               `bson:"gender" validate:"required"`
	Phone      string             `bson:"phone" validate:"required"`
	Age        string             `bson:"age" validate:"required"`
	Mbti       string             `bson:"mbti" validate:"required"`
	ChangeSeat bool               `bson:"changeSeat" validate:"required"`
	Bottles    int                `bson:"bottles" validate:"required"`
	Transfer   bool               `bson:"transfer" validate:"required"`

	//선택
	Invited string `bson:"invited"`
}
