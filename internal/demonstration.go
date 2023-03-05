package internal

import (
	"encoding/json"
	"net/http"
	"time"
)

func (h *Handle) Demonstrations(w http.ResponseWriter, r *http.Request) {

	count, err := h.bqClient.GetAnnouncementCount(time.Now().Add(time.Hour * (-7)).Unix())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"demonstrations": struct {
		Count int64 `json:"count"`
	}{Count: count}})
}
