package httpapi_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"pick-up/backend/internal/command"
	"pick-up/backend/internal/cookie"
	"pick-up/backend/internal/httpapi"
	"pick-up/backend/internal/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCartController(t *testing.T) {
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
	store, err := SetupStore()
	if err != nil {
		t.Fatal(err)
	}
	selectQuestion, options, err := SetupSelectQuestion(store.ID)
	if err != nil {
		t.Fatal(err)
	}
	customQuestion, err := SetupCustomQuestion(store.ID)
	if err != nil {
		t.Fatal(err)
	}

	// Setup a product
	product, err := SetupProduct(store.ID)
	if err != nil {
		t.Fatal(err)
	}
	if AssociateProductWithSelectQuestion(product, selectQuestion); err != nil {
		t.Fatal(err)
	}
	if AssociateProductWithCustomQuestion(product, customQuestion); err != nil {
		t.Fatal(err)
	}

	// Tests
	t.Run("Unauthorized", func(t *testing.T) {
		routes := [](map[string]string){
			{
				"method": "GET",
				"uri":    fmt.Sprintf("/api/cart/%d/items", store.ID),
			},
			{
				"method": "POST",
				"uri":    fmt.Sprintf("/api/cart/%d/items", store.ID),
			},
			{
				"method": "GET",
				"uri":    fmt.Sprintf("/api/cart/%d/items/%d", store.ID, 1234),
			},
			{
				"method": "PUT",
				"uri":    fmt.Sprintf("/api/cart/%d/items/%d", store.ID, 1234),
			},
			{
				"method": "DELETE",
				"uri":    fmt.Sprintf("/api/cart/%d/items/%d", store.ID, 1234),
			},
		}
		for _, route := range routes {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(route["method"], route["uri"], nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, 401, w.Code)
		}
	})

	t.Run("ListItem", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/cart/%d/items", store.ID), nil)
		req.Header.Set("Cookie", cookie)
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("NewItem,GetItem,UpdateItem,DeleteItem", func(t *testing.T) {
		var cartItemId uint

		// NewItem
		{
			method := "POST"
			uri := fmt.Sprintf("/api/cart/%d/items", store.ID)
			body := JsonBody(httpapi.NewItemInput{
				ProductId: product.ID,
				Amount:    1,
				Selects: models.SelectAnswers{
					{
						SelectQuestionId: selectQuestion.ID,
						Options:          []uint{options[0]},
					},
				},
				Customs: models.CustomAnswers{
					{
						CustomQuestionId: customQuestion.ID,
						Text:             "沒有拉",
					},
				},
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(method, uri, body)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Cookie", cookie)
			r.ServeHTTP(w, req)

			assert.Equal(t, 200, w.Code)

			// Get CartItemId from response for following sub-tests
			resp := httpapi.NewItemResponse{}
			err := json.Unmarshal(w.Body.Bytes(), &resp)
			if err != nil {
				t.Fatal(err)
			}
			cartItemId = resp.CartItemId
		}

		uri := fmt.Sprintf("/api/cart/%d/items/%d", store.ID, cartItemId)

		// GetItem
		{
			method := "GET"

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(method, uri, nil)
			req.Header.Set("Cookie", cookie)
			r.ServeHTTP(w, req)

			assert.Equal(t, 200, w.Code)
		}

		// UpdateItem
		{
			method := "PUT"
			body := JsonBody(httpapi.UpdateItemInput{
				Amount: 2,
				SelectAnswers: models.SelectAnswers{
					{
						SelectQuestionId: selectQuestion.ID,
						Options:          []uint{options[0]},
					},
				},
				CustomAnswers: models.CustomAnswers{
					{
						CustomQuestionId: customQuestion.ID,
						Text:             "真的沒有拉",
					},
				},
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(method, uri, body)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Cookie", cookie)
			r.ServeHTTP(w, req)

			assert.Equal(t, 204, w.Code)
		}

		// DeleteItem
		{
			method := "DELETE"

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(method, uri, nil)
			req.Header.Set("Cookie", cookie)
			r.ServeHTTP(w, req)

			assert.Equal(t, 204, w.Code)
		}
	})
}
