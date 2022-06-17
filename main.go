package main

import (
	"encoding/json"
	"final-project/entity"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// func main() {
// 	userSvc := service.NewUserSvc()
// 	if user, err := userSvc.Register(&entity.User{
// 		Username: "usamah07",
// 		Email:    "usamah@gmail.com",
// 		Password: "usamah",
// 		Age:      10,
// 	}); err != nil {
// 		fmt.Printf("error when register %+v", err)
// 	} else {
// 		fmt.Printf("succes when register %+v", user)
// 	}
// }

// var users = []entity.User{
// 	{ID: 1, Username: "usamah", Email: "usamah@gmail.com", Password: "usamah"},
// 	{ID: 2, Username: "abdurrahman", Email: "abdurrahman@gmail.com", Password: "abdurrahman"},
// 	{ID: 3, Username: "eko", Email: "eko@gmail.com", Password: "eko"},
// }

var users = map[int]entity.User{
	1: {
		ID:       1,
		Username: "usamah",
		Email:    "usamah@gmail.com",
		Password: "usamah",
	},
	2: {
		ID:       2,
		Username: "abdurrahman",
		Email:    "abdurrahman@gmail.com",
		Password: "abdurrahman",
	},
}

var PORT = ":8080"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users/", UserHandler)
	r.HandleFunc("/users/{id}", UserHandler)

	fmt.Println("Application is listening on port", PORT)

	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]
	fmt.Println(id)

	if r.Method == "GET" {
		if id != "" {
			if ID, err := strconv.Atoi(id); err == nil {
				jsonData, _ := json.Marshal(users[ID])
				w.Write(jsonData)
			}
		} else {
			jsonData, _ := json.Marshal(&users)
			w.Write(jsonData)
		}
	}

	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var user entity.User
		err := decoder.Decode(&user)
		if err != nil {
			panic(err)
		}

		//menambahkan data user baru ke data
		users[user.ID] = user

		json.NewEncoder(w).Encode(user)
		return
	}

	if r.Method == "PUT" {
		decoder := json.NewDecoder(r.Body)
		var user entity.User
		err := decoder.Decode(&user)
		if err != nil {
			panic(err)
		}

		users[user.ID] = user

		json.NewEncoder(w).Encode(user)
		return
	}

	if r.Method == "DELETE" {
		idInt, _ := strconv.Atoi(id)
		delete(users, idInt)
	}
}
