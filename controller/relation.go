package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/securemist/douyin-mini/model/constant"
	"github.com/securemist/douyin-mini/model/resp"
	"github.com/securemist/douyin-mini/service"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	resp.Response
	UserList *[]resp.User `json:"user_list"`
}

// 关注操作
func RelationAction(c *gin.Context) {
	id, _ := c.Get("userId")
	userId := id.(int64)

	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)

	if userId != toUserId {
		service.RelationAction(userId, toUserId)
	}

	c.JSON(http.StatusOK, constant.GENERAL_SUCCESS)
}

// 关注列表
func FollowList(c *gin.Context) {
	id, _ := c.Get("userId")
	currentUserId := id.(int64)

	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	userList := service.GetFollowList(currentUserId, userId)
	c.JSON(http.StatusOK, UserListResponse{
		constant.GENERAL_SUCCESS,
		&userList,
	})
}

// 粉丝列表
func FollowerList(c *gin.Context) {
	id, _ := c.Get("userId")
	currentUserId := id.(int64)

	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	userList := service.GetFollowerList(currentUserId, userId)
	c.JSON(http.StatusOK, UserListResponse{
		constant.GENERAL_SUCCESS,
		&userList,
	})
}

// 好友列表
func FriendList(c *gin.Context) {
	id, _ := c.Get("userId")
	currentUserId := id.(int64)

	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	userList := service.GetFriendList(currentUserId, userId)
	c.JSON(http.StatusOK, UserListResponse{
		constant.GENERAL_SUCCESS,
		&userList,
	})
}
