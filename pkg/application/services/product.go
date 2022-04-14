package services

import (
	"net/http"
	"store/pkg/application/repository"
	"store/pkg/domain"
)

func Product(info domain.Store) (response domain.Response) {
	if err := info.Validate(); err != nil {
		return domain.Response{
			Code:    http.StatusBadRequest,
			Error:   err.Error(),
			Message: "",
		}
	}

	if err := repository.AddProduct(info); err != nil {
		return domain.Response{
			Code:  http.StatusInternalServerError,
			Error: "Can't add into db",
		}
	}

	return domain.Response{
		Code:    http.StatusOK,
		Error:   "",
		Message: "Success add product",
	}
}
