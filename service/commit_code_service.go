package service

import (
	"oj/cache"
	"oj/model"
	"oj/serializer"

	"github.com/gin-gonic/gin"
)

// CommitCodeService 提交代码服务
type CommitCodeService struct {
	Code string `json:"code"`
}

// Commit 提交代码
func (service *CommitCodeService) Commit(c *gin.Context) serializer.Response {
	id, ok := c.Get("UserID")
	if !ok {
		return serializer.Err(312, "提交代码失败", nil)
	}
	commit, err := model.CreateCommit(id.(uint), service.Code)
	if err != nil {
		return serializer.DBErr("", err)
	}
	cache.RedisClient.Ping()
	cache.RedisClient.Set("lock", 0, 0)
	cache.RedisClient.LPush("code", service.Code)
	cache.RedisClient.LPush("commit", commit.ID)
	cache.RedisClient.Set("lock", 1, 0)
	return serializer.Success()
}
