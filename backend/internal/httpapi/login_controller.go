package httpapi

import (
	"log"
	"net/http"
	"the-video-project/backend/internal/cookie"
	"the-video-project/backend/internal/models"
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
	user := models.User{}
	if err := c.Bind(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "資料格式不符",
		})
	}
	if user.GetByEmail(body.Email); user.ID == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "帳號或密碼錯誤",
		})
	}
	if !user.IsPasswordMatched(body.Password) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "帳號或密碼錯誤",
		})
	}

	expiredAt := time.Now().AddDate(0, 0, 2)
	token, err := cookie.CreateToken(user.ID, expiredAt)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(500, gin.H{
			"message": "伺服器錯誤",
		})
	}
	c.SetCookie(cookie.CookieName, token, int(expiredAt.Unix()-time.Now().Unix()), "/", "", false, true) // todo
	c.AbortWithStatusJSON(http.StatusNoContent, nil)

}

func (ctl LoginController) Logout(c *gin.Context) {
	token := ""
	c.SetCookie(cookie.CookieName, token, -1, "/", "", false, true)
	c.AbortWithStatusJSON(http.StatusNoContent, nil)
}
