package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Christheoreo/project-manager/database"
	"github.com/Christheoreo/project-manager/handlers"
	"github.com/Christheoreo/project-manager/middleware"
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

	// Set up collections
	userCollection := client.Database("projects").Collection("users")

	// Set up models
	userModel := models.User{
		Collection: userCollection,
	}

	// Set up handlers
	userHandler := handlers.UserHandler{
		UserModel: userModel,
	}

	authHandler := handlers.AuthHandler{
		UserModel: userModel,
	}

	// Set up middleware
	jwtMiddleware := middleware.JWTMiddleware{
		UserModel: userModel,
	}

	// Applying top level middleware
	r.Use(middleware.HeadersMiddleware)

	// Unprotected routes
	r.HandleFunc("/users/register", userHandler.RegisterHandler).Methods("POST")
	r.HandleFunc("/auth/login", authHandler.LoginHandler).Methods("POST")

	//pR = protectedRoutes
	pR := r.PathPrefix("/").Subrouter()

	// Apply middleware
	pR.Use(jwtMiddleware.Middleware)

	// Define routes
	pR.HandleFunc("/users/me", userHandler.GetRequester).Methods("GET")

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	http.ListenAndServe(port, r)
}
