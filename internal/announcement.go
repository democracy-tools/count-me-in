package internal

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type AnnouncementRequest struct {
	UserId            string   `json:"user-id"`
	DeviceId          string   `json:"device-id"`
	SeenDevices       []string `json:"seen-devices"`
	LocationLatitute  string   `json:"location-latitute"`
	LocationLongitude string   `json:"location-longitude"`
	Timestamp         int64    `json:"timestamp"`
}

func Announcements(w http.ResponseWriter, r *http.Request) {

	var announcments map[string][]AnnouncementRequest
	err := json.NewDecoder(r.Body).Decode(&announcments)
	if err != nil {
		log.Infof("failed to decode request announcments with %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
