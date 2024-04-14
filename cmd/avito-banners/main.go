package main

import (
	"AvitoBanner/internal/config"
	"AvitoBanner/internal/storage/sqlite"
	"fmt"
	"os"
)

func main() {
	// TODO: config
	cfg := config.MustLoad()
	fmt.Println(cfg)

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		os.Exit(1)
	}

	_ = storage
}
