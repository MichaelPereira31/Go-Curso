package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/go-curso-michaelpereira31/configs"
	"github.com/go-curso-michaelpereira31/internal/entity"
	"github.com/go-curso-michaelpereira31/internal/infra/database"
	"github.com/go-curso-michaelpereira31/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "github.com/go-curso-michaelpereira31/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Go Expert API
// @version 1.0
// @description Product API with authentication

// @contact.name Michael
// @contact.url https://www.linkedin.com/in/michaelpereira31/

// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productHandler := handlers.NewProductHandler(database.NewProductDB(db))
	userHandler := handlers.NewUserHandler(database.NewUserDB(db))
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("JwtExpirationIn", configs.JWTExpirationIn))
	r.Use(LogRequest)
	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/{id}", productHandler.GetProduct)
		r.Get("/", productHandler.GetProducts)
		r.Post("/", productHandler.CreateProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/user", userHandler.CreateUser)
	r.Get("/user/{email}", userHandler.GetUserByEmail)
	r.Get("/users", userHandler.GetUsers)
	r.Delete("/user/{id}", userHandler.DeleteUserById)
	r.Post("/user/generate_token", userHandler.GetJWT)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/doc.json")))

	http.ListenAndServe(":8080", r)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
