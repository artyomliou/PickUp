package httpapi_test

import (
	"testing"
	"the-video-project/backend/internal/httpapi"

	"github.com/stretchr/testify/assert"
)

func TestOrigin(t *testing.T) {
	r := httpapi.SetupRouter()

	t.Run("null", func(t *testing.T) {
		w := SendReqWithOrigin(r, "null")
		assert.Equal(t, "null", w.Header().Get("Access-Control-Allow-Origin"))
		assert.Equal(t, 200, w.Code)
	})
	t.Run("localhost", func(t *testing.T) {
		w := SendReqWithOrigin(r, "http://localhost")
		assert.Equal(t, "http://localhost", w.Header().Get("Access-Control-Allow-Origin"))
		assert.Equal(t, 200, w.Code)
	})
	t.Run("localhost:8080", func(t *testing.T) {
		w := SendReqWithOrigin(r, "http://localhost:8080")
		assert.Equal(t, "http://localhost:8080", w.Header().Get("Access-Control-Allow-Origin"))
		assert.Equal(t, 200, w.Code)
	})
	t.Run("127.0.0.1", func(t *testing.T) {
		w := SendReqWithOrigin(r, "http://127.0.0.1")
		assert.Equal(t, "http://127.0.0.1", w.Header().Get("Access-Control-Allow-Origin"))
		assert.Equal(t, 200, w.Code)
	})
	t.Run("google.com", func(t *testing.T) {
		w := SendReqWithOrigin(r, "https://google.com")
		assert.Equal(t, "", w.Header().Get("Access-Control-Allow-Origin"))
		assert.Equal(t, 403, w.Code)
	})
}

func TestCredendials(t *testing.T) {
	r := httpapi.SetupRouter()

	t.Run("valid domain", func(t *testing.T) {
		w := SendReqWithOrigin(r, "http://localhost:8080")
		assert.Equal(t, "true", w.Header().Get("Access-Control-Allow-Credentials"))
	})
	t.Run("invalid domain", func(t *testing.T) {
		w := SendReqWithOrigin(r, "https://google.com")
		assert.Equal(t, "", w.Header().Get("Access-Control-Allow-Credentials"))
	})
}
