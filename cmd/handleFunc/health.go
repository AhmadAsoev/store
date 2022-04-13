package handleFunc

import (
	"net/http"
	"store/pkg/domain"
)

func Health(w http.ResponseWriter, r *http.Request) {
	response := domain.Response{
		Code:    http.StatusOK,
		Error:   "",
		Message: "Server is ready for work",
	}

	response.Send(w, "Health")
}
