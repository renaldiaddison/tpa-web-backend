package middlewares

import (
	"context"
	"net/http"
)

type ctxString string

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		if auth == "" {
			next.ServeHTTP(w, r)
			return
		}

		getAuthToken := auth[len("Bearer "):]

		validate, err := JwtValidate(context.Background(), getAuthToken)
		if err != nil || !validate.Valid {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		customClaim, _ := validate.Claims.(*JwtCustomClaim)

		ctx := context.WithValue(r.Context(), ctxString("JwtValue"), customClaim)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func GetJwtValueData(ctx context.Context) *JwtCustomClaim {
	raw, _ := ctx.Value(ctxString("JwtValue")).(*JwtCustomClaim)
	return raw
}
