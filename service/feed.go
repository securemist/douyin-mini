package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/securemist/douyin-mini/model/constant"
	"github.com/securemist/douyin-mini/model/db"
	"github.com/securemist/douyin-mini/model/resp"
	"github.com/securemist/douyin-mini/util"
	"sort"
	"time"
)

func GetFeedVideoList(currentUserId int64, lastTime time.Time) VideoList {
	Db := util.GetDbConnection()

	// 查出 1. 排除自己的 2. 在给定时间之前的 3.按时间倒序的(放在切片中排序) 十条视频
	var workList []db.Work

	err := Db.Select(&workList, "SELECT id, user_id, play_url, cover_url,title, create_time FROM user_work WHERE user_id != ? AND deleted = 0 AND create_time < ? ORDER BY create_time DESC LIMIT 0, ?", currentUserId, util.TimeFormat(lastTime), constant.DEFAULT_FEED_LIST_SIZE)
	HandleSqlError(err)
	Db.Close()

	if len(workList) == 0 {
		return nil
	}

	videoList := HandleWorkList(currentUserId, workList)

	// 按时间倒序排序
	sort.Sort(videoList)
	return videoList
}

func HandleWorkList(currentUserId int64, workList []db.Work) VideoList {
	Db := util.GetDbConnection()
	defer Db.Close()

	itemCount := len(workList) // feed视频数量
	videoChan := make(chan resp.Video, itemCount)

	// 多线程查询视频信息
	for _, work := range workList {
		go add0(Db, work, videoChan, currentUserId)
	}
	for {
		if len(videoChan) == itemCount {
			close(videoChan)
			break
		}
	}

	// 将管道内容转存进切片
	videoList := VideoList{}
	for video := range videoChan {
		videoList = append(videoList, video)
	}
	return videoList
}

// @Description:  视频列表根据视频倒序排序
// @param videoList
// @return []resp.Video
func VideoListSort(videoList []resp.Video) []resp.Video {
	sortedVideoList := make([]resp.Video, len(videoList))
	copy(sortedVideoList, videoList)
	// 作品按时间倒序
	sort.Slice(sortedVideoList, func(i, j int) bool {
		t1 := util.TimeStringToUnix(videoList[i].CreateTime)
		t2 := util.TimeStringToUnix(videoList[j].CreateTime)
		return t1 > t2
	})
	return sortedVideoList
}

func GetTimeByWorkId(nextWorkId int64) time.Time {
	Db := util.GetDbConnection()
	defer Db.Close()

	var nextTime string
	err := Db.QueryRow("SELECT create_time FROM user_work WHERE id = ?", nextWorkId).Scan(&nextTime)
	HandleSqlError(err)

	return util.UnixToTime(util.TimeStringToUnix(nextTime))
}

func add0(Db *sqlx.DB, work db.Work, videoChan chan resp.Video, currentUserId int64) {
	isFavorite := false
	var commentCount int64
	var favoriteCount int64

	// 作品点赞数量
	_ = Db.QueryRow("SELECT COUNT(id) FROM user_favorite WHERE work_id = ? AND deleted = 0", work.Id).Scan(&favoriteCount)
	// 作品评论数量
	_ = Db.QueryRow("SELECT COUNT(id) FROM work_comment WHERE work_id = ? AND deleted = 0", work.Id).Scan(&commentCount)
	// 当前用户是否点赞该作品
	count := 0
	_ = Db.QueryRow("SELECT COUNT(id) FROM user_favorite WHERE user_id = ? AND work_id = ? AND deleted = 0", currentUserId, work.Id).Scan(&count)
	isFavorite = count != 0
	// 作者信息
	author := GetUserInfo(currentUserId, work.UserId)

	video := resp.Video{
		Id:            work.Id,
		Author:        author,
		CoverUrl:      work.CoverUrl,
		PlayUrl:       work.PlayUrl,
		Title:         work.Title,
		FavoriteCount: favoriteCount,
		IsFavorite:    isFavorite,
		CommentCount:  commentCount,
		CreateTime:    work.CreateTime,
	}
	videoChan <- video

}
