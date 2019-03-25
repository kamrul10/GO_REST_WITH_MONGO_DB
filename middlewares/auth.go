
package middlewares

import (
	"log"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/matryer/respond.v1"
	"net/http"
	"context"
	"os"
  
)

func JwtMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		jwt_string := os.Getenv("JWT_SECRET")
		jwt_secret := []byte(jwt_string)
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				 return jwt_secret, nil
		})
	
		if err != nil {
			respond.With(w, r, http.StatusUnauthorized, err)
		} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), "authUser",claims)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		} else {
				log.Printf("Invalid JWT Token")
				respond.With(w, r, http.StatusUnauthorized, "Invalid")
		}
		
})

	
}
