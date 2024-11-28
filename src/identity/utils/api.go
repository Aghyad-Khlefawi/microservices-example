package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleServerError(message string, err error, c *gin.Context) {
	LogError(message, err)
	WriteJsonContent(c, Message {Msg: message},http.StatusInternalServerError)
}

func HandleBadRequest(message string, c *gin.Context) {
	WriteJsonContent(c, Message {Msg: message},http.StatusBadRequest)
}

type Message struct{
	Msg string `json:"msg"`
}
