package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Christheoreo/project-manager/database"
	"github.com/Christheoreo/project-manager/handlers"
	"github.com/Christheoreo/project-manager/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	/**
	 * Set up the database connection
	 * Then defer the closing.
	 **/

	client, err := database.EstablishClient()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")
	defer database.EndClient(client)

	r := mux.NewRouter()

	userCollection := client.Database("projects").Collection("users")
	userModel := models.User{
		Collection: userCollection,
	}

	userHandler := handlers.UserHandler{
		UserModel: userModel,
	}

	authHandler := handlers.AuthHandler{
		UserModel: userModel,
	}

	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/users/register", userHandler.RegisterHandler).Methods("POST")
	r.HandleFunc("/auth/login", authHandler.LoginHandler).Methods("POST")

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	http.ListenAndServe(port, r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"message": "Hello!"}`)
}
