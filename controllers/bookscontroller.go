package controllers

import (
	"gofiber-paginate/common"
	"gofiber-paginate/database"
	"gofiber-paginate/helpers"
	"gofiber-paginate/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetPaginatedBooks(c *fiber.Ctx) error {
	books := []models.Book{}

	// per_page (size), sort_order
	perPage := c.Query("per_page", "10")
	sortOrder := c.Query("sort_order", "desc")
	cursor := c.Query("cursor", "")

	limit, err := strconv.ParseInt(perPage, 10, 64)

	if limit < 1 || limit > 100 {
		limit = 10
	}

	if err != nil {
		return c.Status(500).JSON("Invalid per_page option")
	}

	// query
	isFirstPage := cursor == ""
	pointsNext := false

	
	query := database.DB
	query, pointsNext, err = database.GetPaginatedQuery(query, pointsNext, cursor, sortOrder)

	if err != nil {
		return c.Status(500).JSON("unable to get the paginated query")
	}

	err = query.Limit(int(limit) +1).Find(&books).Error
	if err != nil {
		return c.Status(500).JSON("Unable to get the books data")
	}

	hasPagination := len(books) > int(limit)

	if hasPagination {
		books = books[:limit]
	}

	if !isFirstPage && !pointsNext {
		books = helpers.Reverse(books)
	}

	pageInfo := database.CalculatePagination(isFirstPage, hasPagination, int(limit), books, pointsNext)

	response := common.ResponseDTO {
		Success: true,
		Data : books,
		Pagination: pageInfo,
	}

	return c.Status(fiber.StatusOK).JSON(response)
	// get the data

	// manipulate data, calculate
}