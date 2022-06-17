package main

import (
	"encoding/json"
	"final-project/entity"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
	http.HandleFunc("/users/", UserHandler)

	fmt.Println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		fmt.Println(r.URL.Path[1:])
		paths := strings.Split(r.URL.Path[1:], "/")
		fmt.Printf("%+v + \n", paths)

		if len(paths) == 2 && paths[1] != "" {
			fmt.Println(paths[1])
			if ID, err := strconv.Atoi(paths[1]); err == nil {
				fmt.Println(users[ID])
				jsonData, _ := json.Marshal(users[ID])
				w.Header().Add("Content-Type", "application/json")
				w.Write(jsonData)
			}
		} else {
			jsonData, _ := json.Marshal(&users)
			w.Header().Add("Content-Type", "application/json")
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
		paths := strings.Split(r.URL.Path[1:], "/")
		id, _ := strconv.Atoi(paths[1])
		delete(users, id)
	}
}
