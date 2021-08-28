package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Christheoreo/project-manager/database"
	"github.com/Christheoreo/project-manager/handlers"
	"github.com/Christheoreo/project-manager/repositories"
	"github.com/Christheoreo/project-manager/services"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set up the database connection
	// Then defer the closing.

	pool, err := database.EstablishConnectionPool()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Connected!")
	defer database.CloseConnectionPool(pool)

	r := mux.NewRouter()

	// Set up Repositories
	usersRepository := repositories.UsersRepositoryPostgres{
		Pool: pool,
	}
	prioritiesRepository := repositories.PrioritiesRepositoryPostgres{
		Pool: pool,
	}
	tagsRepository := repositories.TagsRepositoryPostgres{
		Pool: pool,
	}
	projectsRepository := repositories.ProjectsRepositoryPostgres{
		Pool: pool,
	}

	// Set up the services
	usersService := services.UsersService{
		UsersRepository: &usersRepository,
	}
	prioritiesService := services.PrioritiesService{
		PrioritiesRepository: &prioritiesRepository,
	}
	tagsService := services.TagsService{
		TagsRepository: &tagsRepository,
	}
	projectsService := services.ProjectsService{
		ProjectsRepository: &projectsRepository,
	}

	// Set up handlers
	usersHandler := handlers.UserHandler{
		UsersService: &usersService,
	}

	authHandler := handlers.AuthHandler{
		UsersService: &usersService,
	}

	tagsHandler := handlers.TagsHandler{
		TagsService: &tagsService,
	}
	projectsHandler := handlers.ProjectsHandler{
		ProjectsService:   &projectsService,
		TagsService:       &tagsService,
		PrioritiesService: &prioritiesService,
	}

	// Set up middleware
	jwtMiddleware := handlers.JWTMiddleware{
		UsersService: &usersService,
	}

	// Applying top level middleware
	r.Use(handlers.HeadersMiddleware)
	r.Use(handlers.CorsMiddleware)

	// Unprotected routes
	r.HandleFunc("/users/register", usersHandler.RegisterHandler).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/auth/login", authHandler.LoginHandler).Methods(http.MethodPost, http.MethodOptions)

	//pR = protectedRoutes
	pR := r.PathPrefix("/").Subrouter()

	// Apply middleware
	pR.Use(jwtMiddleware.Middleware)

	// Define routes
	pR.HandleFunc("/users/me", usersHandler.GetRequester).Methods(http.MethodGet, http.MethodOptions)
	pR.HandleFunc("/tags", tagsHandler.Create).Methods(http.MethodPost, http.MethodOptions)
	pR.HandleFunc("/tags", tagsHandler.GetAllForRequester).Methods(http.MethodGet, http.MethodOptions)
	pR.HandleFunc("/projects", projectsHandler.CreateProjectHandler).Methods(http.MethodPost, http.MethodOptions)
	pR.HandleFunc("/projects", projectsHandler.GetMyProjects).Methods(http.MethodGet, http.MethodOptions)
	pR.HandleFunc("/projects/{projectID}", projectsHandler.GetMyProject).Methods(http.MethodGet, http.MethodOptions)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	http.ListenAndServe(port, r)
}
