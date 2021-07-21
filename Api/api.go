package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Akhenaten-Sama/go-fintech/helpers"
	"github.com/Akhenaten-Sama/go-fintech/users"
	"github.com/gorilla/mux"
)

type Login struct {
	Username string
	Password string
}

type ErrResponse struct {
	Message string
}

func login (w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)


	var formattedBody Login
	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)

	login := users.Login(formattedBody.Username, formattedBody.Password)
	if login["message"] == "all is fine" {
		resp := login
		json.NewEncoder(w).Encode(resp)
	} else {
		resp:= ErrResponse{Message: "Wrong username or Password"}
		json.NewEncoder(w).Encode(resp)
	}
}

func StartAPi(){
	router := mux.NewRouter()
   router.HandleFunc("/login", login).Methods("POST")
   fmt.Println("App is running on port : 9090")
   log.Fatal(http.ListenAndServe(":9090", router))
}