package usercontroller

import (
	"bytes"
	"encoding/json"
	"os"
	"time"

	// "fmt"
	"io"
	"net/http"

	"github.com/BarryDaBee/go-chow-deck/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	// "github.com/gorilla/mux"
	// "github.com/kamva/mgm/v3"
	// "go.mongodb.org/mongo-driver/bson"
)

func SignIn(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	err = json.NewDecoder(bytes.NewReader(body)).Decode(&user)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(&models.Response{
			Error: true,
			Msg:   err.Error(),
		})
		return
	}

	token, err := GenerateToken(user.ID, 24) // Assuming user.ID is unique identifier for the user.
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(&models.Response{
			Error: true,
			Msg:   "Failed to generate token",
		})
		return
	}

	// Respond with success and include the generated token.
	err = json.NewEncoder(w).Encode(&models.Response{
		Error: false,
		Data: map[string]interface{}{
			"token": token,
			"user":  user,
		},
		Msg: "Sign in successful",
	})
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(&models.Response{
			Error: true,
			Msg:   "Failed to encode response",
		})
		return
	}

}

func SignUp(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	w.Header().Set("Content-Type", "application/json")

	var newUser models.User
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&newUser)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(&models.Response{
			Error: true,
			Msg:   err.Error(),
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)

	newUser.Password = hashedPassword

	CreateUser(newUser)

	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(&models.Response{
			Error: true,
			Msg:   "Failed to hash password",
		})
		return
	}

	// Save the user to the database, including the hashed password.
	// Replace this with your actual database saving logic.
	// For demonstration purposes, assume the user is saved successfully.

	token, err := GenerateToken(newUser.ID, 24)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(&models.Response{
			Error: true,
			Msg:   "Failed to generate token",
		})
		return
	}

	err = json.NewEncoder(w).Encode(&models.Response{
		Error: false,
		Data: map[string]interface{}{
			"token": token,
			"user":  newUser,
		},
		Msg: "Sign up successful",
	})
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(&models.Response{
			Error: true,
			Msg:   "Failed to encode response",
		})
		return
	}
}

func GenerateToken(data interface{}, args ...int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	hour := time.Hour * 24

	if len(args) > 0 {
		hour = time.Hour * time.Duration(args[0])
	}

	claims := token.Claims.(jwt.MapClaims)
	claims["data"] = data
	claims["exp"] = time.Now().Add(hour).Unix()

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {

		return "", err
	}

	return t, nil
}
