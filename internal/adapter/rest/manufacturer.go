package rest

import (
	"encoding/json"
	"fmt"
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
	cr.Get("/", api.List)
	cr.Post("/", api.Create)
	cr.Get("/{id}", api.Get)
	cr.Put("/{id}", api.Update)
	cr.Delete("/{id}", api.Delete)
	return cr
}

// writeJSON writes v as a JSON response with the given HTTP status code.
func (api *ManufacturerAPI) writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		api.logger.Error("Failed to encode response", zap.Error(err))
	}
}

func (api *ManufacturerAPI) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "invalid manufacturer ID"})
		return
	}
	req := manufacturer.GetManufacturerRequest{ID: uint32(id)}

	output, err := api.svc.GetByID(ctx, req)
	if err != nil {
		api.logger.Error("Failed to get manufacturer", zap.Error(err))
		api.writeJSON(w, http.StatusInternalServerError, resthttp.HttpResponse{Error: "Failed to get manufacturer"})
		return
	}

	api.writeJSON(w, http.StatusOK, output)
}

func (api *ManufacturerAPI) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req manufacturer.ListManufacturerRequest
	req.FromQueryParams(r.URL.Query())

	output, err := api.svc.List(ctx, req)
	if err != nil {
		api.logger.Error("Failed to list manufacturers", zap.Error(err))
		api.writeJSON(w, http.StatusInternalServerError, resthttp.HttpResponse{Error: "Failed to list manufacturers"})
		return
	}

	api.writeJSON(w, http.StatusOK, output)
}

func (api *ManufacturerAPI) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req manufacturer.CreateManufacturerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "invalid request body"})
		return
	}
	if req.Name == "" {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "name is required"})
		return
	}

	output, err := api.svc.Create(ctx, req)
	if err != nil {
		api.logger.Error("Failed to create manufacturer", zap.Error(err))
		api.writeJSON(w, http.StatusInternalServerError, resthttp.HttpResponse{Error: "Failed to create manufacturer"})
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/manufacturers/%d", output.Data.ID))
	api.writeJSON(w, http.StatusCreated, output)
}

func (api *ManufacturerAPI) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "invalid manufacturer ID"})
		return
	}
	var req manufacturer.UpdateManufacturerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "invalid request body"})
		return
	}
	req.ID = uint32(id)

	output, err := api.svc.Update(ctx, req)
	if err != nil {
		api.logger.Error("Failed to update manufacturer", zap.Error(err))
		api.writeJSON(w, http.StatusInternalServerError, resthttp.HttpResponse{Error: "Failed to update manufacturer"})
		return
	}

	api.writeJSON(w, http.StatusOK, output)
}

func (api *ManufacturerAPI) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "invalid manufacturer ID"})
		return
	}
	var req manufacturer.DeleteManufacturerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "invalid request body"})
		return
	}
	req.ID = uint32(id)

	output, err := api.svc.Delete(ctx, req)
	if err != nil {
		api.logger.Error("Failed to delete manufacturer", zap.Error(err))
		api.writeJSON(w, http.StatusInternalServerError, resthttp.HttpResponse{Error: "Failed to delete manufacturer"})
		return
	}

	api.writeJSON(w, http.StatusOK, output)
}
