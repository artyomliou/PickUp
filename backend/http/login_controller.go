package http

import (
	"net/http"
	"the-video-project/backend/http/middlewares"
	"the-video-project/backend/models"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginController struct{}

func (l LoginController) IsLoggedIn(c *gin.Context) {
	cookie, err := c.Cookie(middlewares.CookieName)
	c.JSON(http.StatusOK, gin.H{
		"status":       "success",
		"is-logged-in": len(cookie) > 0 && err == nil,
	})
}

func (l LoginController) Login(c *gin.Context) {
	type ReqBody struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var body ReqBody
	user := &models.User{}
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
		token := middlewares.CreateToken(user.ID, expiredAt)
		c.SetCookie(middlewares.CookieName, token, int(expiredAt.Unix()-time.Now().Unix()), "/", "", false, true) // todo
		c.JSON(http.StatusNoContent, nil)
	}

}

func (l LoginController) Logout(c *gin.Context) {
	token := ""
	c.SetCookie(middlewares.CookieName, token, -1, "/", "", false, true)
	c.JSON(http.StatusNoContent, nil)
}
