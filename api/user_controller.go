package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lyy42995004/IM-Go/internal/model"
	"github.com/lyy42995004/IM-Go/internal/service"
	"github.com/lyy42995004/IM-Go/pkg/common/response"
	"github.com/lyy42995004/IM-Go/pkg/log"
)

// 处理用户登录的逻辑
func Login(c *gin.Context) {
	var user model.User
	c.ShouldBindJSON(&user)	
	log.Debug("user", log.Any("user", user))

	// if service.UserService.Login(&user) {
	// 	c.JSON(http.StatusOK, response.SuccessMsg(user))
	// 	return
	// }

	c.JSON(http.StatusOK, response.FailMsg("login failed"))
}

func Register(c *gin.Context) {
}

func ModifyUserInfo(c *gin.Context) {
}

func GetUserDetails(c *gin.Context) {
}

func GetUserOrGroupByName(c *gin.Context) {
}

// 获取所有用户的列表信息
func GetUserList(c *gin.Context) {
}

func AddFriend(c *gin.Context) {
}
