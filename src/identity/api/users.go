package api

import (
	"net/http"

	"github.com/aghyad-khlefawi/identity/pkg/users"
	"github.com/aghyad-khlefawi/identity/utils"
	"github.com/gin-gonic/gin"
)

func HandleCreateUser(c *gin.Context, api *ApiContext) {
	var request CreateUserRequest
	err:=c.BindJSON(&request)

	if err != nil {
		utils.HandleBadRequest("Invalid request structure",c)
		return
	}
  
	_,err =users.CreateUser(request.Email,request.Email,request.GeneratePassword)

	if err!=nil{
		if err== users.ErrPasswordNotMeetingRequirement{
			utils.HandleBadRequest(err.Error(),c)
			return
		}
		utils.HandleServerError("Failed to create a user",err,c)
		return
	}
	// Generate password if needed
	
  utils.WriteJsonContent(c, utils.Message{Msg: "User created"},http.StatusOK)
}



type CreateUserRequest struct {
	Email            string `json:"email"`
	Password         string `json:"password"`
	GeneratePassword bool   `json:"generatePassword"`
}
