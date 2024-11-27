package utils

import (
	"net/http"
)

func HandleServerError(message string, err error, w http.ResponseWriter) {
	LogError(message, err)
	error:=WriteJsonContent(w, Message {Msg: message},http.StatusInternalServerError)
	LogError("Message Serialization failed",error)
}

func HandleBadRequest(message string, w http.ResponseWriter) {
	WriteJsonContent(w, Message {Msg: message},http.StatusBadRequest)
}

type Message struct{
	Msg string `json:"msg"`
}
