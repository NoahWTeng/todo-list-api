package mdw

import (
	"context"
	"github.com/NoahWTeng/todo-api-go/src/app/users"
	"github.com/NoahWTeng/todo-api-go/src/infra/helpers/errors"
	"github.com/NoahWTeng/todo-api-go/src/infra/helpers/response"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

//type Claims struct {
//	Name  string `json:"name"`
//	Email string `json:"email"`
//	jwt.StandardClaims
//}

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		jwtKey := users.JwtKey
		token := r.Header.Get("Authorization")
		removeBearer := strings.ReplaceAll(token, "Bearer ", "")

		tkn, _ := jwt.ParseWithClaims(removeBearer, &users.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if claims, ok := tkn.Claims.(*users.Claims); ok && tkn.Valid {
			ctx := context.WithValue(r.Context(), "user", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			err := errors.Unauthorized("")
			response.Error(w, r, err.Status, err.Message)
		}

	})
}
