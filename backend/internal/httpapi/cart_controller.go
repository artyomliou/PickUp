package httpapi

import (
	"log"
	"the-video-project/backend/internal/db"
	"the-video-project/backend/internal/httpapi/resp"
	"the-video-project/backend/internal/models"

	"github.com/gin-gonic/gin"
)

type (
	CartController struct{}
	CartUri        struct {
		StoreId uint `uri:"storeId" binding:"required" valid:"int"`
	}
	CartItemUri struct {
		StoreId uint `uri:"storeId" binding:"required" valid:"int"`
		ItemId  uint `uri:"itemId" binding:"required" valid:"int"`
	}
	NewItemInput struct {
		ProductId uint `binding:"required" valid:"int"`
		Amount    uint `binding:"required" valid:"int"`
		Selects   models.SelectAnswers
		Customs   models.CustomAnswers
	}
	NewItemResponse struct {
		CartItemId uint `binding:"required" valid:"int"`
	}
	UpdateItemInput struct {
		Amount  uint `binding:"required" valid:"int"`
		Selects models.SelectAnswers
		Customs models.CustomAnswers
	}
)

func (ctl CartController) ListItem(c *gin.Context) {
	// input validation
	uri := CartUri{}
	if err := c.BindUri(&uri); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(400, resp.UriErrorResp)
		return
	}

	// db operation
	cart, err := models.GetCart(c.GetUint("uid"), uri.StoreId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(500, resp.StdErrorResp)
		return
	}

	c.AbortWithStatusJSON(200, gin.H{
		"items": cart.Items,
	})
}
func (ctl CartController) NewItem(c *gin.Context) {
	// input validation
	uri := CartUri{}
	if err := c.BindUri(&uri); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(400, resp.UriErrorResp)
		return
	}

	input := NewItemInput{}
	if err := c.Bind(&input); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(400, resp.BodyErrorResp)
		return
	}

	// db operation
	cart, err := models.GetCart(c.GetUint("uid"), uri.StoreId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(500, resp.StdErrorResp)
		return
	}

	item := models.CartItem{
		CartId:    cart.ID,
		ProductId: input.ProductId,
		Amount:    input.Amount,
		Selects:   input.Selects,
		Customs:   input.Customs,
	}
	if tx := db.Conn().Save(&item); tx.Error != nil {
		log.Println(tx.Error)
		c.AbortWithStatusJSON(500, resp.DbWriteErrorResp)
		return
	}

	c.AbortWithStatusJSON(200, NewItemResponse{
		CartItemId: item.ID,
	})
}
func (ctl CartController) GetItem(c *gin.Context) {
	// input validation
	uri := CartItemUri{}
	if err := c.BindUri(&uri); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(400, resp.UriErrorResp)
		return
	}

	// db operation
	item, err := models.GetCartItem(uri.ItemId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(500, resp.DbReadErrorResp)
		return
	}

	c.AbortWithStatusJSON(200, gin.H{
		"item": item,
	})
}
func (ctl CartController) UpdateItem(c *gin.Context) {
	// input validation
	uri := CartItemUri{}
	if err := c.BindUri(&uri); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(400, resp.UriErrorResp)
		return
	}

	newItem := UpdateItemInput{}
	if err := c.Bind(&newItem); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(400, resp.BodyErrorResp)
		return
	}

	// db operation
	oldItem, err := models.GetCartItem(uri.ItemId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(500, resp.DbReadErrorResp)
		return
	}

	oldItem.Amount = newItem.Amount
	oldItem.Selects = newItem.Selects
	oldItem.Customs = newItem.Customs
	if tx := db.Conn().Save(&oldItem); tx.Error != nil {
		log.Println(tx.Error)
		c.AbortWithStatusJSON(500, resp.DbWriteErrorResp)
		return
	}

	c.AbortWithStatus(204)
}
func (ctl CartController) RemoveItem(c *gin.Context) {
	// input validation
	uri := CartItemUri{}
	if err := c.BindUri(&uri); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(400, resp.UriErrorResp)
		return
	}

	// db operation
	if tx := db.Conn().Delete(&models.CartItem{}, uri.ItemId); tx.Error != nil {
		log.Println(tx.Error)
		c.AbortWithStatusJSON(500, resp.DbWriteErrorResp)
		return
	}

	c.AbortWithStatus(204)
}
