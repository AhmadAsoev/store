package services

import (
	"net/http"
	"store/pkg/domain"
)

var Products []domain.Store

func Product(info domain.Store) (response domain.Response) {
	if err := info.Validate(); err != nil {
		return domain.Response{
			Code:    http.StatusBadRequest,
			Error:   err.Error(),
			Message: "",
		}
	}

	Products = append(Products, info)

	return domain.Response{
		Code:    http.StatusOK,
		Error:   "",
		Message: "Success add product",
	}
}
