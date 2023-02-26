package internal_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/democracy-tools/countmein/internal"
	"github.com/democracy-tools/countmein/internal/bq"
	"github.com/stretchr/testify/require"
)

func TestAnnouncement(t *testing.T) {

	body, err := json.Marshal(&map[string][]internal.Announcement{"announcments": {{
		UserId:     "dw12f",
		UserDevice: internal.Device{Id: "d123", Type: "tag"},
		SeenDevice: internal.Device{Id: "d234", Type: "iphone 14"},
		Location:   internal.Location{Latitute: 32.05766501361105, Longitude: 34.76640727232065},
		Timestamp:  time.Now().Unix(),
	}}})
	require.NoError(t, err)
	r, err := http.NewRequest(http.MethodGet, "/announcements", bytes.NewReader(body))
	require.NoError(t, err)
	w := httptest.NewRecorder()

	internal.NewHandle(bq.NewInMemoryClient()).Announcements(w, r)

	require.Equal(t, http.StatusCreated, w.Result().StatusCode)
}
