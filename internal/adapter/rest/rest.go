package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Liveness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func RootRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/liveness", Liveness)
	return r
}
