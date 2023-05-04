package httpapi

import (
	"log"
	"the-video-project/backend/internal/db"
	"the-video-project/backend/internal/httpapi/resp"
	"the-video-project/backend/internal/models"
	"the-video-project/backend/internal/snowflake"

	"github.com/gin-gonic/gin"
)

type (
	OrderController  struct{}
	CreateOrderInput struct {
		StoreId uint `form:"store_id" binding:"required" valid:"int"`
	}
)

func (ctl OrderController) CreateOrder(c *gin.Context) {
	// input validation
	input := CreateOrderInput{}
	if err := c.Bind(&input); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(400, resp.StdErrorResp)
	}

	// get cart
	cart := models.Cart{
		UserId:  c.GetUint("uid"),
		StoreId: input.StoreId,
	}
	if tx := db.Conn().First(&cart); tx.Error != nil {
		log.Println(tx.Error)
		c.AbortWithStatusJSON(404, resp.StdErrorResp)
	}

	// save order
	order := models.Order{
		StoreId: cart.StoreId,
		UserId:  cart.UserId,
		ID:      snowflake.Generate(),
		Price:   cart.Items.CalculateTotalPrice(),
		Items:   cart.Items,
		Status:  0,
	}
	if tx := db.Conn().Save(&order); tx.Error != nil {
		log.Println(tx.Error)
		c.AbortWithStatusJSON(500, resp.StdErrorResp)
	}

	c.AbortWithStatusJSON(200, gin.H{
		"order": order,
	})
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
