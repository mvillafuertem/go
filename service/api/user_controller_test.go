package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {

	controller := UserController{Router: gin.Default()}
	controller.ping()

	recorder := httptest.NewRecorder()
	response, _ := http.NewRequest("GET", "/ping", nil)
	controller.Router.ServeHTTP(recorder, response)

	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, "{\"message\":\"pong\"}", recorder.Body.String())


}

func TestUserRoute(t *testing.T) {

	controller := UserController{Router: gin.Default()}
	controller.getUser()

	recorder := httptest.NewRecorder()
	response, _ := http.NewRequest("GET", "/user/123", nil)
	controller.Router.ServeHTTP(recorder, response)

	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, "{\"name\":\"Pepe\",\"surname\":\"Pipo\",\"userId\":\"123\"}", recorder.Body.String())


}