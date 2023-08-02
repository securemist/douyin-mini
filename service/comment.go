package service

import (
	"github.com/securemist/douyin-mini/model/db"
	"github.com/securemist/douyin-mini/model/resp"
	"github.com/securemist/douyin-mini/util"
	"sort"
	"time"
)

type CommentList []resp.Comment

func (list CommentList) Len() int {
	return len(list)
}

func (list CommentList) Swap(i, j int) {
	temp := list[i]
	list[i] = list[j]
	list[j] = temp
}

func (list CommentList) Less(i, j int) bool {
	t1 := util.TimeStringToUnix(list[i].CreateDate)
	t2 := util.TimeStringToUnix(list[j].CreateDate)
	return t1 > t2
}

// @Description:  添加评论
// @param userId 用户id
// @param videoId 作品id
// @param commentText 评论内容
// @return int64 评论id
// @return time.Time 评论时间
func AddComment(userId, videoId int64, commentText string) (int64, time.Time) {
	Db := util.GetDbConnection()
	defer Db.Close()

	now := util.TimeNow()
	r, err := Db.Exec("INSERT INTO work_comment VALUES (?, ?, ?, ?, ?, ?)",
		0, videoId, userId, commentText, now, 0)
	HandleSqlError(err)

	commentId, err := r.LastInsertId()
	HandleSqlError(err)

	return commentId, now
}

// @Description: 删除单条评论
// @param commentId 评论Id
func DeleteComment(commentId int64) {
	Db := util.GetDbConnection()
	defer Db.Close()

	r, err := Db.Exec("UPDATE work_comment SET deleted = 1 WHERE id = ?", commentId)
	HandleSqlError(err)

	_, err = r.RowsAffected()
	HandleSqlError(err)
}

// @Description:  获取评论列表
// @param userId 当前用户id
// @param videoId 作品id
// @return CommentList
func GetCommentList(currentUserId, videoId int64) CommentList {
	Db := util.GetDbConnection()
	defer Db.Close()

	var list []db.WorkComment
	err := Db.Select(&list, "SELECT id, user_id, work_id, content, create_time FROM work_comment WHERE work_id = ? AND deleted = 0", videoId)
	HandleSqlError(err)

	var commentList CommentList
	for _, workComment := range list {
		user := GetUserInfo(currentUserId, workComment.UserId)
		comment := resp.Comment{
			Id:         workComment.Id,
			User:       user,
			Content:    workComment.Content,
			CreateDate: workComment.CreateTime,
		}
		commentList = append(commentList, comment)
	}

	// 按日期排序
	sort.Sort(commentList)

	// 转换日期格式
	for index, comment := range commentList {
		commentList[index].CreateDate = util.UnixToTime(util.TimeStringToUnix(comment.CreateDate)).Format("01-02")
	}
	return commentList

}
