package main

import (
	"bufio"
	"fmt"
	"github.com/securemist/douyin-mini/model/constant"
	"github.com/securemist/douyin-mini/service"
	"github.com/securemist/douyin-mini/util"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	download()
}

// 读取./video.txt下的所有url，下载到本地，同时数据库添加记录
func download() {
	open, err := os.Open("./video.txt")
	defer open.Close()

	if err != nil {
		fmt.Println("err1", err)
	}

	content := bufio.NewReader(open)
	index := 1
	for {
		line, _, err := content.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return
		}

		url := string(line)
		coverUrl := url + constant.COVER_SUFFIX

		// 文件名使用雪花算法生成的id
		id := util.GenerateId()
		videoPath := "./static/video/"
		fileName := videoPath + strconv.FormatInt(id, 10) + ".mp4"
		coverPath := "./static/cover/"
		coverName := coverPath + strconv.FormatInt(id, 10) + ".jpg"

		// 下载视频
		resp, err := http.Get(url)
		handleError(err)

		file, err := os.Create(fileName)
		handleError(err)

		_, err = io.Copy(file, resp.Body)
		handleError(err)

		title := "title" + strconv.FormatInt(int64(index), 10)

		// 下载封面图片
		resp, err = http.Get(coverUrl)
		handleError(err)

		file, err = os.Create(coverName)
		handleError(err)

		_, err = io.Copy(file, resp.Body)

		// 本地资源路径
		url = fileName[1:]
		coverUrl = coverName[2:]

		_ = service.AddWork(int64(index/10+1), url, coverUrl, title)

		fmt.Println(strconv.FormatInt(int64(index), 10), "的任务完成了")
		index++
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal("error in generate.go")
	}
}
