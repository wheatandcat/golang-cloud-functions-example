package app

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/rs/xid"
)

func SaveItem(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	client, err := getFireBaseClient(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	guid := xid.New()
	item := Item{
		Status: "OK",
		Day:    time.Now().In(jst).Format("2006-01-02"),
	}
	_, err = client.Collection("results").Doc(guid.String()).Set(ctx, item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	res, err := json.Marshal(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(res)
}
