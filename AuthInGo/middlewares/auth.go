package middlewares

import (
	env "AuthInGo/config/env"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func JwtAuthMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader,"Bearer"){
			http.Error(w, "Invalid Authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader,"Bearer")

		if tokenString == ""{
			http.Error(w,"Token is empty",http.StatusUnauthorized)
			return 
		}

		claims := jwt.MapClaims{}

		_,err := jwt.ParseWithClaims(tokenString,claims,func (token *jwt.Token)(interface{},error){

			return []byte(env.GetString("JWT_SECRET","TOKEN")),nil
		})

		if err!= nil {
			http.Error(w,"Invalid token" + err.Error(),http.StatusUnauthorized)
			return 
		}

		UserId,okId := claims["id"].(float64)

		UserEmail,okEmail := claims["email"]

		if !okId || !okEmail {
			http.Error(w,"Invalid claims",http.StatusUnauthorized)
			return 
		}

		fmt.Println("UserId:",UserId,"Email",UserEmail)

		ctx := context.WithValue(r.Context(),"UserId",int64(UserId))
		ctx = context.WithValue(ctx,"UserEmail",UserEmail)

		next.ServeHTTP(w,r.WithContext(ctx))

	})
}