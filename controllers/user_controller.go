package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/Todari/amsr-server/models"
	"github.com/Todari/amsr-server/services"
	"github.com/Todari/amsr-server/structs"
	"github.com/gin-gonic/gin"
)

func CreateUser() gin.HandlerFunc {
	return func(ctx_ *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var newUser models.User

		if err := ctx_.BindJSON(&newUser); err != nil {
			ctx_.JSON(
				http.StatusBadRequest,
				structs.HttpResponse{
					Success: false,
					Data: map[string]interface{}{
						"type":   "Bind user error",
						"result": err.Error(),
					},
				},
			)
			return
		}

		result, err := services.InsertOneUser(ctx, newUser)
		if err != nil {
			ctx_.JSON(
				http.StatusInternalServerError,
				structs.HttpResponse{
					Success: false,
					Data: map[string]interface{}{
						"type":   "InsertOneProperty error",
						"result": err.Error(),
					},
				},
			)
			return
		}

		ctx_.JSON(
			http.StatusCreated,
			structs.HttpResponse{
				Success: true,
				Data: map[string]interface{}{
					"result": result.InsertedID,
				},
			},
		)
	}
}
