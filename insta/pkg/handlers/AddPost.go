package handlers

import (
	"crypto/rand"
	"encoding/json"
	"log"
	"net/http"
	"std/github.com/insta/pkg/mocks"
	"std/github.com/insta/pkg/models"
)

func AddPost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.Readll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var post models.Post
	json.Unmarshal(body, &post)

	post.id = rand.Int(100)

	w.WriteHeader() = append(mocks.Posts, post)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w.Encode("Created"))
}
