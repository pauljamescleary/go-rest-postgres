package router

import (
	"encoding/json"
	"io"
	"net/http"
	"os/user"
	"testing"

	"github.com/pauljamescleary/gomin/pkg/common/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	u := user.User{Name: "test1"}
	req, rec := makeRequest(http.MethodPost, "/users", u)
	c := testEchoContext.NewContext(req, rec)

	if assert.NoError(t, testHandler.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestGetUser(t *testing.T) {
	u := user.User{Name: "test1"}
	req, rec := makeRequest(http.MethodPost, "/users", u)
	c := testEchoContext.NewContext(req, rec)

	if assert.NoError(t, testHandler.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}

	userBytes, _ := io.ReadAll(rec.Body)
	var createdUser *models.User
	_ = json.Unmarshal(userBytes, &createdUser)

	req, rec = makeRequest(http.MethodGet, "/", nil)
	c = testEchoContext.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues(createdUser.ID.String())

	if assert.NoError(t, testHandler.GetUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
