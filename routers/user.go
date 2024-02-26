package routers

import (
	usercontroller "github.com/BarryDaBee/go-chow-deck/controllers/user"
	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	r.HandleFunc("/{id}", usercontroller.GetUserById).Methods("GET")
}
