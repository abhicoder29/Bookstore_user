package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
// c * interfance 
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
