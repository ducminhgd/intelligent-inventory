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
	svc    *manufacturer.ManufacturerService
	logger *zap.Logger
}

func NewManufacturerAPI(svc *manufacturer.ManufacturerService, logger *zap.Logger) *ManufacturerAPI {
	return &ManufacturerAPI{svc: svc, logger: logger}
}

func (api *ManufacturerAPI) Router() http.Handler {
	cr := chi.NewRouter()
	cr.Get("/manufacturers/{id}", api.Get)
	cr.Get("/manufacturers", api.List)
	cr.Post("/manufacturers", api.Create)
	cr.Put("/manufacturers", api.Update)
	cr.Delete("/manufacturers/{id}", api.Delete)
	return cr
}

func (api *ManufacturerAPI) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}
	req := manufacturer.GetManufacturerRequest{ID: uint32(id)}

	output, err := api.svc.GetByID(ctx, req)
	if err != nil {
		api.logger.Error("Failed to get manufacturer", zap.Error(err))
		http.Error(w, "Failed to get manufacturer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(output); err != nil {
		api.logger.Error("Failed to encode response", zap.Error(err))
	}
}

func (api *ManufacturerAPI) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req manufacturer.ListManufacturerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	output, err := api.svc.List(ctx, req)
	if err != nil {
		api.logger.Error("Failed to list manufacturers", zap.Error(err))
		http.Error(w, "Failed to list manufacturers", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(output); err != nil {
		api.logger.Error("Failed to encode response", zap.Error(err))
	}
}

func (api *ManufacturerAPI) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req manufacturer.CreateManufacturerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	output, err := api.svc.Create(ctx, req)
	if err != nil {
		api.logger.Error("Failed to create manufacturer", zap.Error(err))
		http.Error(w, "Failed to create manufacturer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(output); err != nil {
		api.logger.Error("Failed to encode response", zap.Error(err))
	}
}

func (api *ManufacturerAPI) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req manufacturer.UpdateManufacturerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	output, err := api.svc.Update(ctx, req)
	if err != nil {
		api.logger.Error("Failed to update manufacturer", zap.Error(err))
		http.Error(w, "Failed to update manufacturer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(output); err != nil {
		api.logger.Error("Failed to encode response", zap.Error(err))
	}
}

func (api *ManufacturerAPI) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}
	req := manufacturer.DeleteManufacturerRequest{ID: uint32(id)}

	output, err := api.svc.Delete(ctx, req)
	if err != nil {
		api.logger.Error("Failed to delete manufacturer", zap.Error(err))
		http.Error(w, "Failed to delete manufacturer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(output); err != nil {
		api.logger.Error("Failed to encode response", zap.Error(err))
	}
}
