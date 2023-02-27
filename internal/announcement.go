package internal

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/democracy-tools/countmein/internal/bq"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type AnnouncementDB struct {
	Id                string  `bigquery:"id"`
	UserId            string  `bigquery:"user_id"`
	UserDeviceId      string  `bigquery:"user_device_id"`
	UserDeviceType    string  `bigquery:"user_device_type"`
	SeenDeviceId      string  `bigquery:"seen_device_id"`
	SeenDeviceType    string  `bigquery:"seen_device_type"`
	LocationLatitute  float64 `bigquery:"location_latitute"`
	LocationLongitude float64 `bigquery:"location_longitude"`
	UserTimestamp     int64   `bigquery:"user_timestamp"`
	ServerTimestamp   int64   `bigquery:"server_timestamp"`
}

type Announcement struct {
	UserId     string   `json:"user_id"`
	UserDevice Device   `json:"device_id"`
	SeenDevice Device   `json:"seen_device"`
	Location   Location `json:"location"`
	Timestamp  int64    `json:"timestamp"`
}

type Location struct {
	Latitute  float64 `json:"latitute"`
	Longitude float64 `json:"longitude"`
}

type Device struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type Handle struct{ bqClient bq.Client }

func NewHandle(bqClient bq.Client) *Handle {

	return &Handle{bqClient: bqClient}
}

func (h *Handle) Announcements(w http.ResponseWriter, r *http.Request) {

	announcments, code := getAnnouncements(r)
	if code != http.StatusOK {
		w.WriteHeader(code)
		return
	}

	err := h.bqClient.Insert(bq.TableAnnouncement, toDBAnnouncements(announcments))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func getAnnouncements(r *http.Request) ([]*Announcement, int) {

	var announcments map[string][]*Announcement
	err := json.NewDecoder(r.Body).Decode(&announcments)
	if err != nil {
		log.Infof("failed to decode announcments request with %v", err)
		return nil, http.StatusBadRequest
	}

	res, ok := announcments["announcements"]
	if !ok {
		log.Info("no 'announcements' key found in request")
		return nil, http.StatusBadRequest
	}

	return res, http.StatusOK
}

func toDBAnnouncements(announcments []*Announcement) []*AnnouncementDB {

	res := []*AnnouncementDB{}
	for _, currAnnouncement := range announcments {
		res = append(res, toDBAnnouncement(currAnnouncement))
	}

	return res
}

func toDBAnnouncement(announcement *Announcement) *AnnouncementDB {

	return &AnnouncementDB{
		Id:                uuid.NewString(),
		UserId:            announcement.UserId,
		UserDeviceId:      announcement.UserDevice.Id,
		UserDeviceType:    announcement.UserDevice.Type,
		SeenDeviceId:      announcement.SeenDevice.Id,
		SeenDeviceType:    announcement.SeenDevice.Type,
		LocationLatitute:  announcement.Location.Latitute,
		LocationLongitude: announcement.Location.Longitude,
		UserTimestamp:     announcement.Timestamp,
		ServerTimestamp:   time.Now().Unix(),
	}
}
