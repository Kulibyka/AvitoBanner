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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var cl Client
	json.NewDecoder(r.Body).Decode(&cl)
	checkLogin(cl)
}

func checkLogin(cl Client) string {

	if _, exist := clients[cl.Username]; !exist || clients[cl.Username][0] != cl.Password {
		fmt.Println("NOT CORRECT")
		err := "error"
		return err
	}

	validToken, err := auth.CreateToken(clients[cl.Username][1])
	fmt.Println(validToken)

	if err != nil {
		fmt.Println(err)
	}

	return validToken
}
