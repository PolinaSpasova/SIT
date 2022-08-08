package main

import (
	"api/pkg/db"
	"api/pkg/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)

	router := echo.New()

	router.GET("/categories", h.GetAllCategories)
	router.POST("/categories", h.AddCategory)
	router.GET("/categories/:id", h.GetCategoryByID)
	router.PUT("/categories/:id", h.UpdateCategory)
	router.DELETE("/categories/:id", h.DeleteCategory)

	router.GET("/authors", h.GetAllAuthors)
	router.POST("/authors", h.AddAuthor)
	router.GET("/authors/:id", h.GetAuthorByID)
	router.PUT("/authors/:id", h.UpdateAuthor)
	router.DELETE("/authors/:id", h.DeleteAuthor)

	router.GET("/books", h.GetAllBooks)
	router.POST("/books", h.AddBook)
	router.GET("/books/:id", h.GetBookByID)
	router.PUT("/books/:id", h.UpdateBook)
	router.DELETE("/books/:id", h.DeleteBook)

	router.Logger.Fatal(router.Start(":3000"))

}
