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
	tagCollection := client.Database("projects").Collection("tags")

	// Set up models
	userModel := models.User{
		Collection: userCollection,
	}
	tagModel := models.Tag{
		Collection: tagCollection,
	}

	// Set up handlers
	userHandler := handlers.UserHandler{
		UserModel: userModel,
	}

	authHandler := handlers.AuthHandler{
		UserModel: userModel,
	}
	tagHandler := handlers.TagsHandler{
		TagModel: tagModel,
	}

	// Set up middleware
	jwtMiddleware := middleware.JWTMiddleware{
		UserModel: userModel,
	}

	// Applying top level middleware
	r.Use(middleware.HeadersMiddleware)
	r.Use(middleware.CorsMiddleware)

	// Unprotected routes
	r.HandleFunc("/users/register", userHandler.RegisterHandler).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/auth/login", authHandler.LoginHandler).Methods(http.MethodPost, http.MethodOptions)

	//pR = protectedRoutes
	pR := r.PathPrefix("/").Subrouter()

	// Apply middleware
	pR.Use(jwtMiddleware.Middleware)

	// Define routes
	pR.HandleFunc("/users/me", userHandler.GetRequester).Methods(http.MethodGet, http.MethodOptions)
	/**
		@todo this needs testing
		carry on here in postman!
	**/
	pR.HandleFunc("/tags", tagHandler.RegisterHandler).Methods(http.MethodPost, http.MethodOptions)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	http.ListenAndServe(port, r)
}
