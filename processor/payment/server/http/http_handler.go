package http

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
	"processor/payment"
	"processor/payment/models"
)

type handler struct {
	service payment.Service
}

//NewHandler instantiate a handler and fix it with the router
func NewHandler(s payment.Service, r *chi.Mux) {
	h := handler{
		service: s,
	}

	h.AssignRoute(r)
}

func (h *handler) ProcessPayment(w http.ResponseWriter, r *http.Request) {
	resp := models.Response{
		Success: false,
		Errors:  models.Error{},
	}

	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Couldn't process request", http.StatusBadRequest)
		return
	}

	var req models.Request

	err = json.Unmarshal(bytes, &req)

	if err != nil {
		http.Error(w, "Couldn't process request", http.StatusBadRequest)
		return
	}

	success := h.service.ProcessPayment(req, &resp.Errors)

	if len(resp.Errors.Validation) > 0 {
		http.Error(w, resp.Errors.Validation[0].Error(), http.StatusBadRequest)
		return
	}

	if len(resp.Errors.Internal) > 0 {
		http.Error(w, resp.Errors.Internal[0].Error(), http.StatusInternalServerError)
		return
	}

	resp.Success = success

	res, err := json.Marshal(resp)

	if err != nil {
		http.Error(w, "error making response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
