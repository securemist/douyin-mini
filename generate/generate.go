package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/securemist/douyin-mini/model/constant"
	"github.com/securemist/douyin-mini/service"
	"github.com/securemist/douyin-mini/util"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	download()
	//add()
}

var dir = "./generate/video/"

// 读取./video.txt下的所有url，上传到你的oss上面，同时数据库添加记录
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
		// 这是一个视频的url，得到这个url的文件流
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal("url err in generate :", err)
		}

		src := resp.Body

		playUrl := util.UploadIO(constant.Video_Path, src)
		title := "title" + strconv.FormatInt(int64(index), 10)
		coverUrl := playUrl + "x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600"

		_ = service.AddWork(1, playUrl, coverUrl, title)

		fmt.Println(strconv.FormatInt(int64(index), 10), "的任务完成了")
		index++
	}
}

// 弃用
func add() {

	// 读取目录下的所有文件
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	// 遍历目录下的所有文件，并将文件名添加到切片中
	for index, file := range files {
		go func(index int, file fs.DirEntry) {
			// 排除子目录，只处理文件
			if !file.IsDir() {
				fileName := file.Name()
				filePath := filepath.Join(dir, fileName)

				fmt.Println(fileName)
				data, err := os.ReadFile(filePath)
				if err != nil {
					log.Fatal(err)
				}

				src := bytes.NewReader(data)
				playUrl := util.Upload(constant.Video_Path, fileName, src)
				title := "title" + strconv.FormatInt(int64(index), 10)
				coverUrl := playUrl + "x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600"

				_ = service.AddWork(1, playUrl, coverUrl, title)

				fmt.Println(strconv.FormatInt(int64(index), 10), "的任务完成了")
			}
		}(index, file)
	}
}
