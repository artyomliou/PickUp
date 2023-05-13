package httpapi_test

import (
	"fmt"
	"pick-up/backend/internal/command"
	"pick-up/backend/internal/httpapi"
	"pick-up/backend/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoreController(t *testing.T) {
	// Setup
	t.Setenv("DB_DRIVER", "sqlite")
	if err := (&command.MigrateCommand{}).Run(); err != nil {
		panic(err)
	}
	r := httpapi.SetupRouter()

	// Setup a store
	store, err := models.SeedStore()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("/api/stores", func(t *testing.T) {
		w := sendReq(r, "GET", "/api/stores")
		assert.Equal(t, 200, w.Code)
	})

	t.Run("/api/stores/:id", func(t *testing.T) {
		t.Run("Non-exist", func(t *testing.T) {
			w := sendReq(r, "GET", fmt.Sprintf("/api/stores/%d", 99999))
			assert.Equal(t, 404, w.Code)
		})

		t.Run("Exist", func(t *testing.T) {
			w := sendReq(r, "GET", fmt.Sprintf("/api/stores/%d", store.ID))
			assert.Equal(t, 200, w.Code)
		})
	})
}
