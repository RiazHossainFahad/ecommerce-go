package util

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type Pagination struct {
	Page           int    `json:"page"`
	Limit          int    `json:"limit"`
	TotalItems     int    `json:"total_items"`
	TotalPages     int    `json:"total_pages"`
	HasNext        bool   `json:"has_next"`
	HasPrev        bool   `json:"has_prev"`
	CurrentPageURL string `json:"current_page_url"`
	NextPageURL    string `json:"next_page_url,omitempty"`
	PrevPageURL    string `json:"prev_page_url,omitempty"`
	FirstPageURL   string `json:"first_page_url"`
	LastPageURL    string `json:"last_page_url,omitempty"`
}

type PaginatedData struct {
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

type PaginationBuilder struct {
	request     *http.Request
	data        interface{}
	page        int
	limit       int
	totalItems  int
	baseURL     string
	queryParams url.Values
	path        string
}

func NewPaginationBuilder(r *http.Request, data interface{}, page, limit, totalItems int) *PaginationBuilder {
	// Copy query params to avoid modifying original
	params := r.URL.Query()
	params.Del("page") // Remove page to avoid duplication

	return &PaginationBuilder{
		request:     r,
		data:        data,
		page:        page,
		limit:       limit,
		totalItems:  totalItems,
		baseURL:     getBaseURL(r),
		queryParams: params,
		path:        r.URL.Path,
	}
}

func (b *PaginationBuilder) Build() PaginatedData {
	// Validate and set defaults
	if b.page < 1 {
		b.page = 1
	}
	if b.limit < 1 {
		b.limit = 10
	}

	// Calculate total pages
	totalPages := 0
	if b.limit > 0 && b.totalItems > 0 {
		totalPages = (b.totalItems + b.limit - 1) / b.limit
	}

	// Build pagination object
	pagination := Pagination{
		Page:           b.page,
		Limit:          b.limit,
		TotalItems:     b.totalItems,
		TotalPages:     totalPages,
		HasNext:        b.page < totalPages,
		HasPrev:        b.page > 1,
		CurrentPageURL: b.buildPageURL(b.page),
		FirstPageURL:   b.buildPageURL(1),
	}

	// Add last page URL if total pages > 0
	if totalPages > 0 {
		pagination.LastPageURL = b.buildPageURL(totalPages)
	}

	// Add next and prev URLs if they exist
	if b.page < totalPages {
		pagination.NextPageURL = b.buildPageURL(b.page + 1)
	}
	if b.page > 1 {
		pagination.PrevPageURL = b.buildPageURL(b.page - 1)
	}

	return PaginatedData{
		Data:       b.data,
		Pagination: pagination,
	}
}

func (b *PaginationBuilder) buildPageURL(page int) string {
	// Create copy of query params
	params := url.Values{}
	for k, v := range b.queryParams {
		params[k] = v
	}

	// Add page parameter only if page > 1
	if page > 1 {
		params.Set("page", strconv.Itoa(page))
	}

	// Build URL with query string if params exist
	if len(params) > 0 {
		return fmt.Sprintf("%s%s?%s", b.baseURL, b.path, params.Encode())
	}
	return fmt.Sprintf("%s%s", b.baseURL, b.path)
}

// Helper function to get base URL from request
func getBaseURL(r *http.Request) string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	return fmt.Sprintf("%s://%s", scheme, r.Host)
}
