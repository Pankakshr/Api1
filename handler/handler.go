package handler

import (
	"Api1/modle"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var allUser = []modle.User{
	{Name: "vikash", Id: 12, Add: "kheri"},
	{Name: "pawan", Id: 11, Add: "khweri1"},
	{Name: "dav", Id: 31, Add: "kheri2"},
	{Name: "ram", Id: 24, Add: "kheri"},
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	//var newUser modle.User
	if r.Method == http.MethodDelete {
		//json.NewDecoder(r.Body).Decode(newUser)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		x := r.PathValue("id")
		id, err := strconv.Atoi(x)
		if err != nil {
			log.Fatalf("err is %v", err)
		}
		for i, v := range allUser {
			if v.Id == id {
				allUser = append(allUser[:i], allUser[i+1:]...)
			}

		}
		json.NewEncoder(w).Encode("given id is deleted")
		fmt.Println(allUser)

	}

}
