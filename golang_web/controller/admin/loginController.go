package admin

import (
	"blog_web/db/service"
	"blog_web/model"
	"blog_web/response"
	"blog_web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LoginController struct {
	userService *service.UserService
}

func NewLoginRouter() *LoginController {
	return &LoginController{
		userService: service.NewUserService(),
	}
}

// 登录时携带token
func (l *LoginController) Login(ctx *gin.Context) *response.Response {
	var u model.User
	err := ctx.ShouldBind(&u)
	if response.CheckError(err, "Bind param error") {
		return response.NewResponseOkND(response.LoginFailed)
	}

	user, err := l.userService.CheckUser(u.Username, u.Password)
	if response.CheckError(err, "Username or Password incorrect, IP:%s", ctx.GetHeader("X-Forwarded-For")) {
		println("note here 6 error")
		return response.NewResponseOkND(response.LoginFailed)
	}

	token, err := utils.CreateToken(uint32(user.Id), user.Username, time.Hour*24)
	if response.CheckError(err, "Generate token error") {
		ctx.Status(http.StatusInternalServerError)
		return nil
	}
	return response.NewResponseOk(response.LoginSuccess, token, user.Id)
}

func LoginAuthenticationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 解析Authorization头部里的token
		token := ctx.GetHeader("Authorization")
		if token == "" {
			// 说明跳过了上面的Login函数，直接由url跳转
			utils.Logger().Warning("未获得授权, ip:%s", ctx.Request.RemoteAddr)
			ctx.JSON(http.StatusOK, &(response.NewResponseOkND(response.Unauthorized).R))
			ctx.Abort()
			return
		}

		if _, _, ok := utils.VerifyToken(token); !ok {
			ctx.JSON(http.StatusOK, &(response.NewResponseOkND(response.Unauthorized).R))
			// 停止执行当前请求的剩余中间件和处理程序函数
			ctx.Abort()
			return
		}
		// 继续执行链中的下一个中间件或处理程序函数
		ctx.Next()
	}
}
