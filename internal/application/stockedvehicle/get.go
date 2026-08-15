package stockedvehicle

import (
	"net/url"
	"strconv"

	"github.com/ducminhgd/intelligent-inventory/internal/application/http"
	"github.com/ducminhgd/intelligent-inventory/internal/domain"
)

type GetStockedVehicleRequest struct {
	ID uint32 `json:"id"`
}

type GetStockedVehicleResponse struct {
	http.HttpResponse
	Data *domain.StockedVehicle `json:"data"`
}

type ListStockedVehicleDataResponse struct {
	Page     int                      `json:"page"`
	PageSize int                      `json:"page_size"`
	Records  []*domain.StockedVehicle `json:"records"`
}

type ListStockedVehicleResponse struct {
	http.ListResponse
	Data ListStockedVehicleDataResponse `json:"data"`
}

type ListStockedVehicleRequest struct {
	http.ListRequest

	ID_In []uint32 `json:"id_in" form:"id_in"`

	ModelID_In []uint32 `json:"model_id_in" form:"model_id_in"`

	VIN string `json:"vin" form:"vin"`

	Name_iLike string `json:"name_ilike" form:"name_ilike"`
}

func (r *ListStockedVehicleRequest) FromQueryParams(values url.Values) {
	if v := values.Get("page"); v != "" {
		if page, err := strconv.Atoi(v); err == nil {
			r.Page = page
		}
	}
	if v := values.Get("page_size"); v != "" {
		if pageSize, err := strconv.Atoi(v); err == nil {
			r.PageSize = pageSize
		}
	}
	if v := values.Get("name_ilike"); v != "" {
		r.Name_iLike = v
	}
	if v := values.Get("vin"); v != "" {
		r.VIN = v
	}
	for _, v := range values["id_in"] {
		if id, err := strconv.ParseUint(v, 10, 32); err == nil {
			r.ID_In = append(r.ID_In, uint32(id))
		}
	}
	for _, v := range values["model_id_in"] {
		if id, err := strconv.ParseUint(v, 10, 32); err == nil {
			r.ModelID_In = append(r.ModelID_In, uint32(id))
		}
	}
}
