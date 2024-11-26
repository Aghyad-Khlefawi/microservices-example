package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aghyad-khlefawi/identity/pkg/models"
	"github.com/aghyad-khlefawi/identity/utils"
	"github.com/sethvargo/go-password/password"
)

func HandleCreateUser(w http.ResponseWriter, r *http.Request, api *ApiContext) {
	decoder := json.NewDecoder(r.Body)
	var request CreateUserRequest
	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if request.GeneratePassword {
		res, err := password.Generate(14, 3, 2, false, true)
		if err != nil {
			utils.HandleServerError("Failed to generate password", err,w)
			return
		}

		request.Password = res
	}

	api.sc.MongoClient.Database("identity").Collection("users").InsertOne(context.TODO(), models.User{
		Email:    request.Email,
		Password: request.Password,
	})

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "User created")
}

type CreateUserRequest struct {
	Email            string `json:"email"`
	Password         string `json:"password"`
	GeneratePassword bool   `json:"generatePassword"`
}
