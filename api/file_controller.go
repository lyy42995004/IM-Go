package api

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lyy42995004/IM-Go/internal/config"
	"github.com/lyy42995004/IM-Go/internal/service"
	"github.com/lyy42995004/IM-Go/pkg/common/response"
	"github.com/lyy42995004/IM-Go/pkg/log"
)

// 前端通过文件名称获取文件流，显示文件
func GetFile(c *gin.Context) {
	fileName := c.Param("fileName")
	log.Info(fileName)
	data, _ :=  os.ReadFile(config.GetConfig().StaticPath.FilePath + fileName)
	c.Writer.Write(data)
}

// 上传头像等文件
func SaveFile(c *gin.Context) {
	namePreffix := uuid.New().String()

	userUuid := c.PostForm("uuid")

	file, _ := c.FormFile("file")
	fileName := file.Filename
	index := strings.LastIndex(fileName, ".")
	suffix := fileName[index:]

	newFileName := namePreffix + suffix

	log.Info("file", log.Any("file name", config.GetConfig().StaticPath.FilePath+newFileName))
	log.Info("userUuid", log.Any("userUuid name", userUuid))

	c.SaveUploadedFile(file, config.GetConfig().StaticPath.FilePath+newFileName)
	err := service.UserService.ModifyUserAvatar(newFileName, userUuid)
	if err != nil {
		c.JSON(http.StatusOK, response.FailMsg(err.Error()))
	}
	c.JSON(http.StatusOK, response.SuccessMsg(newFileName))
}

