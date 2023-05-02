package httpapi

import (
	"strconv"
	"the-video-project/backend/internal/db"
	"the-video-project/backend/internal/models"

	"github.com/gin-gonic/gin"
)

type CartController struct{}

func (ctl CartController) ListItem(c *gin.Context) {
	userId := c.GetUint("uid")
	storeId, err := strconv.Atoi(c.Param("storeId"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Store ID 必須是數字",
		})
		return
	}

	cart := models.Cart{}
	tx := db.Conn().Where(models.Cart{
		UserId:  userId,
		StoreId: uint(storeId),
	}).Attrs(models.Cart{
		Items: "",
	}).FirstOrCreate(&cart)
	if tx.Error != nil {
		c.JSON(500, gin.H{
			"message": "寫入失敗",
		})
	} else {
		c.JSON(200, gin.H{
			"items": cart.Items,
		})
	}
}
func (ctl CartController) NewItem(c *gin.Context) {
	// TODO
}
func (ctl CartController) GetItem(c *gin.Context) {
	// TODO
}
func (ctl CartController) UpdateItem(c *gin.Context) {
	// TODO
}
func (ctl CartController) RemoveItem(c *gin.Context) {
	// TODO
}
