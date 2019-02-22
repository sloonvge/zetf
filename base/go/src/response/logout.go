package response

import "github.com/gin-gonic/gin"

func Logout(c *gin.Context) {

	c.String(200, "Logout!")
}
