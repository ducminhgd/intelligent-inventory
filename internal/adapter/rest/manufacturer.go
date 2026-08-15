package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	resthttp "github.com/ducminhgd/intelligent-inventory/internal/application/http"
	"github.com/ducminhgd/intelligent-inventory/internal/application/manufacturer"
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
	cr.Put("/manufacturers/{id}", api.Update)
	cr.Delete("/manufacturers/{id}", api.Delete)
	return cr
}

func (api *ManufacturerAPI) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		json.NewEncoder(w).Encode(resthttp.HttpResponse{Error: "invalid user ID"})
		return
	}
	req := manufacturer.GetManufacturerRequest{ID: uint32(id)}

	output, err := api.svc.GetByID(ctx, req)
	if err != nil {
		api.logger.Error("Failed to get manufacturer", zap.Error(err))
		json.NewEncoder(w).Encode(resthttp.HttpResponse{Error: "Failed to get manufacturer"})
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
	req.FromQueryParams(r.URL.Query())

	output, err := api.svc.List(ctx, req)
	if err != nil {
		api.logger.Error("Failed to list manufacturers", zap.Error(err))
		json.NewEncoder(w).Encode(resthttp.HttpResponse{Error: "Failed to list manufacturers"})
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
		json.NewEncoder(w).Encode(resthttp.HttpResponse{Error: "invalid request body"})
		return
	}

	output, err := api.svc.Create(ctx, req)
	if err != nil {
		api.logger.Error("Failed to create manufacturer", zap.Error(err))
		json.NewEncoder(w).Encode(resthttp.HttpResponse{Error: "Failed to create manufacturer"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(output); err != nil {
		api.logger.Error("Failed to encode response", zap.Error(err))
	}
}

func (api *ManufacturerAPI) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		json.NewEncoder(w).Encode(resthttp.HttpResponse{Error: "invalid user ID"})
		return
	}
	var req manufacturer.UpdateManufacturerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		json.NewEncoder(w).Encode(resthttp.HttpResponse{Error: "invalid request body"})
		return
	}
	req.ID = uint32(id)

	output, err := api.svc.Update(ctx, req)
	if err != nil {
		api.logger.Error("Failed to update manufacturer", zap.Error(err))
		json.NewEncoder(w).Encode(resthttp.HttpResponse{Error: "Failed to update manufacturer"})
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
		json.NewEncoder(w).Encode(resthttp.HttpResponse{Error: "invalid user ID"})
		return
	}
	var req manufacturer.DeleteManufacturerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		json.NewEncoder(w).Encode(resthttp.HttpResponse{Error: "invalid request body"})
		return
	}
	req.ID = uint32(id)

	output, err := api.svc.Delete(ctx, req)
	if err != nil {
		api.logger.Error("Failed to delete manufacturer", zap.Error(err))
		json.NewEncoder(w).Encode(resthttp.HttpResponse{Error: "Failed to delete manufacturer"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(output); err != nil {
		api.logger.Error("Failed to encode response", zap.Error(err))
	}
}
