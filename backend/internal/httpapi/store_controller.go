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
	ListStoreResponse struct {
		Stores []models.Store `json:"stores"`
	}
	GetStoreResponse struct {
		Store models.Store `json:"store"`
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
	tx := db.Conn().First(&store, input.Id)
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
	// TODO categories
}
