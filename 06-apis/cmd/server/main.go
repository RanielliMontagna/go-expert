package main

import (
	"goexpert/apis/configs"
	"goexpert/apis/internal/entity"
	"goexpert/apis/internal/infra/database"
	"goexpert/apis/internal/infra/webserver/handlers"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/products/{id}", productHandler.GetProduct)
	r.Get("/products", productHandler.GetProducts)
	r.Post("/products", productHandler.CreateProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)

	r.Post("/users", userHandler.Create)

	http.ListenAndServe(":8000", r)
}
