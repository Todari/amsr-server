package services

import (
	"context"

	"github.com/Todari/amsr-server/configs"
	"github.com/Todari/amsr-server/models"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection = configs.GetCollection(configs.DB, "users")

func InsertOneUser(ctx context.Context, user models.User) (*mongo.InsertOneResult, error) {
	newUser := models.User{
		Id:         user.Id,
		Round:      user.Round,
		Privacy:    user.Privacy,
		Name:       user.Name,
		Gender:     user.Gender,
		Phone:      user.Phone,
		Age:        user.Age,
		Mbti:       user.Mbti,
		ChangeSeat: user.ChangeSeat,
		Bottles:    user.Bottles,
		Transfer:   user.Transfer,
		Invited:    user.Invited,
	}

	return userCollection.InsertOne(ctx, newUser)
}

// func FindManyUsers(ctx context.Context) (*mongo.Cursor, error) {
// 	return userCollection.Find(ctx, bson.M{})
// }

// func FindOneUser(ctx context.Context, match bson.M) *mongo.SingleResult {
// 	return userCollection.FindOne(ctx, match)
// }

// func UpdateOneUser(ctx context.Context, match bson.M, update bson.M) (*mongo.UpdateResult, error) {
// 	return userCollection.UpdateOne(ctx, match, bson.D{{"$set", update}})
// }
