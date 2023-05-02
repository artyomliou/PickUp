package httpapi

import (
	"the-video-project/backend/internal/cookie"

	"github.com/gin-gonic/gin"
)

func SetupRouter() (router *gin.Engine) {
	router = gin.Default()
	router.Use(corsMiddleware())

	// Public endpoint
	api := router.Group("/api")
	{
		loginRoutes(api)
		storeRoutes(api)

	}

	// Protected endpoint
	authorized := api.Group("/")
	authorized.Use(cookie.AuthRequired())
	{
		cartRoutes(api)
		orderRoutes(api)
	}

	return
}

func loginRoutes(rg *gin.RouterGroup) {
	ctl := new(LoginController)
	rg.GET("/is-logged-in", ctl.IsLoggedIn)
	rg.POST("/login", ctl.Login)
	rg.GET("/logout", ctl.Logout)
}

func storeRoutes(rg *gin.RouterGroup) {
	ctl := new(StoreController)
	rg.GET("/stores", ctl.ListStore)
	rg.GET("/stores/{storeId}", ctl.GetStore)
}

func cartRoutes(rg *gin.RouterGroup) {
	ctl := new(CartController)
	rg.GET("/carts/{storeId}/items", ctl.ListItem)
	rg.POST("/carts/{storeId}/items", ctl.NewItem)
	rg.GET("/carts/{storeId}/items/{itemId}", ctl.GetItem)
	rg.PUT("/carts/{storeId}/items/{itemId}", ctl.UpdateItem)
	rg.DELETE("/carts/{storeId}/items/{itemId}", ctl.RemoveItem)
}

func orderRoutes(rg *gin.RouterGroup) {
	ctl := new(OrderController)
	rg.POST("/orders", ctl.CreateOrder)
	rg.GET("/orders/{orderId}/status", ctl.GetOrderStatus)
	rg.GET("/orders/{orderId}", ctl.GetOrder)
}
