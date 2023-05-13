package httpapi_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func sendReq(router *gin.Engine, method string, uri string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, uri, nil)
	router.ServeHTTP(w, req)
	return w
}

func SendReqWithOrigin(router *gin.Engine, origin string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/helloworld", nil)
	req.Header.Set("Origin", origin)
	router.ServeHTTP(w, req)
	return w
}

func SendReqWithBody(router *gin.Engine, method string, uri string, body *bytes.Buffer) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, uri, body)
	req.Header.Set("Content-type", "application/json")
	router.ServeHTTP(w, req)
	return w
}

func JsonBody(jsonMap any) *bytes.Buffer {
	jsonByte, _ := json.Marshal(jsonMap)
	return bytes.NewBuffer(jsonByte)
}
