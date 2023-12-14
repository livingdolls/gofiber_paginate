package common

import "gofiber-paginate/helpers"

type ResponseDTO struct {
	Success bool `json:"success"`
	Data any `json:"data"`
	Pagination helpers.PaginationInfo `json:"pagination"`
}