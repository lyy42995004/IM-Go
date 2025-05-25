package service

import (
	"github.com/lyy42995004/IM-Go/internal/dao/pool"
	"github.com/lyy42995004/IM-Go/internal/model"
	"github.com/lyy42995004/IM-Go/pkg/common/constant"
	"github.com/lyy42995004/IM-Go/pkg/common/request"
	"github.com/lyy42995004/IM-Go/pkg/common/response"
	"github.com/lyy42995004/IM-Go/pkg/errors"
	"github.com/lyy42995004/IM-Go/pkg/log"
	"github.com/lyy42995004/IM-Go/pkg/protocol"
	"gorm.io/gorm"
)

type messageService struct{}
var MessageService = new (messageService)

// 获取消息列表
func (m *messageService) GetMessages(message request.MessageRequest) ([]response.MessageResponse, error) {
	db := pool.GetDB()

	mirgate := &model.Message{}
	pool.GetDB().AutoMigrate(&mirgate)

	// 处理用户间消息请求
	if message.MessageType == constant.MESSAGE_TYPE_USER {
		var queryUser *model.User
		db.First(&queryUser)

		if queryUser.Id == 0 {
			return nil, errors.New("用户不存在")
		}

		var friend *model.User
		db.First(&friend, "username = ?", message.FriendUsername)
		if friend.Id == 0 {
			return nil, errors.New("用户不存在")
		}

		var messages []response.MessageResponse

		db.Raw(`SELECT m.id, m.from_user_id, m.to_user_id, m.content, m.content_type, m.url,
				m.created_at, u.username AS from_username, u.avatar, to_user.username AS to_username 
				FROM messages AS m 
				LEFT JOIN users AS u ON m.from_user_id = u.id
				LEFT JOIN users AS to_user ON m.to_user_id = to_user.id
				WHERE from_user_id IN (?, ?) AND to_user_id IN (?, ?)`,
			queryUser.Id, friend.Id, queryUser.Id, friend.Id).Scan(&messages)

		return messages, nil
	}

	// 处理群组消息请求
	if message.MessageType == constant.MESSAGE_TYPE_GROUP {
		messages, err := fetchGroupMessage(db, message.Uuid)
		if err != nil {
			return nil, err
		}

		return messages, nil
	}

	return nil, errors.New("不支持查询类型")
}

// 根据群组 UUID 从数据库中获取该群组的消息列表
func fetchGroupMessage(db *gorm.DB, toUuid string) ([]response.MessageResponse, error) {
	var group model.Group
	db.First(&group, "uuid = ?", toUuid)
	if group.ID <= 0 {
		return nil, errors.New("群组不存在")
	}

	var messages []response.MessageResponse

	db.Raw(`SELECT m.id, m.from_user_id, m.to_user_id, m.content, m.content_type,
			m.url, m.created_at, u.username AS from_username, u.avatar 
			FROM messages AS m 
			LEFT JOIN users AS u ON m.from_user_id = u.id 
			WHERE m.message_type = 2 AND m.to_user_id = ?`,
		group.ID).Scan(&messages)

	return messages, nil
}

// 将传入的消息保存到数据库
func (m *messageService) SaveMessage(message protocol.Message) {
	db := pool.GetDB()

	var fromUser model.User
	db.Find(&fromUser, "uuid = ?", message.From)
	if fromUser.Id == 0 {
		log.Error("SaveMessage not find from user", log.Any("SaveMessage not find from user", fromUser.Id))
		return
	}

	var toUserId int32 = 0

	// 处理用户间消息的接收者
	if message.MessageType == constant.MESSAGE_TYPE_USER {
		var toUser model.User
		db.Find(&toUser, "uuid = ?", message.To)
		if toUser.Id == 0 {
			return
		}
		toUserId = toUser.Id
	}

	// 处理群组消息的接收者
	if message.MessageType == constant.MESSAGE_TYPE_GROUP {
		var group model.Group
		db.Find(&group, "uuid = ?", message.To)
		if group.ID == 0 {
			return
		}
		toUserId = group.ID
	}

	saveMessage := model.Message{
		FromUserId:  fromUser.Id,
		ToUserId:    toUserId,
		Content:     message.Content,
		ContentType: int16(message.ContentType),
		MessageType: int16(message.MessageType),
		Url:         message.Url,
	}
	db.Save(&saveMessage)
}
