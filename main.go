package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/BarryDaBee/go-chow-deck/models"
	"github.com/BarryDaBee/go-chow-deck/routers"
	"github.com/gorilla/mux"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	const connectionString = "mongodb+srv://potato_man:secret123@cluster0.c7x7l.mongodb.net/?retryWrites=true&w=majority"
	err := mgm.SetDefaultConfig(nil, "go_chow_deck", options.Client().ApplyURI(connectionString))

	checkNilError(err)

	var user models.User = models.User{}
	fmt.Println(user.ID.Hex())

	err = mgm.Coll(&models.User{}).Create(&user)
	checkNilError(err)

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(&models.Response{
			Error: false,
			Data:  "Server is live",
			Msg:   "Server Live",
		})
	})
	userRouter := r.PathPrefix("/user").Subrouter()
	authRouter := r.PathPrefix("/auth").Subrouter()

	PORT := 4000
	routers.UserRoutes(userRouter)
	routers.AuthRoutes(authRouter)
	fmt.Printf("Server is running on Port %v", PORT)
	log.Fatal(http.ListenAndServe(":4000", r))

}

func checkNilError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
