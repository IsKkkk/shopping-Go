package api_helper

import (
	"shopping/utils/pagination"

	"github.com/gin-gonic/gin"
)

var userIdText = "userId"

// 从context获得用户id
func GetUserId(g *gin.Context) uint {
	return uint(pagination.ParseInt(g.GetString(userIdText), -1))
}
