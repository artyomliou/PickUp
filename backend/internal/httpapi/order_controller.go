package httpapi

import (
	"the-video-project/backend/internal/db"
	"the-video-project/backend/internal/models"

	"github.com/gin-gonic/gin"
)

type OrderController struct{}

func (ctl OrderController) CreateOrder(c *gin.Context) {
	// TODO
}

func (ctl OrderController) GetOrderStatus(c *gin.Context) {
	order := &models.Order{}
	tx := db.Conn().Find(&order, c.Param("orderId"))
	if tx.Error != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"message": "指定店家不存在",
		})
	}

	c.AbortWithStatusJSON(200, gin.H{
		"status": order.Status,
	})
}
func (ctl OrderController) GetOrder(c *gin.Context) {
	order := &models.Order{}
	tx := db.Conn().Find(&order, c.Param("orderId"))
	if tx.Error != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"message": "指定店家不存在",
		})
	}

	c.AbortWithStatusJSON(200, gin.H{
		"order": order,
	})
}
