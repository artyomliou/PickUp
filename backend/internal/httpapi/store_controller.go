package httpapi

import (
	"log"
	"pick-up/backend/internal/db"
	"pick-up/backend/internal/models"

	"github.com/gin-gonic/gin"
)

type (
	StoreController struct{}
	GetStoreInput   struct {
		Id uint `uri:"id" binding:"required" valid:"number"`
	}
	GetProductInput struct {
		StoreId   uint `uri:"id" binding:"required" valid:"number"`
		ProductId uint `uri:"pid" binding:"required" valid:"number"`
	}
	ListStoreResponse struct {
		Stores []models.Store `json:"stores"`
	}
	GetStoreResponse struct {
		Store models.Store `json:"store"`
	}
	GetProductResponse struct {
		Product models.Product `json:"product"`
	}
)

func (ctl StoreController) ListStore(c *gin.Context) {
	stores := []models.Store{}
	db.Conn().Find(&stores, models.Store{
		Status: models.StoreStatus(models.StoreStatusOpened), // Get opened store
	})
	c.AbortWithStatusJSON(200, ListStoreResponse{
		Stores: stores,
	})
}

func (ctl StoreController) GetStore(c *gin.Context) {
	// input validation
	input := GetStoreInput{}
	if err := c.ShouldBindUri(&input); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "參數錯誤",
		})
	}

	// db operation
	store := models.Store{}
	conn := db.Conn()
	conn = conn.Preload("Categories.Products.SelectQuestions.SelectOptions")
	conn = conn.Preload("Categories.Products.CustomQuestions")
	tx := conn.First(&store, input.Id)
	if tx.Error != nil {
		log.Println(tx.Error)
		c.AbortWithStatusJSON(404, gin.H{
			"message": "指定店家不存在",
		})
	} else {
		c.AbortWithStatusJSON(200, GetStoreResponse{
			Store: store,
		})
	}
}

func (ctl StoreController) GetProduct(c *gin.Context) {
	// input validation
	input := GetProductInput{}
	if err := c.ShouldBindUri(&input); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "參數錯誤",
		})
	}

	// db operation
	product := models.Product{}
	conn := db.Conn()
	conn = conn.Preload("SelectQuestions.SelectOptions")
	conn = conn.Preload("CustomQuestions")
	tx := conn.Where("store_id", input.StoreId).First(&product, input.ProductId)
	if tx.Error != nil {
		log.Println(tx.Error)
		c.AbortWithStatusJSON(404, gin.H{
			"message": "指定商品不存在",
		})
	} else {
		c.AbortWithStatusJSON(200, GetProductResponse{
			Product: product,
		})
	}
}
