package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WriteJsonContent(c *gin.Context, content any, status int) error {
	c.JSON(status, content)
	return nil
}

func DeserializeJsonRequest[T any](r *http.Request) (*T, error) {
	decoder := json.NewDecoder(r.Body)
	var request T
	err := decoder.Decode(&request)
	return &request, err
}
