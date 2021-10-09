package handlers

import (
	"crypto/rand"
	"encoding/json"
	"log"
	"net/http"
	"std/github.com/insta/pkg/mocks"
	"std/github.com/insta/pkg/models"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.Readll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var user models.User
	json.Unmarshal(body, &user)

	user.id = rand.Int(100)
	mocks.Users = append(mocks.Users, user)

	w.WriteHeader() = append(mocks.Users, user)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w.Encode("Created"))
}
