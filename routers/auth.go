package routers

import (
	usercontroller "github.com/BarryDaBee/go-chow-deck/controllers/user"
	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	r.HandleFunc("/sign-in", usercontroller.SignIn).Methods("POST")
}
