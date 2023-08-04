package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/securemist/douyin-mini/model/constant"
	"github.com/securemist/douyin-mini/model/resp"
	"github.com/securemist/douyin-mini/service"
	"github.com/securemist/douyin-mini/util"
	"log"
	"net/http"
	"strconv"
)

type VideoListResponse struct {
	resp.Response
	VideoList *service.VideoList `json:"video_list"`
}

type VideoPublishResponse struct {
	resp.Response
}

// 视频发布接口，这个接口前端app有问题，这里没有问题
func Publish(c *gin.Context) {
	// 这里使用postman测试会出问题，接收不到文件，所以我采用html表单进行测试
	id, _ := c.Get("userId")
	userId := id.(int64)

	title := c.PostForm("title")
	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusOK, VideoPublishResponse{
			Response: constant.FILE_UPLOAD_FAILED,
		})
	}

	src, _ := file.Open()

	// 处理接收到的文件，这里放在到本地
	// 文件名为雪花id
	vid := util.GenerateId()
	playUrl := "./static/video/" + strconv.FormatInt(vid, 10) + ".mp4"
	// 由于前端app设计的缺陷，封面url没有必要取做，这里新发布的视频就同一使用默认的封面
	coverUrl := "./static/bg/default.jpg"

	err = util.Upload(playUrl, src)
	if err != nil {
		c.JSON(http.StatusOK, VideoPublishResponse{
			Response: constant.FILE_UPLOAD_FAILED,
		})
	}

	// 将记录添加进数据库
	_, err = service.AddWork(userId, playUrl, coverUrl, title)

	if err != nil {
		c.JSON(http.StatusOK, VideoPublishResponse{
			Response: constant.FILE_UPLOAD_FAILED,
		})
	}

	c.JSON(http.StatusOK, VideoPublishResponse{
		Response: constant.FILE_UPLOAD_SUCCESS,
	})
}

// 列出用户的所有视频列表
func PublishList(c *gin.Context) {
	id, _ := c.Get("userId")
	currentUserId := id.(int64)

	// 要查看的用户id
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		// 参数异常
		log.Println("params error string => int64 : ", c.Query("user_id"))
		return
	}

	// 作者信息
	// 获取指定用户视频列表
	videoList := service.GetVideoList(currentUserId, userId)
	if len(videoList) == 0 {
		videoList = nil
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response:  constant.GENERAL_SUCCESS,
		VideoList: &videoList,
	})
}
