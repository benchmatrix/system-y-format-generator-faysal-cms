package authMiddleWare

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"system-y-format-generator-faysal-cms/authMiddleWare/AuthModels"
	"system-y-format-generator-faysal-cms/authMiddleWare/Service"

	"github.com/golang-jwt/jwt"
)

func AuthMiddleWare(Db *sql.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Cache-Control", "no-store")
			w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self' 'unsafe-inline';")

			reqToken := r.Header.Get("Authorization")
			if (reqToken) == "" || &reqToken == nil || !strings.HasPrefix(reqToken, "Bearer") || strings.HasPrefix(reqToken, "Bearer null") {
				w.WriteHeader(http.StatusUnauthorized)
				check := &AuthModels.CustomError{ErrorCode: 401, Status: "Unauthorized", Description: "User is not authorized"}
				json.NewEncoder(w).Encode(check)
				return
			}
			splitToken := strings.Split(reqToken, "Bearer ")
			reqToken = splitToken[1]
			reqToken1 := reqToken
			claims := &AuthModels.Claims{}
			secretKey, err1 := Service.GetSecret(Db)
			if err1 != nil {
				fmt.Println(err1)
			}

			tokenCheck, err := jwt.ParseWithClaims(reqToken1, claims, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(secretKey), nil
			})
			if err != nil {
				fmt.Println(err)
			}

			if !tokenCheck.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				check := &AuthModels.CustomError{ErrorCode: 401, Status: "Unauthorized", Description: "Token not valid, Re-login Required!"}
				json.NewEncoder(w).Encode(check)
				return
			}

			tokenValid := Service.CheckTokenDb(Db, reqToken)

			if (tokenValid) == true {
				next.ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				check := &AuthModels.CustomError{ErrorCode: 401, Status: "Unauthorized", Description: "No token is linked to the user!"}
				json.NewEncoder(w).Encode(check)
				return
			}
		})
	}
}
