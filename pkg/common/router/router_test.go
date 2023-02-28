package router

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
	"github.com/pauljamescleary/gomin/pkg/common/config"
	"github.com/pauljamescleary/gomin/pkg/common/handler"
)

var testEchoContext *echo.Echo
var testHandler *handler.Handler
var testConfig *config.Config

func init() {
	testConfig = &config.Config{
		Port:  1323,
		DbUrl: "postgres://test:test@localhost:5435/gomin",
	}
	testHandler = handler.LoadHandlerFromConfig(*testConfig)
	testEchoContext = SetupRouter(testHandler)
}

func makeRequest(method, url string, body interface{}) (*http.Request, *httptest.ResponseRecorder) {
	requestBody, _ := json.Marshal(body)
	req := httptest.NewRequest(method, url, bytes.NewBuffer(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	return req, rec
}
