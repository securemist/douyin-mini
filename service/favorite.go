package service

import (
	"github.com/securemist/douyin-mini/model/db"
	"github.com/securemist/douyin-mini/util"
	"sort"
)

func FavoriteAction(userId, videoId int64) {
	Db := util.GetDbConnection()
	r := 0
	_ = Db.QueryRow("SELECT COUNT(id) FROM user_favorite WHERE user_id = ? AND work_id = ?",
		userId, videoId).Scan(&r)

	// 数据库没有记录，是第一次点赞
	var err error
	if r == 0 {
		_, err = Db.Exec("INSERT INTO user_favorite VALUES (?, ?, ?, ?, ?)",
			int64(0), videoId, userId, util.TimeNow(), 0)
	} else { // 后续的点赞或者取消点赞
		_, err = Db.Exec("UPDATE user_favorite SET deleted = CASE WHEN deleted = 0 THEN 1 WHEN deleted = 1 THEN 0 END WHERE user_id = ? AND work_id = ?",
			userId, videoId)
	}

	HandleSqlError(err)
}

func GetFavoriteList(currentUserId, userId int64) VideoList {
	Db := util.GetDbConnection()
	defer Db.Close()

	// 查出目标用户所有点赞的视频
	var workList []db.Work

	err := Db.Select(&workList, "SELECT id, user_id, play_url, cover_url,title, create_time FROM user_work WHERE id in (SELECT work_id FROM user_favorite WHERE user_id = ? AND deleted = 0) AND deleted = 0", userId)
	HandleSqlError(err)

	videoList := HandleWorkList(currentUserId, workList)
	sort.Sort(videoList)

	return videoList
}
