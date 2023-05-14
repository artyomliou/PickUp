package httpapi

import (
	"log"
	"net/http"
	"pick-up/backend/internal/cookie"
	"pick-up/backend/internal/models"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	LoginController struct{}
	LoginBody       struct {
		Email    string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
	}
)

func (ctl LoginController) IsLoggedIn(c *gin.Context) {
	tokenStr, err := c.Cookie(cookie.CookieName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"isLoggedIn": false,
		})
		return
	}

	_, err = cookie.ParseTokenForUid(tokenStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"isLoggedIn": false,
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"isLoggedIn": true,
	})
}

func (ctl LoginController) Login(c *gin.Context) {
	body := LoginBody{}
	if err := c.Bind(&body); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "資料格式不符",
		})
		return
	}

	user, err := models.GetUserByEmail(body.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "帳號或密碼錯誤",
		})
		return
	}
	if !user.IsPasswordMatched(body.Password) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "帳號或密碼錯誤",
		})
		return
	}

	expiredAt := time.Now().AddDate(0, 0, 2)
	token, err := cookie.CreateToken(user.ID, expiredAt)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(500, gin.H{
			"message": "伺服器錯誤",
		})
		return
	}

	c.SetCookie(cookie.CookieName, token, int(expiredAt.Unix()-time.Now().Unix()), "/", "", false, true) // todo
	c.AbortWithStatusJSON(http.StatusNoContent, nil)
}

func (ctl LoginController) Logout(c *gin.Context) {
	token := ""
	c.SetCookie(cookie.CookieName, token, -1, "/", "", false, true)
	c.AbortWithStatusJSON(http.StatusNoContent, nil)
}
