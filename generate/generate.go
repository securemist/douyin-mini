package main

import (
	"bufio"
	"fmt"
	"github.com/securemist/douyin-mini/model/constant"
	"github.com/securemist/douyin-mini/service"
	"github.com/securemist/douyin-mini/util"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync"
)

func main() {
	err := download()

	if err != nil {
		log.Fatal("generate data error : ", err)
	}
}

// 读取./video.txt下的所有url，上传到你的oss上面，同时数据库添加记录
func download() error {
	open, err := os.Open("./video.txt")
	defer open.Close()

	if err != nil {
		return err
	}

	content := bufio.NewReader(open)

	fmt.Println("during generate data")

	task := sync.WaitGroup{}

	for {
		line, _, err := content.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		task.Add(1)
		go func(line []byte) {
			url := string(line)
			// 这是一个视频的url，得到这个url的文件流
			resp, err := http.Get(url)
			if err != nil {
				panic("utl get error")
			}

			src := resp.Body

			// 上传到你自己的oss
			playUrl := util.UploadIO("video/", src)
			title := "这是我的作品标题哇"
			coverUrl := playUrl + constant.COVER_SUFFIX

			_ = service.AddWork(int64(rand.Intn(5)), playUrl, coverUrl, title)

			fmt.Print("==")
			task.Done()
		}(line)
	}

	task.Wait()

	fmt.Print("\n")
	fmt.Println("data generate success")
	return nil
}
