package httpapi

import (
	"fmt"
	"testing"
	"the-video-project/backend/internal/command"
	"the-video-project/backend/internal/db"
	"the-video-project/backend/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestStoreController(t *testing.T) {
	// Setup
	t.Setenv("DB_DRIVER", "sqlite")
	if err := (&command.MigrateCommand{}).Run(); err != nil {
		panic(err)
	}
	r := SetupRouter()

	// Setup a store
	store := models.Store{
		Name:     "某間商店",
		Pic:      "http://localhost:8080/store-pic/12345678.jpg",
		Status:   models.StoreStatus(models.StoreStatusOpened),
		OpenedAt: "09:00",
		ClosedAt: "18:00",
	}
	tx := db.Conn().Save(&store)
	if tx.Error != nil {
		t.Fatal(tx.Error)
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
