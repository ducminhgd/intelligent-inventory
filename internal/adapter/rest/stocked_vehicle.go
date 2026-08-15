package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	resthttp "github.com/ducminhgd/intelligent-inventory/internal/application/http"
	"github.com/ducminhgd/intelligent-inventory/internal/application/stockedvehicle"
)

type StockedVehicleAPI struct {
	svc    *stockedvehicle.StockedVehicleService
	logger *zap.Logger
}

func NewStockedVehicleAPI(svc *stockedvehicle.StockedVehicleService, logger *zap.Logger) *StockedVehicleAPI {
	return &StockedVehicleAPI{svc: svc, logger: logger}
}

func (api *StockedVehicleAPI) Router() http.Handler {
	cr := chi.NewRouter()
	cr.Get("/", api.List)
	cr.Post("/", api.Create)
	cr.Get("/{id}", api.Get)
	cr.Put("/{id}", api.Update)
	cr.Delete("/{id}", api.Delete)
	return cr
}

// writeJSON writes v as a JSON response with the given HTTP status code.
func (api *StockedVehicleAPI) writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		api.logger.Error("Failed to encode response", zap.Error(err))
	}
}

func (api *StockedVehicleAPI) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "invalid stocked vehicle ID"})
		return
	}
	req := stockedvehicle.GetStockedVehicleRequest{ID: uint32(id)}

	output, err := api.svc.GetByID(ctx, req)
	if err != nil {
		api.logger.Error("Failed to get stocked vehicle", zap.Error(err))
		api.writeJSON(w, http.StatusInternalServerError, resthttp.HttpResponse{Error: "Failed to get stocked vehicle"})
		return
	}

	api.writeJSON(w, http.StatusOK, output)
}

func (api *StockedVehicleAPI) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req stockedvehicle.ListStockedVehicleRequest
	req.FromQueryParams(r.URL.Query())

	output, err := api.svc.List(ctx, req)
	if err != nil {
		api.logger.Error("Failed to list stocked vehicles", zap.Error(err))
		api.writeJSON(w, http.StatusInternalServerError, resthttp.HttpResponse{Error: "Failed to list stocked vehicles"})
		return
	}

	api.writeJSON(w, http.StatusOK, output)
}

func (api *StockedVehicleAPI) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req stockedvehicle.CreateStockedVehicleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "invalid request body"})
		return
	}
	if req.VIN == "" {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "vin is required"})
		return
	}
	if req.ModelID == 0 {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "model_id is required"})
		return
	}
	if req.Name == "" {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "name is required"})
		return
	}
	if req.Price <= 0 {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "price must be positive"})
		return
	}
	if req.Action != "" && !req.Action.Valid() {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "invalid action"})
		return
	}

	output, err := api.svc.Create(ctx, req)
	if err != nil {
		api.logger.Error("Failed to create stocked vehicle", zap.Error(err))
		api.writeJSON(w, http.StatusInternalServerError, resthttp.HttpResponse{Error: "Failed to create stocked vehicle"})
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/stocked-vehicles/%d", output.Data.ID))
	api.writeJSON(w, http.StatusCreated, output)
}

func (api *StockedVehicleAPI) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "invalid stocked vehicle ID"})
		return
	}
	var req stockedvehicle.UpdateStockedVehicleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "invalid request body"})
		return
	}
	req.ID = uint32(id)
	if req.VIN == "" {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "vin is required"})
		return
	}
	if req.ModelID == 0 {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "model_id is required"})
		return
	}
	if req.Name == "" {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "name is required"})
		return
	}
	if req.Price <= 0 {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "price must be positive"})
		return
	}
	if req.Action != "" && !req.Action.Valid() {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "invalid action"})
		return
	}

	output, err := api.svc.Update(ctx, req)
	if err != nil {
		api.logger.Error("Failed to update stocked vehicle", zap.Error(err))
		api.writeJSON(w, http.StatusInternalServerError, resthttp.HttpResponse{Error: "Failed to update stocked vehicle"})
		return
	}

	api.writeJSON(w, http.StatusOK, output)
}

func (api *StockedVehicleAPI) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "invalid stocked vehicle ID"})
		return
	}
	var req stockedvehicle.DeleteStockedVehicleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		api.writeJSON(w, http.StatusBadRequest, resthttp.HttpResponse{Error: "invalid request body"})
		return
	}
	req.ID = uint32(id)

	output, err := api.svc.Delete(ctx, req)
	if err != nil {
		api.logger.Error("Failed to delete stocked vehicle", zap.Error(err))
		api.writeJSON(w, http.StatusInternalServerError, resthttp.HttpResponse{Error: "Failed to delete stocked vehicle"})
		return
	}

	api.writeJSON(w, http.StatusOK, output)
}
