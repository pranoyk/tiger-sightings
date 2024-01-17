// middleware_test.go
package middleware

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pranoyk/tiger-sightings/helper"
	"github.com/stretchr/testify/assert"
)

func TestJwtAuthMiddleware(t *testing.T) {
	r := gin.Default()
	r.Use(JwtAuthMiddleware())

	r.GET("/test", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	w := helper.PerformRequest(r, "GET", "/test", nil)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
