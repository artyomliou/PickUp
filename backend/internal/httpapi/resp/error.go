package resp

import (
	"github.com/gin-gonic/gin"
)

var StdErrorResp = gin.H{
	"message": "發生錯誤",
}

var UriErrorResp = gin.H{
	"message": "URL參數有錯誤",
}

var BodyErrorResp = gin.H{
	"message": "HTTP Body有錯誤",
}

var DbReadErrorResp = gin.H{
	"message": "發生錯誤 (1001)",
}

var DbWriteErrorResp = gin.H{
	"message": "發生錯誤 (1002)",
}
