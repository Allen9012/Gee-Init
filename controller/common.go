package controller

import (
	"errors"
	"gee-Init/util"

	"github.com/gin-gonic/gin"
)

var ErrorIdNotExist = errors.New("用户不可用")

func getUserId(c *gin.Context) (int64, error) {
	value, exist := c.Get(util.KeyUserId)
	if !exist {
		return -1, ErrorIdNotExist
	}
	userId, ok := value.(int64)
	if !ok {
		return -1, ErrorIdNotExist
	}
	return userId, nil
}
