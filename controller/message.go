package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/securemist/douyin-mini/model/constant"
	"github.com/securemist/douyin-mini/model/resp"
	"github.com/securemist/douyin-mini/service"
	"net/http"
	"strconv"
)

type ChatResponse struct {
	resp.Response
	MessageList *[]resp.Message `json:"message_list"`
}

// 发送消息
func MessageAction(c *gin.Context) {
	id, _ := c.Get("userId")
	userId := id.(int64)

	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	content := c.Query("content")

	service.HandleMessageAction(userId, toUserId, content)

	c.JSON(http.StatusOK, constant.GENERAL_SUCCESS)
}

// 消息列表
func MessageChat(c *gin.Context) {
	id, _ := c.Get("userId")
	userId := id.(int64)

	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)

	messageList := service.GetMessageChat(userId, toUserId)

	c.JSON(http.StatusOK, ChatResponse{
		constant.GENERAL_SUCCESS,
		&messageList,
	})
}
