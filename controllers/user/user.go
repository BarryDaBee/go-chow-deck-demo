package usercontroller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BarryDaBee/go-chow-deck/models"
	"github.com/gorilla/mux"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUserById(w http.ResponseWriter, r *http.Request) {

	// body, err := ioutil.ReadAll(r.Body)

	// if err != nil {
	// 	w.WriteHeader(400)
	// 	json.NewEncoder(w).Encode(&models.Response{
	// 		Error: true,
	// 		Msg:   err.Error(),
	// 	})
	// 	return
	// }

	id := mux.Vars(r)["id"]

	filter := bson.M{"_id": id}

	coll := mgm.Coll(&models.User{})

	// if err != nil {
	// 	json.NewEncoder(w).Encode(&models.Response{
	// 		Error: true,
	// 		Msg:   "Invalid Id provided",
	// 	})
	// 	return
	// }

	var users []models.User
	err := coll.SimpleFind(&users, filter)

	if err != nil {
		json.NewEncoder(w).Encode(&models.Response{
			Error: true,
			Msg:   "User not found",
		})
		return
	}

	fmt.Println(err.Error())
	json.NewEncoder(w).Encode(&models.Response{
		Error: false,
		Data:  &users[0],
		Msg:   "success",
	})

}

func CreateUser(user models.User) error {
	ctx := context.TODO()

	err := mgm.Coll(&user).CreateWithCtx(ctx, &user)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	return nil
}
