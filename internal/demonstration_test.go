package internal_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/democracy-tools/countmein/internal"
	"github.com/democracy-tools/countmein/internal/bq"
	"github.com/stretchr/testify/require"
)

func TestHandle_Demonstrations(t *testing.T) {

	r, err := http.NewRequest(http.MethodGet, "/demonstrations", nil)
	require.NoError(t, err)
	w := httptest.NewRecorder()

	internal.NewHandle(bq.NewInMemoryClient()).Demonstrations(w, r)

	require.Equal(t, http.StatusOK, w.Result().StatusCode)
}
