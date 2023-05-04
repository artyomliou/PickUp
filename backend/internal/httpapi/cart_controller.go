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
	CartController struct{}
	CartInput      struct {
		StoreId uint `form:"storeId" binding:"required" valid:"int"`
	}
	ItemInput struct {
		StoreId uint `form:"storeId" binding:"required" valid:"int"`
		ItemId  uint `form:"itemId" binding:"required" valid:"int"`
	}
)

func (ctl CartController) ListItem(c *gin.Context) {
	// input validation
	input := CartInput{}
	if err := c.BindUri(&input); err != nil {
		c.AbortWithStatusJSON(400, resp.StdErrorResp)
	}

	// db operation
	cart, err := models.GetCart(c.GetUint("uid"), input.StoreId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(500, resp.StdErrorResp)
	}

	c.AbortWithStatusJSON(200, gin.H{
		"items": cart.Items,
	})
}
func (ctl CartController) NewItem(c *gin.Context) {
	// input validation
	input := CartInput{}
	if err := c.BindUri(&input); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(400, resp.StdErrorResp)
	}

	item := models.CartItem{}
	if err := c.Bind(&item); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(400, resp.StdErrorResp)
	}

	// db operation
	cart, err := models.GetCart(c.GetUint("uid"), input.StoreId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(500, resp.StdErrorResp)
	}

	item.ID = snowflake.Generate()
	item.CartId = cart.ID
	if tx := db.Conn().Save(&item); tx.Error != nil {
		log.Println(tx.Error)
		c.AbortWithStatusJSON(500, resp.StdErrorResp)
	}

	c.AbortWithStatusJSON(200, gin.H{
		"cartItemid": item.ID,
	})
}
func (ctl CartController) GetItem(c *gin.Context) {
	// input validation
	input := ItemInput{}
	if err := c.BindUri(&input); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(400, resp.StdErrorResp)
	}

	// db operation
	item, err := models.GetCartItem(input.ItemId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(500, resp.StdErrorResp)
	}

	c.AbortWithStatusJSON(200, gin.H{
		"item": item,
	})
}
func (ctl CartController) UpdateItem(c *gin.Context) {
	// input validation
	input := ItemInput{}
	if err := c.BindUri(&input); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(400, resp.StdErrorResp)
	}

	newItem := models.CartItem{}
	if err := c.Bind(&newItem); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(400, resp.StdErrorResp)
	}

	// db operation
	oldItem, err := models.GetCartItem(input.ItemId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(500, resp.StdErrorResp)
	}

	oldItem.Amount = newItem.Amount
	oldItem.Selects = newItem.Selects
	oldItem.Customs = newItem.Customs
	if tx := db.Conn().Save(&oldItem); tx.Error != nil {
		log.Println(tx.Error)
		c.AbortWithStatusJSON(500, resp.StdErrorResp)
	}

	c.AbortWithStatus(204)
}
func (ctl CartController) RemoveItem(c *gin.Context) {
	// input validation
	input := ItemInput{}
	if err := c.BindUri(&input); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(400, resp.StdErrorResp)
	}

	// db operation
	if tx := db.Conn().Delete(&models.CartItem{}, input.ItemId); tx.Error != nil {
		log.Println(tx.Error)
		c.AbortWithStatusJSON(500, resp.StdErrorResp)
	}

	c.AbortWithStatus(204)
}
