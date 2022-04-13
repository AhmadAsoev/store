package handleFunc

import (
	"encoding/json"
	"log"
	"net/http"
	"store/pkg/application/services"
	"store/pkg/domain"
)

func Product(w http.ResponseWriter, r *http.Request) {
	var request domain.Store

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte("Can not Decode to json")); err != nil {
			log.Println("HandleFunc/Product/Decode/Write")
			return
		}
		log.Println("HandleFunc/Product/Decode")
	}

	response := services.Product(request)
	response.Send(w, "Product")
}
