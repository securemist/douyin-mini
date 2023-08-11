package service

import (
	"github.com/jmoiron/sqlx"
	conf "github.com/securemist/douyin-mini/config"
	"github.com/securemist/douyin-mini/model/db"
	"github.com/securemist/douyin-mini/model/resp"
	"github.com/securemist/douyin-mini/util"
	"log"
	"sort"
)

type Service struct {
}

type VideoList []resp.Video

func (list VideoList) Len() int {
	return len(list)
}

func (list VideoList) Swap(i, j int) {
	temp := list[i]
	list[i] = list[j]
	list[j] = temp
}

func (list VideoList) Less(i, j int) bool {
	t1 := util.TimeStringToUnix(list[i].CreateTime)
	t2 := util.TimeStringToUnix(list[j].CreateTime)
	return t1 > t2
	return t2 <= t1
}

func GetVideoList(currentUseId, userId int64) VideoList {
	Db := util.GetDbConnection()
	defer Db.Close()

	// 查询出的指定用户的所有视频
	var workList []db.Work
	err := Db.Select(&workList, "SELECT id, user_id, play_url, cover_url, title, create_time FROM user_work WHERE user_id = ? AND deleted = 0 ORDER BY create_time DESC", userId)
	HandleSqlError(err)

	itemCount := len(workList) // 用户作品的数量

	videoChan := make(chan resp.Video, itemCount)
	author := GetUserInfo(currentUseId, userId)

	// 并发遍历作品列表，每个协程生成一个video放入管道，管道满时继续运行
	// 这里的并发安全行亟待验证 TODO
	for index, work := range workList {
		go add(index, work, currentUseId, Db, videoChan, author)
	}

	for {
		if len(videoChan) == itemCount {
			close(videoChan)
			break
		}
	}

	// 将管道内容转存进切片
	var videoList VideoList
	for video := range videoChan {
		videoList = append(videoList, video)
	}

	// 按时间倒序排序，重写Less方法
	sort.Sort(videoList)
	return videoList
}

// @Description:  用户发布作品
// @param userId 用户id
// @param playUrl 视频url
// @param coverUrl 封面url
// @return int64 添加记录之后生成的作品id
func AddWork(userId int64, playUrl, coverUrl, title string) (int64, error) {
	Db := util.GetDbConnection()
	defer Db.Close()

	r, err := Db.Exec("INSERT INTO user_work (user_id, play_url, cover_url,title,  create_time) VALUES ( ?, ?, ?, ?, ?)", userId, playUrl, coverUrl, title, util.TimeNow())
	if err != nil {
		return 0, err
	}

	workId, err := r.LastInsertId()
	if err != nil {
		return 0, err
	}
	return workId, nil
}

func add(index int, work db.Work, currentUserId int64, Db *sqlx.DB, videoChan chan resp.Video, author resp.User) {

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

	video := resp.Video{
		Id:            work.Id,
		Author:        author,
		CoverUrl:      conf.Project_url_suffix + work.CoverUrl,
		PlayUrl:       conf.Project_url_suffix + work.PlayUrl,
		Title:         work.Title,
		FavoriteCount: favoriteCount,
		IsFavorite:    isFavorite,
		CommentCount:  commentCount,
		CreateTime:    work.CreateTime,
	}
	videoChan <- video
}

func HandleSqlError(err error) {
	if err != nil {
		log.Println("sql exec failed : ", err)
	}
}
