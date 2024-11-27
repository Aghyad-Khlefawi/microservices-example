package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aghyad-khlefawi/identity/pkg/models"
	"github.com/aghyad-khlefawi/identity/utils"
	"github.com/sethvargo/go-password/password"
)

func HandleCreateUser(w http.ResponseWriter, r *http.Request, api *ApiContext) {
	request, err := utils.DeserializeJsonRequest[CreateUserRequest](r)
	if err != nil {
		utils.HandleBadRequest("Invalid request structure", w)
		return
	}

	// Generate password if needed
	if request.GeneratePassword {
		res, err := password.Generate(14, 3, 2, false, true)
		if err != nil {
			utils.HandleServerError("Failed to generate password", err, w)
			return
		}
		request.Password = res

	} else {
		if len(request.Password) < 10 {
			utils.HandleBadRequest("Password must be at least 10 characters long", w)
			return
		}
	}

	// Hash the password
	hashedPassword, err := utils.HashPasswword(request.Password)
	if err != nil {
		utils.HandleServerError("Error creating the user", err,w)
		return
	}

	// Create the user record
	api.sc.MongoClient.Database("identity").Collection("users").InsertOne(context.TODO(), models.User{
		Email:    request.Email,
		Password: hashedPassword,
	})

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "User created")
}



type CreateUserRequest struct {
	Email            string `json:"email"`
	Password         string `json:"password"`
	GeneratePassword bool   `json:"generatePassword"`
}
