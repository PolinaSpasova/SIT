package handlers

import (
	"api/pkg/models"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (h handler) GetAllCategories(ctx echo.Context) error {
	categories := new(models.Categories)
	if err := h.DB.Find(&categories.AllCat).Error; err != nil {
		return err
	}
	for i := 0; i < len(categories.AllCat); i++ {
		result := h.DB.First(&categories.AllCat[i].Books, "category_id", categories.AllCat[i].Id)
		if result.Error != nil {
			if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return result.Error
			}
		}
	}

	return ctx.JSON(http.StatusOK, categories)
}

func (h handler) AddCategory(ctx echo.Context) error {
	c := new(models.Category)
	if err := ctx.Bind(c); err != nil {
		return err
	}

	if err := h.DB.Create(&c).Error; err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, c)
}

func (h handler) GetCategoryByID(ctx echo.Context) error {
	id := ctx.Param("id")
	res := new(models.Category)

	if err := h.DB.First(res, id).Error; err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h handler) UpdateCategory(ctx echo.Context) error {
	id := ctx.Param("id")
	c := new(models.Category)
	updated := new(models.Category)
	if err := ctx.Bind(c); err != nil {
		return err
	}

	if err := h.DB.First(updated, id).Error; err != nil {
		return err
	}

	updated.Name = c.Name

	h.DB.Save(&updated)

	return ctx.JSON(http.StatusOK, updated)
}

func (h handler) DeleteCategory(ctx echo.Context) error {
	id := ctx.Param("id")
	category := new(models.Category)
	books := []models.Book{}

	if err := h.DB.Clauses(clause.Returning{}).Where("id = ?", id).Delete(category).Error; err != nil {
		return err
	}

	h.DB.Where("category_id = ?", category.Id).Delete(&books)

	return ctx.JSON(http.StatusOK, category)
}
