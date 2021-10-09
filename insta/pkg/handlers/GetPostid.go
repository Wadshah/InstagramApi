package handlers

import (
	"fmt"
	"net/http"
)

func GetPost(w http.ResponseWriter, r *http.Request) {
   
	 vars := mux.Vars(r)
	 id, _ =strconv.Atop(vars["ID"])

	 for post.ID =id{
    if post.ID =id{
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-type","application/json")
		json.NewEncoder(w).Encode(post)
		break
	}
	 }

}
