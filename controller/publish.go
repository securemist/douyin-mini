package controller

import (
	"fmt"
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

// 视频发布接口
func Publish(c *gin.Context) {
	// 这里使用postman测试会出问题，接收不到文件，所以我采用html表单进行测试
	id, _ := c.Get("userId")
	userId := id.(int64)

	title := c.PostForm("title")
	file, err := c.FormFile("file")

	if err != nil {
		fmt.Println("file is nil ", err)
		return
	}

	src, _ := file.Open()

	// 处理接收到的文件，这里使用 oss 对象存储
	// 上传之后的文件URL
	playUrl := util.Upload(constant.Video_Path, file.Filename, src)

	if playUrl == "" {
		c.JSON(http.StatusOK, VideoPublishResponse{
			Response: constant.FILE_UPLOAD_FAILED,
		})
	}

	// 阿里云视频截帧，见 https://help.aliyun.com/zh/oss/user-guide/video-snapshots
	coverUrl := playUrl + "x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600"

	// 将记录添加进数据库
	_ = service.AddWork(userId, playUrl, coverUrl, title)
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
