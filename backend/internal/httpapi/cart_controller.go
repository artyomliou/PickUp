package httpapi

import (
	"log"
	"pick-up/backend/internal/db"
	"pick-up/backend/internal/httpapi/resp"
	"pick-up/backend/internal/models"

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

	UpdateItemInput struct {
		Amount        uint                 `json:"amount" binding:"required" valid:"int"`
		SelectAnswers models.SelectAnswers `json:"selectAnswers" binding:"required"`
		CustomAnswers models.CustomAnswers `json:"customAnswers" binding:"required"`
	}
	ListItemResponse struct {
		Items []*models.CartItem `json:"items"`
	}
	NewItemResponse struct {
		CartItemId uint `json:"cartItemId"`
	}
	GetItemResponse struct {
		Item models.CartItem `json:"item"`
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

	c.AbortWithStatusJSON(200, ListItemResponse{
		Items: cart.Items,
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
		CartId:        cart.ID,
		ProductId:     input.ProductId,
		Amount:        input.Amount,
		SelectAnswers: input.Selects,
		CustomAnswers: input.Customs,
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

	c.AbortWithStatusJSON(200, GetItemResponse{
		Item: *item,
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
	oldItem.SelectAnswers = newItem.SelectAnswers
	oldItem.CustomAnswers = newItem.CustomAnswers
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
