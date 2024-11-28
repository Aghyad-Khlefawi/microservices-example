package api

import (
	"context"
	"net/http"

	"github.com/aghyad-khlefawi/identity/pkg/models"
	"github.com/aghyad-khlefawi/identity/utils"
	"github.com/gin-gonic/gin"
	"github.com/sethvargo/go-password/password"
)

func HandleCreateUser(c *gin.Context, api *ApiContext) {
	var request CreateUserRequest
	err:=c.BindJSON(&request)

	if err != nil {
		utils.HandleBadRequest("Invalid request structure",c)
		return
	}

	// Generate password if needed
	if request.GeneratePassword {
		res, err := password.Generate(14, 3, 2, false, true)
		if err != nil {
			utils.HandleServerError("Failed to generate password", err,c)
			return
		}
		request.Password = res

	} else {
		if len(request.Password) < 10 {
			utils.HandleBadRequest("Password must be at least 10 characters long",c)
			return
		}
	}

	// Hash the password
	hashedPassword, err := utils.HashPasswword(request.Password)
	if err != nil {
		utils.HandleServerError("Error creating the user", err,c)
		return
	}

	// Create the user record
	api.sc.MongoClient.Database("identity").Collection("users").InsertOne(context.TODO(), models.User{
		Email:    request.Email,
		Password: hashedPassword,
	})


  utils.WriteJsonContent(c, utils.Message{Msg: "User created"},http.StatusOK)
}



type CreateUserRequest struct {
	Email            string `json:"email"`
	Password         string `json:"password"`
	GeneratePassword bool   `json:"generatePassword"`
}
