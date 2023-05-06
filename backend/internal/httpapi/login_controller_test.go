package httpapi

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	_ "the-video-project/backend/configs"
	"the-video-project/backend/internal/command"
	"the-video-project/backend/internal/cookie"
	"the-video-project/backend/internal/models"

	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	commands := []command.Command{
		command.MigrateCommand{},
		command.SeedCommand{},
	}
	for _, cmd := range commands {
		if err := cmd.Run(); err != nil {
			panic(err)
		}
	}
}

func TestApiIsLoggedIn(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/is-logged-in", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	assert.JSONEq(t, w.Body.String(), `{"is-logged-in": false, "status": "success"}`)
}

func TestApiLogin(t *testing.T) {
	router := SetupRouter()

	email := fmt.Sprintf("%s@not-exist.com", uuid.New())
	password := "ultra_secret"
	(&models.User{}).Create(email, password)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/login", jsonBody(map[string]any{
		"email":    email,
		"password": password,
	}))
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 204)
	for _, k := range w.Result().Cookies() {
		if k.Name == cookie.CookieName {
			assert.NotEmpty(t, k.MaxAge)
			assert.NotEmpty(t, k.Value)
		}
	}
}

func jsonBody(jsonMap map[string]any) *bytes.Buffer {
	jsonByte, _ := json.Marshal(jsonMap)
	return bytes.NewBuffer(jsonByte)
}

func TestApiLogout(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/logout", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 204)
	for _, k := range w.Result().Cookies() {
		if k.Name == cookie.CookieName {
			assert.Equal(t, k.MaxAge, -1)
			assert.Equal(t, k.Value, "")
		}
	}
}
