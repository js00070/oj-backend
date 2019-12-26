package service

import (
	"oj/model"
	"oj/serializer"

	"github.com/gin-gonic/gin"
)

// CommitCodeService 提交代码服务
type CommitCodeService struct {
	Code      string `json:"code"`
	Lang      string `json:"lang"`
	ProblemID uint   `json:"pid"`
}

// Commit 提交代码
func (service *CommitCodeService) Commit(c *gin.Context) serializer.Response {
	id, ok := c.Get("UserID")
	if !ok {
		return serializer.Err(312, "提交代码失败", nil)
	}
	_, err := model.CreateCommit(id.(uint), service.ProblemID, service.Code, service.Lang)
	if err != nil {
		return serializer.DBErr("", err)
	}
	return serializer.Success()
}
