package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	resthttp "github.com/ducminhgd/intelligent-inventory/internal/application/http"
	"github.com/ducminhgd/intelligent-inventory/internal/application/model"
)

type ModelAPI struct {
	svc    *model.ModelService
	logger *zap.Logger
}

func NewModelAPI(svc *model.ModelService, logger *zap.Logger) *ModelAPI {
	return &ModelAPI{svc: svc, logger: logger}
}

func (api *ModelAPI) Router() http.Handler {
	cr := chi.NewRouter()
	cr.Get("/models/{id}", api.Get)
	cr.Get("/models", api.List)
	cr.Post("/models", api.Create)
	cr.Put("/models/{id}", api.Update)
	cr.Delete("/models/{id}", api.Delete)
	return cr
}

// writeJSON writes v as a JSON response with the given HTTP status code.
func (api *ModelAPI) writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		api.logger.Error("Failed to encode response", zap.Error(err))
	}
}

func (api *ModelAPI) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "invalid model ID"})
		return
	}
	req := model.GetModelRequest{ID: uint32(id)}

	output, err := api.svc.GetByID(ctx, req)
	if err != nil {
		api.logger.Error("Failed to get model", zap.Error(err))
		api.writeJSON(w, http.StatusInternalServerError, resthttp.HttpResponse{Error: "Failed to get model"})
		return
	}

	api.writeJSON(w, http.StatusOK, output)
}

func (api *ModelAPI) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req model.ListModelRequest
	req.FromQueryParams(r.URL.Query())

	output, err := api.svc.List(ctx, req)
	if err != nil {
		api.logger.Error("Failed to list models", zap.Error(err))
		api.writeJSON(w, http.StatusInternalServerError, resthttp.HttpResponse{Error: "Failed to list models"})
		return
	}

	api.writeJSON(w, http.StatusOK, output)
}

func (api *ModelAPI) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req model.CreateModelRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "invalid request body"})
		return
	}
	if req.ManufacturerID == 0 {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "manufacturer_id is required"})
		return
	}
	if req.Name == "" {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "name is required"})
		return
	}

	output, err := api.svc.Create(ctx, req)
	if err != nil {
		api.logger.Error("Failed to create model", zap.Error(err))
		api.writeJSON(w, http.StatusInternalServerError, resthttp.HttpResponse{Error: "Failed to create model"})
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/models/%d", output.Data.ID))
	api.writeJSON(w, http.StatusCreated, output)
}

func (api *ModelAPI) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "invalid model ID"})
		return
	}
	var req model.UpdateModelRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "invalid request body"})
		return
	}
	req.ID = uint32(id)
	if req.ManufacturerID == 0 {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "manufacturer_id is required"})
		return
	}
	if req.Name == "" {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "name is required"})
		return
	}

	output, err := api.svc.Update(ctx, req)
	if err != nil {
		api.logger.Error("Failed to update model", zap.Error(err))
		api.writeJSON(w, http.StatusInternalServerError, resthttp.HttpResponse{Error: "Failed to update model"})
		return
	}

	api.writeJSON(w, http.StatusOK, output)
}

func (api *ModelAPI) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "invalid model ID"})
		return
	}
	var req model.DeleteModelRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "invalid request body"})
		return
	}
	req.ID = uint32(id)

	output, err := api.svc.Delete(ctx, req)
	if err != nil {
		api.logger.Error("Failed to delete model", zap.Error(err))
		api.writeJSON(w, http.StatusInternalServerError, resthttp.HttpResponse{Error: "Failed to delete model"})
		return
	}

	api.writeJSON(w, http.StatusOK, output)
}
