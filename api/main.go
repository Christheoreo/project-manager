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

	pool, err := database.EstablishConnectionPool()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Connected!")
	defer database.CloseConnectionPool(pool)

	r := mux.NewRouter()

	// Set up collections

	// Set up models
	userModel := models.User{
		Pool: pool,
	}

	tagsModel := models.Tag{
		Pool: pool,
	}
	projectModel := models.Project{
		Pool: pool,
	}

	// Set up handlers
	userHandler := handlers.UserHandler{
		UserModel: userModel,
	}

	authHandler := handlers.AuthHandler{
		UserModel: userModel,
	}

	tagsHandler := handlers.TagsHandler{
		TagModel: tagsModel,
	}
	projectsHandler := handlers.ProjectsHandler{
		ProjectModel: projectModel,
		TagModel:     tagsModel,
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
	pR.HandleFunc("/tags", tagsHandler.RegisterHandler).Methods(http.MethodPost, http.MethodOptions)
	pR.HandleFunc("/tags", tagsHandler.GetAllForRequester).Methods(http.MethodGet, http.MethodOptions)
	pR.HandleFunc("/projects", projectsHandler.CreateProjectHandler).Methods(http.MethodPost, http.MethodOptions)
	pR.HandleFunc("/projects", projectsHandler.GetMyProjects).Methods(http.MethodGet, http.MethodOptions)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	http.ListenAndServe(port, r)
}
