package model

import (
	"net/url"
	"strconv"

	"github.com/ducminhgd/intelligent-inventory/internal/application/http"
	"github.com/ducminhgd/intelligent-inventory/internal/domain"
)

type GetModelRequest struct {
	ID uint32 `json:"id"`
}

type GetModelResponse struct {
	http.HttpResponse
	Data *domain.Model `json:"data"`
}

type ListModelDataResponse struct {
	Page     int             `json:"page"`
	PageSize int             `json:"page_size"`
	Records  []*domain.Model `json:"records"`
}

type ListModelResponse struct {
	http.ListResponse
	Data ListModelDataResponse `json:"data"`
}

type ListModelRequest struct {
	http.ListRequest

	ID_In []uint32 `json:"id_in" form:"id_in"`

	ManufacturerID_In []uint32 `json:"manufacturer_id_in" form:"manufacturer_id_in"`

	Name_iLike string `json:"name_ilike" form:"name_ilike"`
}

func (r *ListModelRequest) FromQueryParams(values url.Values) {
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
	for _, v := range values["id_in"] {
		if id, err := strconv.ParseUint(v, 10, 32); err == nil {
			r.ID_In = append(r.ID_In, uint32(id))
		}
	}
	for _, v := range values["manufacturer_id_in"] {
		if id, err := strconv.ParseUint(v, 10, 32); err == nil {
			r.ManufacturerID_In = append(r.ManufacturerID_In, uint32(id))
		}
	}
}
