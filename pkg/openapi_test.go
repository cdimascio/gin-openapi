package openapi

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOpenapi_OK(t *testing.T) {
	t.Run("returns 200", func(t *testing.T) {
		called := false
		router := gin.Default()
		router.Use(ValidateRequests("openapi_test_spec.yml"))
		router.GET("/v1/ping", func(c *gin.Context) {
			called = true
			c.JSON(200, gin.H{
				"id": "pong",
			})
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/v1/ping", nil)
		req.Header.Set("accept", "application/json")
		router.ServeHTTP(w, req)

		assert.True(t, called)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestOpenapi_MissingRequiredField(t *testing.T) {
	t.Run("returns 422", func(t *testing.T) {
		called := false
		router := gin.Default()
		router.Use(ValidateRequests("openapi_test_spec.yml"))
		router.POST("/v1/ping", func(c *gin.Context) {
			called = true
			c.JSON(200, gin.H{
				"id": "pong",
			})
		})

		b := new(bytes.Buffer)
		_ =json.NewEncoder(b).Encode(struct { Id string } { Id: "test" })
		req, _ := http.NewRequest("POST", "/v1/ping", b)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.False(t, called)
		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})
}

func TestOpenapi_UnsupportedMediaType(t *testing.T) {
	t.Run("returns 400", func(t *testing.T) {
		called := false
		router := gin.Default()
		router.Use(ValidateRequests("openapi_test_spec.yml"))
		router.POST("/v1/ping", func(c *gin.Context) {
			called = true
			c.JSON(200, gin.H{
				"id": "pong",
			})
		})

		b := new(bytes.Buffer)
		_ =json.NewEncoder(b).Encode(struct { Id string } { Id: "test" })
		req, _ := http.NewRequest("POST", "/v1/ping", b)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.False(t, called)
		assert.Equal(t, http.StatusUnsupportedMediaType, w.Code)
	})
}

func TestOpenapiNotFound(t *testing.T) {
	t.Run("returns 404", func(t *testing.T) {
		router := gin.Default()
		router.Use(ValidateRequests("openapi_test_spec.yml"))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/not_found", nil)
		req.Header.Set("accept", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}