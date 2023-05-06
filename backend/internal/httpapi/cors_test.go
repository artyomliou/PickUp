package httpapi

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func sendReqWithOrigin(router *gin.Engine, origin string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/helloworld", nil)
	req.Header.Set("Origin", origin)
	router.ServeHTTP(w, req)
	return w
}

func TestOrigin(t *testing.T) {
	r := SetupRouter()

	t.Run("null", func(t *testing.T) {
		w := sendReqWithOrigin(r, "null")
		assert.Equal(t, "null", w.Header().Get("Access-Control-Allow-Origin"))
		assert.Equal(t, 200, w.Code)
	})
	t.Run("localhost", func(t *testing.T) {
		w := sendReqWithOrigin(r, "http://localhost")
		assert.Equal(t, "http://localhost", w.Header().Get("Access-Control-Allow-Origin"))
		assert.Equal(t, 200, w.Code)
	})
	t.Run("localhost:8080", func(t *testing.T) {
		w := sendReqWithOrigin(r, "http://localhost:8080")
		assert.Equal(t, "http://localhost:8080", w.Header().Get("Access-Control-Allow-Origin"))
		assert.Equal(t, 200, w.Code)
	})
	t.Run("127.0.0.1", func(t *testing.T) {
		w := sendReqWithOrigin(r, "http://127.0.0.1")
		assert.Equal(t, "http://127.0.0.1", w.Header().Get("Access-Control-Allow-Origin"))
		assert.Equal(t, 200, w.Code)
	})
	t.Run("google.com", func(t *testing.T) {
		w := sendReqWithOrigin(r, "https://google.com")
		assert.Equal(t, "", w.Header().Get("Access-Control-Allow-Origin"))
		assert.Equal(t, 403, w.Code)
	})
}

func TestCredendials(t *testing.T) {
	r := SetupRouter()

	t.Run("valid domain", func(t *testing.T) {
		w := sendReqWithOrigin(r, "http://localhost:8080")
		assert.Equal(t, "true", w.Header().Get("Access-Control-Allow-Credentials"))
	})
	t.Run("invalid domain", func(t *testing.T) {
		w := sendReqWithOrigin(r, "https://google.com")
		assert.Equal(t, "", w.Header().Get("Access-Control-Allow-Credentials"))
	})
}
