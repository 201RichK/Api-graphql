package security

import (
	"fmt"
	"log"
	"net/http"

	"github.com/201RichK/graphql/utils"
	"github.com/dgrijalva/jwt-go"
)

type APIResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Handle(next http.Handler) http.Handler {
	conf := utils.Conf()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("token")

		if tokenString != "" {
			token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodRSA); ok {
					return nil, fmt.Errorf("Unexpected sign in method %v", token.Header["alg"])
				}

				return []byte(conf.Security.PrivateKey), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid && claims != nil {
				log.Printf("jwt authenticated OK (app: %s)", claims["app"])

				next.ServeHTTP(w, r)
			}
		} else {
			log.Println("token not ookay")
		}
	})
}
