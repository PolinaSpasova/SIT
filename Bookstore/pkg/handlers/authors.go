package handlers

import (
	"api/pkg/models"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (h handler) GetAllAuthors(ctx echo.Context) error {
	authors := new(models.Authors)
	if err := h.DB.Find(&authors.AllAuthors).Error; err != nil {
		return err
	}

	for i := 0; i < len(authors.AllAuthors); i++ {
		result := h.DB.First(&authors.AllAuthors[i].Books, "author_id", authors.AllAuthors[i].Id)
		if result.Error != nil {
			if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return result.Error
			}
		}
	}

	return ctx.JSON(http.StatusOK, authors)
}

func (h handler) AddAuthor(ctx echo.Context) error {
	a := new(models.Author)
	if err := ctx.Bind(a); err != nil {
		return err
	}

	result := h.DB.Create(&a)
	if result.Error != nil {
		return result.Error
	}

	return ctx.JSON(http.StatusOK, a)
}

func (h handler) GetAuthorByID(ctx echo.Context) error {
	id := ctx.Param("id")
	res := new(models.Author)

	result := h.DB.First(res, id)
	if result.Error != nil {
		return result.Error
	}

	if err := h.DB.Find(&res.Books, res.Id).Error; err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h handler) UpdateAuthor(ctx echo.Context) error {
	id := ctx.Param("id")
	a := new(models.Author)
	updated := new(models.Author)
	if err := ctx.Bind(a); err != nil {
		return err
	}

	if err := h.DB.First(updated, id).Error; err != nil {
		return err
	}

	updated.Name = a.Name
	updated.Biography = a.Biography

	h.DB.Save(&updated)

	return ctx.JSON(http.StatusOK, updated)
}

func (h handler) DeleteAuthor(ctx echo.Context) error {
	id := ctx.Param("id")
	author := new(models.Author)
	books := []models.Book{}

	if err := h.DB.Clauses(clause.Returning{}).Where("id = ?", id).Delete(author).Error; err != nil {
		return err
	}

	h.DB.Where("author_id = ?", author.Id).Delete(&books)

	return ctx.JSON(http.StatusOK, author)
}
