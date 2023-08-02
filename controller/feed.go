package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/securemist/douyin-mini/model/constant"
	"github.com/securemist/douyin-mini/model/resp"
	"github.com/securemist/douyin-mini/service"
	"github.com/securemist/douyin-mini/util"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	resp.Response
	VideoList *service.VideoList `json:"video_list,omitempty"`
	NextTime  *int64             `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	// 得到userId
	id, _ := c.Get("userId")
	currentUserId := id.(int64)

	// 解析时间，没有时间就是当前时间
	time0 := c.Query("last_time")
	var lastTime time.Time
	if time0 != "" {
		time1, _ := strconv.ParseInt(time0, 10, 64)
		lastTime = util.UnixToTime(time1)
	} else {
		lastTime = time.Now()
	}

	// 返回按投稿时间倒序的视频列表 十条视频
	videoList := service.GetFeedVideoList(currentUserId, lastTime)
	//本次feed最后一个视频的时间 = 下一次feed的起始时间
	nextWorkId := videoList[len(videoList)-1].Id
	// 查询出时间
	nextTime := util.TimeToUnix(service.GetTimeByWorkId(nextWorkId))

	c.JSON(http.StatusOK, FeedResponse{
		Response:  constant.GENERAL_SUCCESS,
		VideoList: &videoList,
		NextTime:  &nextTime,
	})
}
