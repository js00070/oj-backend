package api

import (
	"oj/service"

	"github.com/gin-gonic/gin"
)

// CommitCode 提交代码
func CommitCode(c *gin.Context) {
	var service service.CommitCodeService
	if err := c.BindJSON(&service); err == nil {
		res := service.Commit(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
