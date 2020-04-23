package main

import (
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/rs/cors"
	"net/http"
	"time"
)

type Error struct {
	Code string `json:"code"`
	Description string `json:"description"`
}

func main() {
	loginSecret := "login-secret"
	zendeskSecret := "zendesk-secret"

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil || cookie == nil {
			writeErrorResponse(w, "00-00", "Token not passed", http.StatusUnauthorized)
			return
		}
		
		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}

			return []byte(loginSecret), nil
		})

		if err != nil {
			writeErrorResponse(w, "00-01", "Token is invalid", http.StatusUnauthorized)
			return
		}

		var loginClaims jwt.MapClaims
		var ok bool
		if loginClaims, ok = token.Claims.(jwt.MapClaims); !ok || !token.Valid {
			writeErrorResponse(w, "00-01", "Token is invalid", http.StatusUnauthorized)
			return
		}

		var sub, email string
		if sub, ok = loginClaims["sub"].(string); !ok || sub == "" {
			writeErrorResponse(w, "00-02", "\"sub\" claim is missed", http.StatusUnprocessableEntity)
			return
		}
		if email, ok = loginClaims["email"].(string); !ok || email == "" {
			writeErrorResponse(w, "00-03", "\"email\" claim is missed", http.StatusUnprocessableEntity)
			return
		}

		zendeskClaims := jwt.MapClaims{}
		zendeskClaims["email"] = email
		zendeskClaims["name"] = email
		zendeskClaims["external_id"] = sub
		zendeskClaims["iat"] = time.Now().UTC().Unix()

		zendeskToken := jwt.NewWithClaims(jwt.SigningMethodHS256, zendeskClaims)
		zendeskTokenString, _ := zendeskToken.SignedString([]byte(zendeskSecret))
		writeResponse(w, zendeskTokenString)
	})

	handler := getCors().Handler(mux)

	http.ListenAndServe("example.com:8001", handler)
}

func getCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"http://example.com"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
}

func writeErrorResponse(w http.ResponseWriter, code, description string, statusCode int) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-type", "application/json")

	encoder := json.NewEncoder(w)

	err := new(Error)
	err.Code = code
	err.Description = description

	encoder.Encode(err)
}

func writeResponse(w http.ResponseWriter, token string) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")

	body := struct {
		Token string `json:"token"`
	}{}
	body.Token = token

	encoder := json.NewEncoder(w)
	encoder.Encode(&body)
}