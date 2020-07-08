package v1

import "github.com/gin-gonic/gin"

// Test 测试api
func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 1,
	})
}
