package middleware

import (
	"context"
	"fmt"
	"fulka-api/util"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		fmt.Println("auth", authHeader)
		if authHeader == "" {
			util.WriteJSONResponse(w, http.StatusUnauthorized, "Missing token", "unauthorized", nil)
			return
		}

		tokenString := authHeader
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			util.WriteJSONResponse(w, http.StatusUnauthorized, "Invalid token format", "unauthorized", nil)

			return
		}

		claims := &jwt.StandardClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(util.GetConfig("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			util.WriteJSONResponse(w, http.StatusUnauthorized, "Invalid token", "unauthorized", nil)

			return
		}

		ctx := context.WithValue(r.Context(), "user", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserFromContext(ctx context.Context) (*jwt.StandardClaims, bool) {
	claims, ok := ctx.Value("user").(*jwt.StandardClaims)
	return claims, ok
}
