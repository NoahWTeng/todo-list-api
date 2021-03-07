package pagination

import (
	"net/http"
	"strconv"
)

var (
	limit = 10
	page  = 1
)

// Pages represents a paginated list of data items.
type Pages struct {
	Page       int         `json:"page"`
	Pages      int         `json:"pages"`
	Limit      int         `json:"limit"`
	TotalCount int         `json:"totalCount"`
	Items      interface{} `json:"items"`
}

func NewFromRequest(r *http.Request) *Pages {
	page := parseInt(r.URL.Query().Get("page"), page)
	limit := parseInt(r.URL.Query().Get("limit"), limit)
	return &Pages{
		Page:  page,
		Limit: limit,
	}
}

// parseInt parses a string into an integer. If parsing is failed, defaultValue will be returned.
func parseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if result, err := strconv.Atoi(value); err == nil {
		return result
	}
	return defaultValue
}

func Update(p *Pages, total int) {
	pageCount := -1

	if total >= 0 {
		pageCount = (total + p.Limit - 1) / p.Limit
		if p.Page > pageCount {
			p.Page = pageCount
		}
	}

	p.Pages = pageCount
	p.TotalCount = total
}
