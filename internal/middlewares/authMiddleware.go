package middlewares

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	SecretKey string
}

func NewAuthMiddleware(secretKey string) *AuthMiddleware {
	return &AuthMiddleware{
		SecretKey: secretKey,
	}
}

func (mw *AuthMiddleware) ValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized)+": empty authorization header", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		tokenString := tokenParts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(mw.SecretKey), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
		if err != nil {
			if err != nil {
				log.Printf("Error parsing token: %v", err)
				http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusUnauthorized)
				return
			}
		}

		var userID float64
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			userID = claims["userId"].(float64)
		} else {
			http.Error(w, "Invalid token payload", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userId", int(userID))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
