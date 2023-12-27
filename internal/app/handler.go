// internal/app/handler.go
package app

import (
	"github.com/gin-gonic/gin"
)

// ReorganizeStringHandler handles the reorganize string functionality
func (a *App) ReorganizeStringHandler(c *gin.Context) {
	s := c.PostForm("s")
	result := a.ReorganizeString(s)
	c.JSON(200, gin.H{"result": result})
}
