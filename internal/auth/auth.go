package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

var mySigningKey = []byte("avitorest")

const (
	UserToken   = "user_token"
	AdminToken  = "admin_token"
	UserRole    = "user"
	AdminRole   = "admin"
	TokenExpire = 24 * time.Hour // Время жизни токена
)

type Banner struct {
	BannerID  int    `json:"banner_id"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	URL       string `json:"url"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func CreateToken(role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["role"] = role
	claims["expired at"] = time.Now().Add(TokenExpire).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("something went wrong: %s", err.Error())
	}

	return tokenString, nil
}

// Функция для проверки JWT-токена на авторизацию
func isTokenValid(tokenString string, role string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil || !token.Valid {
		return false
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["role"] != role || claims["role"] != AdminRole {
		return false
	}
	return true
}

func CheckAuthorization(next http.HandlerFunc, role string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		var isValidToken bool
		if role == AdminRole || role == UserRole {
			isValidToken = isTokenValid(tokenString, role)
		} else {
			http.Error(w, "Invalid role", http.StatusInternalServerError)
			return
		}

		if !isValidToken {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
