package response

import "github.com/gin-gonic/gin"

func Register(c *gin.Context) {
	username := c.Query("username")
	passwd := c.Query("passwd")

	c.String(200, "Welcome, %s %s!", username, passwd)
}
