package api

import (
	"context"
	"net/http"

	"github.com/aghyad-khlefawi/identity/pkg/jwthelper"
	"github.com/aghyad-khlefawi/identity/pkg/models"
	"github.com/aghyad-khlefawi/identity/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func HandleAuthenticateUser(c *gin.Context, api *ApiContext) {
	var request AuthenticateUserRequest
	err := c.BindJSON(&request)

	if err != nil {
		utils.HandleBadRequest("Invalid request structure", c)
		return
	}

	var user models.User
	err = api.sc.MongoClient.Database("identity").Collection("users").FindOne(context.TODO(), bson.M{"email": request.Email}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			unauthorized(c)
			return
		}
		utils.HandleServerError("Error while authenticating the user", err, c)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		unauthorized(c)
		return
	}

	token, err := jwthelper.GenerateJwt(user.Email)
	if err != nil {
		utils.HandleServerError("Failed to generate a token", err, c)
	}
  
	utils.WriteJsonContent(c, AuthenticateUserResponse{Token: token}, http.StatusOK)

	return
}

func unauthorized(c *gin.Context) {
	utils.WriteJsonContent(c, utils.Message{Msg: "Invalid user credentials"}, http.StatusUnauthorized)
}

type AuthenticateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticateUserResponse struct {
	Token string `json:"token"`
}
