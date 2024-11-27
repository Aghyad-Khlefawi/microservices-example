package api

import (
	"context"
	"net/http"

	"github.com/aghyad-khlefawi/identity/pkg/models"
	"github.com/aghyad-khlefawi/identity/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func HandleAuthenticateUser(w http.ResponseWriter, r *http.Request, api *ApiContext) {
	request, err := utils.DeserializeJsonRequest[AuthenticateUserRequest](r)
	if err != nil {
		utils.HandleBadRequest("Invalid request structure", w)
		return
	}

	var user models.User
	err = api.sc.MongoClient.Database("identity").Collection("users").FindOne(context.TODO(),bson.M {"email": request.Email}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			unauthorized(w)
			return
		}
		utils.HandleServerError("Error while authenticating the user", err, w)
		return
	}

	if err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(request.Password)); err!=nil {
		unauthorized(w)
		return
	}
	utils.WriteJsonContent(w, utils.Message{Msg: "Authenticated"},http.StatusOK)
	return
}

func unauthorized(w http.ResponseWriter) {
	utils.WriteJsonContent(w, utils.Message{Msg: "Invalid user credentials"}, http.StatusUnauthorized)
}

type AuthenticateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
