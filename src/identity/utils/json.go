package util

import (
	"encoding/json"
	"net/http"
)

func WriteJsonContent(w http.ResponseWriter, content any) {
	bytes, _ := json.Marshal(content)
	w.Header().Add("Content-Type", "application/json")
	w.Write(bytes)
}
