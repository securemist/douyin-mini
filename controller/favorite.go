package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/securemist/douyin-mini/model/constant"
	"github.com/securemist/douyin-mini/model/resp"
	"github.com/securemist/douyin-mini/service"
	"net/http"
	"strconv"
)

type FavoriteListResponse struct {
	resp.Response
	VideoList *service.VideoList
}

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	id, _ := c.Get("userId")
	userId := id.(int64)

	_ = c.Query("action_type") // 1点赞 2取消点赞

	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	service.FavoriteAction(userId, videoId)

	c.JSON(http.StatusOK, constant.GENERAL_SUCCESS)
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	id, _ := c.Get("userId")
	currentUserId := id.(int64)

	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	videoList := service.GetFavoriteList(currentUserId, userId)

	c.JSON(http.StatusOK, FavoriteListResponse{
		constant.GENERAL_SUCCESS,
		&videoList,
	})
}
