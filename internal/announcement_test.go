package internal_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tufin/count-me-in/internal"
)

func TestAnnouncement(t *testing.T) {

	body, err := json.Marshal(&map[string][]internal.AnnouncementRequest{"announcments": {{
		UserId:            "123",
		DeviceId:          "d123",
		SeenDevices:       []string{"d234", "d235", "d236"},
		LocationLatitute:  "32.05766501361105",
		LocationLongitude: "34.76640727232065",
		Timestamp:         time.Now().Unix(),
	}}})
	require.NoError(t, err)
	r, err := http.NewRequest(http.MethodGet, "/announcements", bytes.NewReader(body))
	require.NoError(t, err)
	w := httptest.NewRecorder()

	internal.Announcements(w, r)

	require.Equal(t, http.StatusOK, w.Result().StatusCode)
}
