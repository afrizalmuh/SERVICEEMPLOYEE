package middlewares

import (
	"context"
	"serviceemployee/helpers"
	"net/http"
)


func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		accessToken := r.Header.Get("Authorization")

		if accessToken == "" {
			helpers.Response(w, 401, "anauthorized", nil)
			return
		}

		employee, err := helpers.ValidateToken(accessToken)
		if err != nil {
			helpers.Response(w, 401, err.Error(), nil)
			return
		}

		ctx := context.WithValue(r.Context(), "employeeinfo", employee)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}