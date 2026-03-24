package util

import "net/http"

type Pagination struct {
	Page       int  `json:"page"`
	Limit      int  `json:"limit"`
	TotalItems int  `json:"total_items"`
	TotalPages int  `json:"total_pages"`
	HasNext    bool `json:"has_next"`
	HasPrev    bool `json:"has_prev"`
}

type PaginatedData struct {
	Data       any        `json:"data"`
	Pagination Pagination `json:"pagination"`
}

func SendPaginatedData(w http.ResponseWriter, data any, page, limit, totalItems int) {
	// Validate inputs
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	// Calculate total pages
	totalPages := 0
	if limit > 0 && totalItems > 0 {
		totalPages = (totalItems + limit - 1) / limit
	}

	response := PaginatedData{
		Data: data,
		Pagination: Pagination{
			Page:       page,
			Limit:      limit,
			TotalItems: totalItems,
			TotalPages: totalPages,
			HasNext:    page < totalPages,
			HasPrev:    page > 1,
		},
	}

	SuccessResponse(w, response, http.StatusOK)
}
