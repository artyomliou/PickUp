package httpapi_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"pick-up/backend/internal/command"
	"pick-up/backend/internal/cookie"
	"pick-up/backend/internal/httpapi"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOrderController(t *testing.T) {
	// Setup
	t.Setenv("DB_DRIVER", "sqlite")
	if err := (&command.MigrateCommand{}).Run(); err != nil {
		panic(err)
	}
	r := httpapi.SetupRouter()

	// Setup a user
	user, _, _, err := SetupUser()
	if err != nil {
		t.Fatal(err)
	}

	// Setup token
	token, err := cookie.CreateToken(user.ID, time.Now().AddDate(0, 0, 2))
	if err != nil {
		t.Fatal(err)
	}
	cookie := cookie.MakeCookieString(token)

	// Setup a store
	dummyStore, err := SetupStore()
	if err != nil {
		t.Fatal(err)
	}
	store, err := SetupStore()
	if err != nil {
		t.Fatal(err)
	}

	// Setup a product
	product, err := SetupProduct(store.ID)
	if err != nil {
		t.Fatal(err)
	}

	// Setup a cart with a item(product)
	_, err = SetupCart(user.ID, store.ID, product.ID)
	if err != nil {
		t.Fatal(err)
	}

	// Tests
	t.Run("Unauthorized", func(t *testing.T) {
		routes := [](map[string]string){
			{
				"method": "POST",
				"uri":    fmt.Sprintf("/api/orders"),
			},
			{
				"method": "GET",
				"uri":    fmt.Sprintf("/api/orders/%d/status", 1234),
			},
			{
				"method": "GET",
				"uri":    fmt.Sprintf("/api/orders/%d", 1234),
			},
		}
		for _, route := range routes {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(route["method"], route["uri"], nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, 401, w.Code)
		}
	})

	t.Run("CreateOrder:Unvisited store", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/orders", JsonBody(httpapi.CreateOrderInput{
			StoreId: dummyStore.ID,
		}))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Cookie", cookie)
		r.ServeHTTP(w, req)

		assert.Equal(t, 424, w.Code)
	})

	t.Run("CreateOrder,GetOrderStatus,GetOrder", func(t *testing.T) {

		// Create order
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/orders", JsonBody(httpapi.CreateOrderInput{
			StoreId: store.ID,
		}))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Cookie", cookie)
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		// Get orderId from response for following sub-tests
		resp := httpapi.CreateOrderResponse{}
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		if err != nil {
			t.Fatal(err, w.Body.String())
		}
		orderId := resp.Order.ID

		// GetOrderStatus
		{
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("/api/orders/%d/status", orderId), nil)
			req.Header.Set("Cookie", cookie)
			r.ServeHTTP(w, req)

			assert.Equal(t, 200, w.Code)
		}

		// GetOrder
		{
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("/api/orders/%d", orderId), nil)
			req.Header.Set("Cookie", cookie)
			r.ServeHTTP(w, req)

			assert.Equal(t, 200, w.Code)
		}
	})
}
