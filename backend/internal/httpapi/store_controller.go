package httpapi

import (
	"the-video-project/backend/internal/db"
	"the-video-project/backend/internal/models"

	"github.com/gin-gonic/gin"
)

type StoreController struct{}

func (ctl StoreController) ListStore(c *gin.Context) {
	stores := []models.Store{}
	db.Conn().Find(&stores)
	c.AbortWithStatusJSON(200, gin.H{
		"stores": stores,
	})
}
func (ctl StoreController) GetStore(c *gin.Context) {
	store := models.Store{}
	tx := db.Conn().First(&store, c.Param("storeId"))
	if tx.Error != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"message": "指定店家不存在",
		})
	} else {
		c.AbortWithStatusJSON(200, gin.H{
			"store": store,
		})
	}
}
