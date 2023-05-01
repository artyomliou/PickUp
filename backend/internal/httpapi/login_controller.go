package httpapi

import (
	"net/http"
	"the-video-project/backend/internal/cookie"
	"the-video-project/backend/internal/models"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginController struct{}

func (ctl LoginController) IsLoggedIn(c *gin.Context) {
	cookie, err := c.Cookie(cookie.CookieName)
	c.JSON(http.StatusOK, gin.H{
		"status":       "success",
		"is-logged-in": len(cookie) > 0 && err == nil,
	})
}

func (ctl LoginController) Login(c *gin.Context) {
	type ReqBody struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var body ReqBody
	user := models.User{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "資料格式不符",
		})
		return
	}
	if user.GetByEmail(body.Email); user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "帳號或密碼錯誤",
		})
		return
	}
	if !user.IsPasswordMatched(body.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "帳號或密碼錯誤",
		})
		return
	} else {
		expiredAt := time.Now().AddDate(0, 0, 2)
		token := cookie.CreateToken(user.ID, expiredAt)
		c.SetCookie(cookie.CookieName, token, int(expiredAt.Unix()-time.Now().Unix()), "/", "", false, true) // todo
		c.JSON(http.StatusNoContent, nil)
	}

}

func (ctl LoginController) Logout(c *gin.Context) {
	token := ""
	c.SetCookie(cookie.CookieName, token, -1, "/", "", false, true)
	c.JSON(http.StatusNoContent, nil)
}
