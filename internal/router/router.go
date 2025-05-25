package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lyy42995004/IM-Go/api/v1"
	"github.com/lyy42995004/IM-Go/pkg/common/response"
	"github.com/lyy42995004/IM-Go/pkg/log"
)

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	server := gin.Default()
	server.Use(Cors())
	server.Use(Recovery)

	socket := RunSocket

	group := server.Group("")
	{
		// 用户管理功能
		group.GET("/user", v1.GetUserList)
		group.GET("/user/:uuid", v1.GetUserDetails)
		group.GET("/user/name", v1.GetUserOrGroupByName)
		group.POST("/user/register", v1.Register)
		group.POST("/user/login", v1.Login)
		group.PUT("/user", v1.ModifyUserInfo)

		group.POST("/friend", v1.AddFriend)

		group.GET("/message", v1.GetMessage)

		group.GET("/file/:fileName", v1.GetFile)
		group.POST("/file", v1.SaveFile)

		group.GET("/group/:uuid", v1.GetGroup)
		group.POST("/group/:uuid", v1.SaveGroup)
		group.POST("/group/join/:userUuid/:groupUuid", v1.JoinGroup)
		group.GET("/group/user/:uuid", v1.GetGroupUsers)

		group.GET("/socket.io", socket)
	}
	return server
}

// 处理跨域请求
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") // 请求头部
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Error("HttpError", log.Any("HttpError", err))
			}
		}()

		c.Next()
	}
}

// 捕获并处理路由处理过程中出现的异常
func Recovery(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Error("gin catch error", log.Any("error", r))
			c.JSON(http.StatusOK, response.FailMsg("系统内部错误"))
		}
	}()
	c.Next()
}
