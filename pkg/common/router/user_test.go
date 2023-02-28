package router

import (
	"net/http"
	"os/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	u := user.User{Name: "test1"}
	req, rec := makeRequest(http.MethodPost, "/users", u)
	c := testEchoContext.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, testHandler.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}
