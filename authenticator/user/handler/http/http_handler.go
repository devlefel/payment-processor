package http

import (
	"authenticator/user"
	"authenticator/user/models"
	"encoding/json"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
)

type handler struct {
	service user.Service
}

//NewHandler instantiate a handler and fix it with the router
func NewHandler(s user.Service, r *chi.Mux) {
	h := handler{
		service: s,
	}

	h.AssignRoute(r)
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	var login models.Login
	jr, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error parsing request", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(jr, &login)

	if err != nil {
		http.Error(w, "Error parsing request", http.StatusBadRequest)
		return
	}

	tokens, errs := h.service.Login(login)

	if errs != nil {
		if len(errs.Validation) > 0 {
			http.Error(w, errs.Validation[0].Error(), http.StatusBadRequest)
			return
		}

		if len(errs.Internal) > 0 {
			http.Error(w, errs.Internal[0].Error(), http.StatusInternalServerError)
			return
		}
	}

	var response = struct {
		Tokens []string
	}{
		Tokens: tokens,
	}

	js, err := json.Marshal(response)

	if err != nil {
		http.Error(w, "error mounting response", http.StatusInternalServerError)
		return
	}

	w.Write(js)
	return
}
