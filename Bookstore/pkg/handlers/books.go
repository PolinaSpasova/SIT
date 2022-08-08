package handlers

import (
	"api/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
)

func (h handler) GetAllBooks(ctx echo.Context) error {
	books := new(models.Books)
	if err := h.DB.Find(&books.AllBooks).Error; err != nil {
		return err
	}
	res := new(models.AllBooksAC)
	for i := 0; i < len(books.AllBooks); i++ {
		var book models.BookAC
		book.Id = books.AllBooks[i].Id
		book.Title = books.AllBooks[i].Title
		book.Price = books.AllBooks[i].Price

		if err := h.DB.Find(&book.Author, books.AllBooks[i].AuthorId).Error; err != nil {
			return err
		}
		if err := h.DB.Find(&book.Category, books.AllBooks[i].CategoryId).Error; err != nil {
			return err
		}

		res.Books = append(res.Books, book)
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h handler) AddBook(ctx echo.Context) error {
	b := new(models.Book)
	if err := ctx.Bind(b); err != nil {
		return err
	}

	if err := h.DB.Create(&b).Error; err != nil {
		return err
	}

	resp := new(models.BookAC)
	resp.Id = b.Id
	resp.Title = b.Title
	resp.Price = b.Price
	if err := h.DB.Find(&resp.Author, b.AuthorId).Error; err != nil {
		return err
	}
	if err := h.DB.Find(&resp.Category, b.CategoryId).Error; err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (h handler) GetBookByID(ctx echo.Context) error {
	id := ctx.Param("id")
	book := new(models.Book)

	if err := h.DB.First(book, id).Error; err != nil {
		return err
	}

	res := new(models.BookAC)
	res.Id = book.Id
	res.Title = book.Title
	res.Price = book.Price

	if err := h.DB.Find(&res.Author, book.AuthorId).Error; err != nil {
		return err
	}
	if err := h.DB.Find(&res.Category, book.CategoryId).Error; err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h handler) UpdateBook(ctx echo.Context) error {
	id := ctx.Param("id")
	book := new(models.Book)
	updated := new(models.BookAC)

	if err := ctx.Bind(updated); err != nil {
		return err
	}

	if err := h.DB.First(book, id).Error; err != nil {
		return err
	}
	if err := h.DB.Model(&book).Updates(&models.Book{Price: updated.Price}).Error; err != nil {
		return err
	}

	updated.Id = book.Id
	updated.Title = book.Title
	updated.Price = book.Price

	if err := h.DB.Find(&updated.Author, book.AuthorId).Error; err != nil {
		return err
	}
	if err := h.DB.Find(&updated.Category, book.CategoryId).Error; err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, updated)
}

func (h handler) DeleteBook(ctx echo.Context) error {
	id := ctx.Param("id")
	book := new(models.Book)
	resp := new(models.BookAC)

	if err := h.DB.Clauses(clause.Returning{}).Where("id = ?", id).Delete(book).Error; err != nil {
		return err
	}
	resp.Id = book.Id
	resp.Title = book.Title
	resp.Price = book.Price

	if err := h.DB.Find(&resp.Author, book.AuthorId).Error; err != nil {
		return err
	}
	if err := h.DB.Find(&resp.Category, book.CategoryId).Error; err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, resp)
}
