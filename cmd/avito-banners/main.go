package main

import (
	"AvitoBanner/internal/auth"
	"AvitoBanner/internal/config"
	"AvitoBanner/internal/handlers/banner"
	"AvitoBanner/internal/handlers/bannerId"
	"AvitoBanner/internal/handlers/login"
	"AvitoBanner/internal/handlers/userBanner"
	"AvitoBanner/internal/storage/sqlite"
	"fmt"
	"github.com/gorilla/mux"
	"os"
)

func main() {
	// TODO: config
	cfg := config.MustLoad()
	fmt.Println(cfg)

	storage, err := sqlite.Open(cfg.StoragePath)
	if err != nil {
		fmt.Errorf("%w", err)
		os.Exit(1)
	}

	_ = storage

	// TODO: server
	router := mux.NewRouter()

	router.HandleFunc("/login", login.Login).Methods("POST")

	router.Handle("/user_banner",
		auth.CheckAuthorization(userBanner.GetUserBanner, auth.UserRole)).Methods("GET")
	router.Handle("/banner",
		auth.CheckAuthorization(banner.GetAllBanners, auth.AdminRole)).Methods("GET")
	router.Handle("/banner",
		auth.CheckAuthorization(banner.CreateBanner, auth.AdminRole)).Methods("POST")
	router.Handle("/banner/{id}",
		auth.CheckAuthorization(bannerId.PatchBanner, auth.AdminRole)).Methods("PATCH")
	router.Handle("/banner/{id}",
		auth.CheckAuthorization(bannerId.DeleteBanner, auth.AdminRole)).Methods("DELETE")

}
