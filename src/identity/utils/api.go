package utils

import (
	"net/http"
)


func HandleServerError(message string,err error, w http.ResponseWriter){
	LogError(message, err)
	w.WriteHeader(http.StatusInternalServerError)
	WriteJsonContent(w,struct{msg string}{msg:message})
}
