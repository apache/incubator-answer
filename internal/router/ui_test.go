package router

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUIRouter_Register(t *testing.T) {
	r := gin.Default()

	NewUIRouter().Register(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUIRouter_Static(t *testing.T) {
	r := gin.Default()

	NewUIRouter().Register(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/static/version.txt", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", w.Body.String())
}
