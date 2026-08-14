package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ducminhgd/intelligent-inventory/internal/application/manufacturer"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type ManufacturerAPI struct {
	getService manufacturer.GetManufacturer
	logger     *zap.Logger
}

func (api *ManufacturerAPI) RegisterRoutes() http.Handler {
	cr := chi.NewRouter()
	cr.Get("/manufacturers/{id}", api.Get)
	return cr
}

func (api *ManufacturerAPI) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	output, err := api.getService.Execute(ctx, manufacturer.GetManufacturerRequest{
		ID: uint32(id),
	})
	if err != nil {
		api.logger.Error("Failed to get manufacturer", zap.Error(err))
		http.Error(w, "Failed to get manufacturer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
