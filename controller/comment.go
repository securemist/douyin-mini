package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/securemist/douyin-mini/model/constant"
	"github.com/securemist/douyin-mini/model/resp"
	"github.com/securemist/douyin-mini/service"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	resp.Response
	CommentList *service.CommentList `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	resp.Response
	Comment *resp.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	id, _ := c.Get("userId")
	userId := id.(int64)

	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	service.FavoriteAction(userId, videoId)

	actionType := c.Query("action_type")
	// 发表评论
	if actionType == "1" {
		commentText := c.Query("comment_text")
		commentId, createTime := service.AddComment(userId, videoId, commentText)
		// 这里接口文档有点不懂，这里只有一个用户，为什么会存在 "is_follow": true, 这个字段
		user := service.GetUserInfo(0, userId)
		comment := resp.Comment{
			commentId,
			user,
			commentText,
			createTime.Format("01-02"),
		}

		c.JSON(http.StatusOK, CommentActionResponse{
			constant.GENERAL_SUCCESS,
			&comment,
		})
		return
	}

	// 删除评论
	if actionType == "2" {
		commentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
		service.DeleteComment(commentId)
		c.JSON(http.StatusOK, CommentActionResponse{
			constant.GENERAL_SUCCESS,
			nil,
		})
		return
	}

}

// CommentList all videos have same demo comment list
func HandleCommentList(c *gin.Context) {
	id, _ := c.Get("userId")
	userId := id.(int64)

	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)

	commentList := service.GetCommentList(userId, videoId)

	c.JSON(http.StatusOK, CommentListResponse{
		constant.GENERAL_SUCCESS,
		&commentList,
	})
}
