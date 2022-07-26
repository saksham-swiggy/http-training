package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	fmt.Println("server started")
	http.HandleFunc("/", UserInfo)
	http.ListenAndServe("localhost:8080", nil)

}

func UserInfo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "application/json")
	getuser := User{
		Name: "Saksham Kumar",
	}
	switch req.Method {
	case "GET":
		err := json.NewEncoder(w).Encode(getuser)
		if err != nil {
			fmt.Println("error in encoding", err)
		}
	case "POST":
		postuser := User{}
		err := json.NewDecoder(req.Body).Decode(&postuser)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Fprint(w, postuser)
	default:
		fmt.Fprint(w, "only use GET and POST methods!!")
	}
}
