package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lyy42995004/IM-Go/internal/service"
	"github.com/lyy42995004/IM-Go/pkg/common/request"
	"github.com/lyy42995004/IM-Go/pkg/common/response"
	"github.com/lyy42995004/IM-Go/pkg/log"
)

// 获取消息列表
func GetMessage(c *gin.Context) {
	log.Info(c.Query("uuid"))

	var messageRequest request.MessageRequest
	err := c.BindQuery(&messageRequest)
	if err != nil {
		log.Error("bindQueryError", log.Any("bindQueryError", err))
	}
	log.Info("messageRequest params: ", log.Any("messageRequest", messageRequest))

	messages, err := service.MessageService.GetMessages(messageRequest)
	if err != nil {
		c.JSON(http.StatusOK, response.FailMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg(messages))
}