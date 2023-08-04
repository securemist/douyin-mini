package main

import (
	"bufio"
	"fmt"
	"github.com/securemist/douyin-mini/model/constant"
	"github.com/securemist/douyin-mini/util"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

func main() {
	generate()
}

// 读取./video.txt下的所有url，下载到本地，同时数据库添加记录
func generate() {
	open, err := os.Open("./video.txt")
	handleError(err)

	defer open.Close()

	// 统计行数 暂时不使用
	scanner := bufio.NewScanner(open)
	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lineCount++
		}
	}

	// 读取文件内容
	_, _ = open.Seek(0, 0)
	content := bufio.NewReader(open)

	fmt.Println("首次运行会生成测试数据，请耐心等待......")

	for {
		line, _, err := content.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return
		}

		// 这是一个视频的url
		url := string(line)
		videoPath := "./static/video/"
		coverPath := "./static/cover/"

		// 文件名策略 同一个作品的视频和方面采用同一个雪花id
		id := util.GenerateId()

		videoFileName := videoPath + strconv.FormatInt(id, 10) + ".mp4"
		coverFileName := coverPath + strconv.FormatInt(id, 10) + ".jpg"

		// 下载视频
		resp, err := http.Get(url)
		handleError(err)
		videoFile, err := os.Create(videoFileName)
		defer videoFile.Close()
		handleError(err)
		_, err = io.Copy(videoFile, resp.Body)
		handleError(err)

		// 下载封面  视频截帧
		url += constant.COVER_SUFFIX
		resp, err = http.Get(url)
		handleError(err)
		coverFile, err := os.Create(coverFileName)
		defer coverFile.Close()
		handleError(err)
		_, err = io.Copy(coverFile, resp.Body)
		handleError(err)

		// 添加记录数据库
		playUrl := videoFileName[1:]
		coverUrl := coverFileName[1:]
		title := "douyin-mini by securemist"

		Db := util.GetDbConnection()
		defer Db.Close()

		r, err := Db.Exec("INSERT INTO user_work (user_id, play_url, cover_url,title,  create_time) VALUES ( ?, ?, ?, ?, ?)", int64(rand.Intn(5)+1), playUrl, coverUrl, title, util.TimeNow())
		handleError(err)
		_, err = r.LastInsertId()
		handleError(err)

		fmt.Print("==")
	}

	fmt.Print("\n")
	fmt.Println("data generate success")

}

func handleError(err error) {
	if err != nil {
		log.Fatal("generate data error : ", err)
	}
}
