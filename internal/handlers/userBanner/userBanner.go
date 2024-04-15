package userBanner

import (
	"AvitoBanner/internal/auth"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"strconv"
)

func GetUserBanner(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./storage/storage_new.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	tagID, err := strconv.Atoi(r.URL.Query().Get("tag_id"))

	if err != nil {
		http.Error(w, "Invalid tag_id", http.StatusBadRequest)
		return
	}

	featureID, err := strconv.Atoi(r.URL.Query().Get("feature_id"))
	if err != nil {
		http.Error(w, "Invalid feature_id", http.StatusBadRequest)
		return
	}

	useLastRevision, err := strconv.ParseBool(r.URL.Query().Get("use_last_revision"))
	if err != nil {
		useLastRevision = false
	}

	banner, err := getUserBannerFromDB(db, tagID, featureID, useLastRevision)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(6)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(banner)
}

func getUserBannerFromDB(db *sql.DB, tagID, featureID int, useLastRevision bool) (*auth.Banner, error) {
	var query string
	if useLastRevision {
		query = `
            SELECT b.id, b.title, b.text, b.url, b.is_active, b.created_at, b.updated_at
            FROM banners b
            INNER JOIN banner_tags bt ON b.id = bt.banner_id
            INNER JOIN banner_features bf ON b.id = bf.banner_id
            WHERE bt.tag_id = ? AND bf.feature_id = ?
            ORDER BY b.updated_at DESC
            LIMIT 1 
        `
	} else {
		query = `
            SELECT b.id, b.title, b.text, b.url, b.is_active, b.created_at, b.updated_at
            FROM banners b
            INNER JOIN banner_tags bt ON b.id = bt.banner_id
            INNER JOIN banner_features bf ON b.id = bf.banner_id
            WHERE bt.tag_id = ? AND bf.feature_id = ?
            ORDER BY b.created_at DESC
            LIMIT 1
        `
	}
	fmt.Println(query)
	row := db.QueryRow(query, tagID, featureID)
	fmt.Println(row)
	var banner auth.Banner
	err := row.Scan(&banner.BannerID, &banner.Title, &banner.Text,
		&banner.URL, &banner.IsActive, &banner.CreatedAt, &banner.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &banner, nil
}
