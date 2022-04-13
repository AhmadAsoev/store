package services

import (
	"net/http"
	"store/pkg/domain"
)

func AllProducts() domain.Response {
	return domain.Response{
		Code:    http.StatusOK,
		Error:   "",
		Message: Products,
	}
}
