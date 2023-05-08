package httpapi_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"pick-up/backend/internal/db"
	"pick-up/backend/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func sendReq(router *gin.Engine, method string, uri string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, uri, nil)
	router.ServeHTTP(w, req)
	return w
}

func SendReqWithOrigin(router *gin.Engine, origin string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/helloworld", nil)
	req.Header.Set("Origin", origin)
	router.ServeHTTP(w, req)
	return w
}

func SendReqWithBody(router *gin.Engine, method string, uri string, body *bytes.Buffer) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, uri, body)
	req.Header.Set("Content-type", "application/json")
	router.ServeHTTP(w, req)
	return w
}

func JsonBody(jsonMap any) *bytes.Buffer {
	jsonByte, _ := json.Marshal(jsonMap)
	return bytes.NewBuffer(jsonByte)
}

func SetupUser() (*models.User, string, string, error) {
	email := fmt.Sprintf("%s@not-exist.com", uuid.New())
	password := "ultra_secret"
	passwordHash, _ := models.HashPassword(password)
	user := models.User{
		Email:    email,
		Password: passwordHash,
	}
	tx := db.Conn().Save(&user)
	if tx.Error != nil {
		return nil, "", "", tx.Error
	}
	return &user, email, password, nil
}

func SetupStore() (*models.Store, error) {
	store := models.Store{
		Name:     "某間商店",
		Pic:      "http://localhost:8080/store-pic/12345678.jpg",
		Status:   models.StoreStatus(models.StoreStatusOpened),
		OpenedAt: "09:00",
		ClosedAt: "18:00",
	}
	tx := db.Conn().Save(&store)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &store, nil
}

func SetupProduct(storeId uint) (*models.Product, error) {
	product := models.Product{
		StoreId: storeId,
		Name:    "烏龍茶",
		Price:   30,
	}
	tx := db.Conn().Save(&product)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &product, nil
}

func SetupSelectQuestion(storeId uint) (*models.SelectQuestion, []uint, error) {
	questionName := "冰量"
	isRequired := true
	atMost := 1
	optNames := []string{
		"正常冰",
		"少冰",
		"微冰",
		"去冰",
		"完全去冰",
	}

	question := models.SelectQuestion{
		StoreId:    storeId,
		Name:       questionName,
		IsRequired: isRequired,
		AtMost:     uint(atMost),
	}
	if tx := db.Conn().Save(&question); tx.Error != nil {
		return nil, nil, tx.Error
	}

	options := make([]uint, len(optNames))
	for _, optName := range optNames {
		optModel := models.SelectOption{
			SelectQuestionId: question.ID,
			Name:             optName,
			Price:            0,
		}
		if tx := db.Conn().Save(&optModel); tx.Error != nil {
			return nil, nil, tx.Error
		}
		options = append(options, optModel.ID)
	}

	return &question, options, nil
}

func AssociateProductWithSelectQuestion(product *models.Product, question *models.SelectQuestion) error {
	return db.Conn().Model(&product).Association("SelectQuestions").Append(question)
}

func SetupCustomQuestion(storeId uint) (*models.CustomQuestion, error) {
	questionName := "有額外需求嗎？"
	placeholder := "請在這邊寫您的需求"

	question := models.CustomQuestion{
		StoreId:     storeId,
		Name:        questionName,
		Placeholder: placeholder,
	}
	if tx := db.Conn().Save(&question); tx.Error != nil {
		return nil, tx.Error
	}

	return &question, nil
}

func AssociateProductWithCustomQuestion(product *models.Product, question *models.CustomQuestion) error {
	return db.Conn().Model(&product).Association("CustomQuestions").Append(question)
}

func SetupCart(userId uint, storeId uint, productId uint) (*models.Cart, error) {
	cart := models.Cart{
		StoreId: storeId,
		UserId:  userId,
	}
	if err := db.Conn().Create(&cart).Error; err != nil {
		return nil, err
	}

	item := models.CartItem{
		CartId:    cart.ID,
		ProductId: productId,
		Amount:    5,
		Selects:   models.SelectAnswers{},
		Customs:   models.CustomAnswers{},
	}
	if err := db.Conn().Create(&item).Error; err != nil {
		return nil, err
	}

	return &cart, nil
}
