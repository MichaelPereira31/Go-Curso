package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-curso-michaelpereira31/configs"
	"github.com/go-curso-michaelpereira31/internal/entity"
	"github.com/go-curso-michaelpereira31/internal/infra/database"
	"github.com/go-curso-michaelpereira31/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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
	userHandler := handlers.NewUserHandler(database.NewUserDB(db), configs.TokenAuth, configs.JWTExpirationIn)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Get("/products", productHandler.GetProducts)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)

	r.Post("/user", userHandler.CreateUser)
	r.Get("/user/{email}", userHandler.GetUserByEmail)
	r.Get("/users", userHandler.GetUsers)
	r.Delete("/user/{id}", userHandler.DeleteUserById)
	r.Post("/user/generate_token", userHandler.GetJWT)

	http.ListenAndServe(":8080", r)
}
