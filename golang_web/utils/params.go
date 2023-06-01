package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func QueryInt(ctx *gin.Context, key string) int {
	param := ctx.Query(key)
	i, _ := strconv.Atoi(param)
	return i
}

func DefaultQueryInt(ctx *gin.Context, key, def string) int {
	param := ctx.DefaultQuery(key, def)
	i, _ := strconv.Atoi(param)
	return i
}
