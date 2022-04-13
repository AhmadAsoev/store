package handleFunc

import (
	"net/http"
	"store/pkg/application/services"
)

func AllProducts(w http.ResponseWriter, r *http.Request) {
	response := services.AllProducts()
	response.Send(w, "AllProducts")
}
