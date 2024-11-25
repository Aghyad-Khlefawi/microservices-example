package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleCreateUser(w http.ResponseWriter, r *http.Request){
	decoder:= json.NewDecoder(r.Body)
	var request CreateUserRequest
	err:= decoder.Decode(&request)
	if err!= nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
  
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w,"User created")
}


type CreateUserRequest struct{
	Email string `json:"email"`
	Password string `json:"password"`
	GeneratePassword bool `json:"generatePassword"`
}
