package handlers

import (
	"fmt"
	"net/http"
)

func GetUserId(w http.ResponseWriter, r *http.Request) {
	 vars := mux.Vars(r)
	 id, _ =strconv.Atop(vars["ID"])

	 for user.ID =id{
    if user.ID =id{
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-type","application/json")
		json.NewEncoder(w).Encode(user)
		break
	}
	 }

}
