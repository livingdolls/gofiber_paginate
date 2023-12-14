package database

import (
	"fmt"
	"gofiber-paginate/helpers"
	"time"

	"gorm.io/gorm"
)

type PaginatedItem interface {
	GetId() string
	GetCreatedAt() time.Time
}

func GetPaginatedQuery(query *gorm.DB, pointsNext bool, cursor string, sortOrder string) (*gorm.DB, bool, error) {
	if cursor != "" {
		decodedCursor, err := helpers.DecodeCursor(cursor)

		if err != nil {
			return nil, pointsNext, err
		}

		pointsNext = decodedCursor["points_next"] == true
		operator, order := getPaginatedOperator(pointsNext, sortOrder)
		whereStr := fmt.Sprintf("(created_at %s OR (created_at = ? AND id %s ?))", operator, operator)
		query = query.Where(whereStr, decodedCursor["created_at"], decodedCursor["created_at"], decodedCursor["id"])

		if order != "" {
			sortOrder = order
		}
	}

	query = query.Order("created_at " + sortOrder)

	return query, pointsNext, nil
}

func getPaginatedOperator(pointsNext bool, sortOrder string) (string, string) {
	if pointsNext && sortOrder == "asc" {
		return ">", ""
	}

	if pointsNext && sortOrder == "desc" {
		return "<", ""
	}

	if !pointsNext && sortOrder == "asc" {
		return "<", "desc"
	}

	if pointsNext && sortOrder == "desc" {
		return "<", "asc"
	}

	return "", "";
}

func CalculatePagination[T PaginatedItem](isFirstPage bool, hasPagination bool, limit int, items []T, pointsNext bool) helpers.PaginationInfo {
	pagination := helpers.PaginationInfo{}
	nextCur := helpers.Cursor{}
	prevCur := helpers.Cursor{}

	if isFirstPage {
		if hasPagination {
			nextCur = helpers.CreateCursor(items[limit-1].GetId(), items[limit-1].GetCreatedAt(), true)
			pagination = helpers.GeneratePager(nextCur, prevCur)
		}
	} else {
		if pointsNext {
			if hasPagination {
				nextCur = helpers.CreateCursor(items[limit-1].GetId(), items[limit-1].GetCreatedAt(), true)
			}

			prevCur = helpers.CreateCursor(items[0].GetId(), items[0].GetCreatedAt(), false)
			pagination = helpers.GeneratePager(nextCur, prevCur)
		} else {
			nextCur = helpers.CreateCursor(items[limit-1].GetId(), items[limit-1].GetCreatedAt(), true)
			if hasPagination {
				prevCur = helpers.CreateCursor(items[0].GetId(), items[0].GetCreatedAt(), false)
			}

			pagination = helpers.GeneratePager(nextCur, prevCur)
		}
	}

	return pagination
}