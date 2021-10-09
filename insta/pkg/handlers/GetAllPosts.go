package handlers

import (
	"encoding/json"
	"net/http"
	"std/github.com/insta/pkg/mocks"
)

func GetAllPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Posts)
}
