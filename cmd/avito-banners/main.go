package main

import (
	"AvitoBanner/internal/auth"
	"AvitoBanner/internal/config"
	"AvitoBanner/internal/handlers/banner"
	"AvitoBanner/internal/handlers/bannerId"
	"AvitoBanner/internal/handlers/login"
	"AvitoBanner/internal/handlers/userBanner"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// TODO: config
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// TODO: server
	router := mux.NewRouter()

	router.HandleFunc("/login", login.Login).Methods("POST")

	router.Handle("/user_banner",
		auth.IsAuthorized(userBanner.GetUserBanner)).Methods("GET")
	router.Handle("/banner",
		auth.CheckAuthorization(banner.GetAllBanners, auth.AdminRole)).Methods("GET")
	router.Handle("/banner",
		auth.CheckAuthorization(banner.CreateBanner, auth.AdminRole)).Methods("POST")
	router.Handle("/banner/{id}",
		auth.CheckAuthorization(bannerId.PatchBanner, auth.AdminRole)).Methods("PATCH")
	router.Handle("/banner/{id}",
		auth.CheckAuthorization(bannerId.DeleteBanner, auth.AdminRole)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
