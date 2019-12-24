package api

import (
	"oj/model"
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

// GetCommitList 获取提交列表
func GetCommitList(c *gin.Context) {
	id, _ := c.Get("UserID")
	commitList, err := model.GetCommitList(id.(uint))
	if err != nil {
		c.JSON(500, ErrorResponse(err))
	} else {
		c.JSON(200, gin.H{
			"commitlist": commitList,
		})
	}
}

// GetProblemList 获取题目列表
func GetProblemList(c *gin.Context) {
	var problemList []model.Problem
	err := model.DB.Find(&problemList).Error
	if err != nil {
		c.JSON(500, ErrorResponse(err))
	} else {
		c.JSON(200, gin.H{
			"problemlist": problemList,
		})
	}
}
