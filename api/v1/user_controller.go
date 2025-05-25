package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lyy42995004/IM-Go/internal/model"
	"github.com/lyy42995004/IM-Go/internal/service"
	"github.com/lyy42995004/IM-Go/pkg/common/request"
	"github.com/lyy42995004/IM-Go/pkg/common/response"
	"github.com/lyy42995004/IM-Go/pkg/log"
)

// 用户注册
func Register(c *gin.Context) {
	var user model.User
	c.ShouldBindJSON(&user)
	log.Debug("user", log.Any("user", user))

	if err := service.UserService.Register(&user); err != nil {
		c.JSON(http.StatusOK, response.FailMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg(nil))
}

// 用户登录
func Login(c *gin.Context) {
	var user model.User
	c.ShouldBindJSON(&user)	
	log.Debug("user", log.Any("user", user))

	if service.UserService.Login(&user) {
		c.JSON(http.StatusOK, response.SuccessMsg(user))
		return
	}
	c.JSON(http.StatusOK, response.FailMsg("login failed"))
}

// 修改用户信息
func ModifyUserInfo(c *gin.Context) {
	var user model.User
	c.ShouldBindJSON(&user)
	log.Debug("user", log.Any("user", user))

	if err := service.UserService.ModifyUserInfo(&user); err != nil {
		c.JSON(http.StatusOK, response.FailMsg(err.Error()))
	}
	c.JSON(http.StatusOK, response.SuccessMsg(nil))
}

// 从数据库中获取用户的详细信息
func GetUserDetails(c *gin.Context) {
	uuid := c.Param("uuid")

	c.JSON(http.StatusOK, response.SuccessMsg(service.UserService.GetUserDetails(uuid)))
}

// 根据名称查找用户或群组
func GetUserOrGroupByName(c *gin.Context) {
	name := c.Query("name")

	c.JSON(http.StatusOK, response.SuccessMsg(service.UserService.GetUserOrGroupByName(name)))
}

// 获取所有用户的列表信息
func GetUserList(c *gin.Context) {
	uuid := c.Param("uuid")

	c.JSON(http.StatusOK, response.SuccessMsg(service.UserService.GetUserList(uuid)))
}

// 添加好友
func AddFriend(c *gin.Context) {
	var userFriendRequest request.FriendRequest
	if err := service.UserService.AddFriend(&userFriendRequest); err != nil {
		c.JSON(http.StatusOK, response.FailMsg(err.Error()))
	}
	c.JSON(http.StatusOK, response.SuccessMsg(nil))
}
