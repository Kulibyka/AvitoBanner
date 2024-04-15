package login

import (
	"AvitoBanner/internal/auth"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var m map[string][2]string
var clients = map[string][2]string{"admin": {"ad_pass", "admin"},
	"user": {"us_pass", "user"}}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cl Client
	err := json.NewDecoder(r.Body).Decode(&cl)
	if err != nil {
		fmt.Errorf("err: %w", err)
	}
	checkLogin(cl, w)
}

func checkLogin(cl Client, w http.ResponseWriter) string {
	if _, exist := clients[cl.Username]; !exist || clients[cl.Username][0] != cl.Password {
		json.NewEncoder(w).Encode(map[string]string{"verdict": "wrong data"})
		err := "error"
		return err
	}

	validToken, err := auth.CreateToken(clients[cl.Username][1])

	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(map[string]string{"verdict": validToken})
	return validToken
}
