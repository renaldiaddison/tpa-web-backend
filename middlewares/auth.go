package middlewares

import (
	"context"
	"net/http"

	"github.com/renaldiaddison/tpa-web-backend/service"
)

type authString string

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		autho := r.Header.Get("Authorization")

		if autho == "" {
			next.ServeHTTP(w, r)
			return
		}

		bearer := "Bearer "
		autho = autho[len(bearer):]

		validate, err := service.JwtValidate(context.Background(), autho)
		if err != nil || !validate.Valid {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		customClaim, _ := validate.Claims.(*service.JwtCustomClaim)

		ctx := context.WithValue(r.Context(), authString("auth"), customClaim)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func CtxValue(ctx context.Context) *service.JwtCustomClaim {
	raw, _ := ctx.Value(authString("auth")).(*service.JwtCustomClaim)
	return raw
}
