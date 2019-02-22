package response

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	username := c.Query("username")
	passwd := c.Query("passwd")


	if username != "" && passwd != "" {
		c.String(200, "Hello, %s %s!", username, passwd)
	} else {
		c.String(200, "Not correct Password!")
	}
}
