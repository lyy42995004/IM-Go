package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/lyy42995004/IM-Go/internal/dao/pool"
	"github.com/lyy42995004/IM-Go/internal/model"
	"github.com/lyy42995004/IM-Go/pkg/common/request"
	"github.com/lyy42995004/IM-Go/pkg/common/response"
	"github.com/lyy42995004/IM-Go/pkg/errors"
	"github.com/lyy42995004/IM-Go/pkg/log"
)

type userService struct {
}

var UserService = new(userService)

// 用户注册
func (u *userService) Register(user *model.User) error {
	db := pool.GetDB()

	var userCount int64
	db.Model(user).Where("username", user.Username).Count(&userCount)
	// SELECT COUNT(*) FROM users WHERE username = '[user.Username]';

	if userCount > 0 {
		return errors.New("user already exists")
	}

	user.Uuid = uuid.New().String()
	user.CreateAt = time.Now()
	user.DeleteAt = 0

	db.Create(&user)
	return nil
}

// 用户登录
func (u *userService) Login(user *model.User) bool {
	pool.GetDB().AutoMigrate(&user)
	log.Debug("user", log.Any("user in service", user))

	db := pool.GetDB()

	var queryUser *model.User
	db.First(&queryUser, "username = ?", user.Username)
	log.Debug("queryUser", log.Any("queryUser", queryUser))

	user.Uuid = queryUser.Uuid
	return queryUser.Password == user.Password
}

// 修改用户信息
func (u *userService) ModifyUserInfo(user *model.User) error {
	db := pool.GetDB()

	var queryUser *model.User
	db.First(&queryUser, "username = ?", user.Username)
	log.Debug("queryUser", log.Any("queryUser", queryUser))

	var nullId int32 = 0
	if nullId == queryUser.Id {
		return errors.New("用户不存在")
	}
	queryUser.Nickname = user.Nickname
	queryUser.Email = user.Email
	queryUser.Password = user.Password

	db.Save(queryUser)
	return nil
}

// 根据用户的 UUID 从数据库中获取用户的详细信息
func (u *userService) GetUserDetails(uuid string) model.User {
	db := pool.GetDB()

	var queryUser *model.User
	db.Select("uuid", "username", "nickname", "avatar").First(&queryUser, "uuid = ?", uuid)
	return *queryUser
}

// 根据名称查找用户或群组
func (u *userService) GetUserOrGroupByName(name string) response.SerachResponse {
	db := pool.GetDB()

	var queryUser *model.User
	db.Select("uuid", "username", "nickname", "avatar").First(&queryUser, "username = ?", name)

	var queryGroup *model.Group
	db.Select("uuid", "name").First(&queryGroup, "name = ?", name)

	return response.SerachResponse{
		User:  *queryUser,
		Group: *queryGroup,
	}
}

// 根据用户的 uuid 获取该用户的好友列表
func (u *userService) GetUserList(uuid string) []model.User {
	db := pool.GetDB()
	var queryUser *model.User
	db.First(&queryUser, "uuid = ?", uuid)
	var nullId int32 = 0
	if nullId == queryUser.Id {
		return nil
	}

	var queryUsers []model.User
	db.Raw(`SELECT u.username, u.uuid, u.avatar
			FROM user_friends AS uf 
			JOIN users AS u ON uf.friend_id = u.id 
			WHERE uf.user_id = ?`, queryUser.Id).Scan(&queryUsers)
	return queryUsers
}

// 添加好友
func (u *userService) AddFriend(userFriendRequest *request.FriendRequest) error {
	db := pool.GetDB()

	// 查询发起添加好友请求的用户
	var queryUser *model.User
	db.First(&queryUser, "uuid = ?", userFriendRequest.Uuid)
	log.Debug("queryUser", log.Any("queryUser", queryUser))
	var nullId int32 = 0
	if nullId == queryUser.Id {
		return errors.New("用户不存在")
	}

	// 查询要添加的好友
	var friend *model.User
	db.First(&friend, "username = ?", userFriendRequest.FriendUsername)
	if nullId == friend.Id {
		return errors.New("无法查询到好友")
	}

	// 创建用户好友关系记录
	userFriend := model.UserFriend{
		UserId:   queryUser.Id,
		FriendId: friend.Id,
	}

	// 检查用户和好友是否已经是好友关系
	var userFriendQuery *model.UserFriend
	db.First(&userFriendQuery, "user_id = ? and friend_id = ?", queryUser.Id, friend.Id)
	if nullId != userFriendQuery.ID {
		return errors.New("该用户已经是你好友")
	}

	return nil
}

// 根据用户的 uuid 修改用户的头像
func (u *userService) ModifyUserAvatar(avatar string, userUuid string) {
	db := pool.GetDB()

	var queryUser *model.User
	db.First(&queryUser, "uuid = ?", userUuid)
	if nu

}
