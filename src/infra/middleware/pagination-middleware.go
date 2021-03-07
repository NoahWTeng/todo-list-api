package mdw

import (
	"context"
	"github.com/NoahWTeng/todo-api-go/src/infra/helpers/pagination"
	"net/http"
)

func Pagination(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newPagination := pagination.NewFromRequest(r)
		ctx := context.WithValue(r.Context(), "pagination", newPagination)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
