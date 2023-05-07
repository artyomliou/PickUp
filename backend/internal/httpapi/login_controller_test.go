package httpapi_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	_ "the-video-project/backend/configs"
	"the-video-project/backend/internal/command"
	"the-video-project/backend/internal/cookie"
	"the-video-project/backend/internal/httpapi"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoginController(t *testing.T) {
	// Setup
	t.Setenv("DB_DRIVER", "sqlite")
	if err := (&command.MigrateCommand{}).Run(); err != nil {
		panic(err)
	}
	r := httpapi.SetupRouter()

	// Setup a user
	user, email, password, err := SetupUser()
	if err != nil {
		t.Fatal(err)
	}

	token, _ := cookie.CreateToken(user.ID, time.Now().AddDate(0, 0, 2))
	cookieStr := cookie.MakeCookieString(token)

	// Tests
	t.Run("/api/is-logged-in", func(t *testing.T) {
		t.Run("false", func(t *testing.T) {
			w := sendReq(r, "GET", "/api/is-logged-in")
			assert.Equal(t, 200, w.Code)
			assert.JSONEq(t, `{"isLoggedIn": false}`, w.Body.String())
		})
		t.Run("true", func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/api/is-logged-in", nil)
			req.Header.Set("Cookie", cookieStr)
			r.ServeHTTP(w, req)

			assert.Equal(t, 200, w.Code)
			assert.JSONEq(t, `{"isLoggedIn": true}`, w.Body.String())
		})
	})

	t.Run("/api/login", func(t *testing.T) {
		t.Run("Correct user/pwd", func(t *testing.T) {
			w := SendReqWithBody(r, "POST", "/api/login", JsonBody(map[string]any{
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
			w := SendReqWithBody(r, "POST", "/api/login", JsonBody(map[string]any{
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
