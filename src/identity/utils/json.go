package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJsonContent(w http.ResponseWriter, content any, status int) error {
	bytes, err := json.Marshal(content)
	if err!=nil{
		LogError("Serialization error",err)
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(bytes)
	return nil
}

func DeserializeJsonRequest[T any](r *http.Request) (*T,error){
	decoder := json.NewDecoder(r.Body)
	var request T
	err := decoder.Decode(&request)
	return &request,err
}
