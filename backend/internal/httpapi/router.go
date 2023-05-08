package httpapi

import (
	"flag"
	"pick-up/backend/internal/cookie"

	"github.com/gin-gonic/gin"
)

func SetupRouter() (router *gin.Engine) {
	router = gin.Default()
	router.Use(corsMiddleware())

	api := router.Group("/api")

	// Testing endpoint
	if flag.Lookup("test.v") != nil {
		testingRoutes(api)
	}

	// Public endpoint
	{
		loginRoutes(api)
		storeRoutes(api)
	}

	// Protected endpoint
	authorized := api.Group("/")
	authorized.Use(cookie.AuthRequired())
	{
		cartRoutes(authorized)
		orderRoutes(authorized)
	}

	return
}

func testingRoutes(rg *gin.RouterGroup) {
	rg.GET("/helloworld", func(ctx *gin.Context) {
		ctx.AbortWithStatusJSON(200, gin.H{
			"message": "hello world",
		})
	})
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
	rg.GET("/stores/:id", ctl.GetStore)
}

func cartRoutes(rg *gin.RouterGroup) {
	ctl := new(CartController)
	rg.GET("/cart/:storeId/items", ctl.ListItem)
	rg.POST("/cart/:storeId/items", ctl.NewItem)
	rg.GET("/cart/:storeId/items/:itemId", ctl.GetItem)
	rg.PUT("/cart/:storeId/items/:itemId", ctl.UpdateItem)
	rg.DELETE("/cart/:storeId/items/:itemId", ctl.RemoveItem)
}

func orderRoutes(rg *gin.RouterGroup) {
	ctl := new(OrderController)
	rg.POST("/orders", ctl.CreateOrder)
	rg.GET("/orders/:orderId/status", ctl.GetOrderStatus)
	rg.GET("/orders/:orderId", ctl.GetOrder)
}
