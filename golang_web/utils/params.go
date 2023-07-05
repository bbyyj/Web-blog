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

//函数的功能是获取指定查询参数的值，并将其转换为整数类型。
//如果查询参数不存在或无法转换为整数，则返回一个默认值。
func DefaultQueryInt(ctx *gin.Context, key, def string) int {
	param := ctx.DefaultQuery(key, def)
	i, _ := strconv.Atoi(param)
	return i
}
