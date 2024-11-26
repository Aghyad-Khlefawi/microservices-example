package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aghyad-khlefawi/identity/types"
)

func HandleCreateUser(w http.ResponseWriter, r *http.Request,api *ApiContext){
	decoder:= json.NewDecoder(r.Body)
	var request CreateUserRequest
	err:= decoder.Decode(&request)
	if err!= nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
  
  api.sc.MongoClient.Database("identity").Collection("users").InsertOne(context.TODO(),types.User{
		Email: request.Email,
		Password: "",
	})

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w,"User created")
}


type CreateUserRequest struct{
	Email string `json:"email"`
	Password string `json:"password"`
	GeneratePassword bool `json:"generatePassword"`
}
