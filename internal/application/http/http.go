package http

type HttpResponse struct {
	Error string `json:"error,omitempty"`
}

type ListResponse struct {
	Page    int         `json:"page"`
	Limit   int         `json:"limit"`
	Total   int         `json:"total"`
	Records interface{} `json:"records"`
}

type ListRequest struct {
	Page     int `json:"page" binding:"required"`
	PageSize int `json:"page_size" binding:"required"`
	Limit    int `json:"limit" binding:"required"`

	CreatedAt_Gte *string `json:"created_at_gte" form:"created_at_gte"`
	CreatedAt_Lte *string `json:"created_at_lte" form:"created_at_lte"`
	CreatedBy     *string `json:"created_by" form:"created_by"`

	UpdatedAt_Gte *string `json:"updated_at_gte" form:"updated_at_gte"`
	UpdatedAt_Lte *string `json:"updated_at_lte" form:"updated_at_lte"`
	UpdatedBy     *string `json:"updated_by" form:"updated_by"`

	DeletedAt_Gte *string `json:"deleted_at_gte" form:"deleted_at_gte"`
	DeletedAt_Lte *string `json:"deleted_at_lte" form:"deleted_at_lte"`
	DeletedBy     *string `json:"deleted_by" form:"deleted_by"`
}
