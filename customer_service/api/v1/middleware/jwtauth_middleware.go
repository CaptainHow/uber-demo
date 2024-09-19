package middlewares

import (
	"net/http"
	"strings"

	"github.com/uber-demo/customer/api/v1/helper"
)

func AuthenticationMiddleware()  func (next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				helper.RespondWithError(w, http.StatusUnauthorized, "Missing Authentication Token!")
				return
			}
			tokenParts := strings.Split(tokenString, " ")
			if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
				helper.RespondWithError(w, http.StatusUnauthorized, "Invalid Authentication Token!")
				return
			}
			
			tokenString = tokenParts[1]

			claims, err := helper.VerifyToken(tokenString)

			if err != nil {
				helper.RespondWithError(w, http.StatusUnauthorized, "Invalid Authentication Token!")
				return
			}

			w.Header().Set("cust_id", claims["cust_id"].(string))
			next.ServeHTTP(w, r)
		})
	}
}