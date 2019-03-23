package utils

import (
	"encoding/json"
	"go-books-list/models"
	"fmt"
	"net/http"
)

func SendError(w http.ResponseWriter, status int, error models.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

func SendSuccess(w http.ResponseWriter, data interface{}) {
	fmt.Println(data)
	json.NewEncoder(w).Encode(data)
}
