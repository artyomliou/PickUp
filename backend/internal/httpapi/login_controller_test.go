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
	"the-video-project/backend/internal/db"
	"the-video-project/backend/internal/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func sendReq(router *gin.Engine, method string, uri string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, uri, nil)
	router.ServeHTTP(w, req)
	return w
}

func sendReqWithHeaders(router *gin.Engine, method string, uri string, headers map[string]string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, uri, nil)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	router.ServeHTTP(w, req)
	return w
}

func sendReqWithBody(router *gin.Engine, method string, uri string, body *bytes.Buffer) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, uri, body)
	req.Header.Set("Content-type", "application/json")
	router.ServeHTTP(w, req)
	return w
}

func jsonBody(jsonMap map[string]any) *bytes.Buffer {
	jsonByte, _ := json.Marshal(jsonMap)
	return bytes.NewBuffer(jsonByte)
}

func TestLoginController(t *testing.T) {
	// Setup
	t.Setenv("DB_DRIVER", "sqlite")
	if err := (&command.MigrateCommand{}).Run(); err != nil {
		panic(err)
	}
	r := SetupRouter()

	// Setup a user
	email := fmt.Sprintf("%s@not-exist.com", uuid.New())
	password := "ultra_secret"
	passwordHash, _ := models.HashPassword(password)
	user := models.User{
		Email:    email,
		Password: passwordHash,
	}
	db.Conn().Save(&user)

	// Tests
	t.Run("/api/is-logged-in", func(t *testing.T) {
		t.Run("false", func(t *testing.T) {
			w := sendReq(r, "GET", "/api/is-logged-in")
			assert.Equal(t, 200, w.Code)
			assert.JSONEq(t, `{"isLoggedIn": false}`, w.Body.String())
		})
		t.Run("true", func(t *testing.T) {
			token, _ := cookie.CreateToken(user.ID, time.Now().AddDate(0, 0, 2))
			w := sendReqWithHeaders(r, "GET", "/api/is-logged-in", map[string]string{
				"Cookie": fmt.Sprintf("auth=%s", token),
			})
			assert.Equal(t, 200, w.Code)
			assert.JSONEq(t, `{"isLoggedIn": true}`, w.Body.String())
		})
	})

	t.Run("/api/login", func(t *testing.T) {
		t.Run("Correct user/pwd", func(t *testing.T) {
			w := sendReqWithBody(r, "POST", "/api/login", jsonBody(map[string]any{
				"email":    email,
				"password": password,
			}))
			assert.Equal(t, 204, w.Code)
			for _, k := range w.Result().Cookies() {
				if k.Name == cookie.CookieName {
					assert.NotEmpty(t, k.MaxAge)
					assert.NotEmpty(t, k.Value)
				}
			}
		})

		t.Run("Incorrect user/pwd", func(t *testing.T) {
			w := sendReqWithBody(r, "POST", "/api/login", jsonBody(map[string]any{
				"email":    "whatever@localhost",
				"password": "12345678",
			}))
			assert.Equal(t, 400, w.Code)
			for _, k := range w.Result().Cookies() {
				if k.Name == cookie.CookieName {
					assert.Empty(t, k.MaxAge)
					assert.Empty(t, k.Value)
				}
			}
		})
	})

	t.Run("/api/logout", func(t *testing.T) {
		w := sendReq(r, "GET", "/api/logout")

		assert.Equal(t, 204, w.Code)
		for _, k := range w.Result().Cookies() {
			if k.Name == cookie.CookieName {
				assert.Equal(t, k.MaxAge, -1)
				assert.Equal(t, k.Value, "")
			}
		}
	})
}
