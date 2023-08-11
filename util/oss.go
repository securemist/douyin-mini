package util

import (
	conf "github.com/securemist/douyin-mini/config"
	"io"
	"log"
	"path/filepath"
	"strconv"
)

// @Description: 上传文件
// @param path 不同的类型对应了不同的上传路径，具体参见 model/constant/oss.go
// @param fileName 文件名
// @param file 文件流
// @param userId 用户id
// @return string 返回上传之后的URL，上传失败返回 ""
func Upload(path string, fileName string, file io.Reader) string {
	// 文件路径 + snowflake id + 扩展名
	objectName := path + strconv.FormatInt(GenerateId(), 10) + filepath.Ext(fileName)
	err := conf.Bucket.PutObject(objectName, file)

	if err != nil {
		log.Println("file upload failed : ", err)
		return ""
	}

	return conf.Oss_url_Suffix + objectName
}

// 从网络下载的文件流，直接上传，不考虑文件扩展名，全部为mp4
func UploadIO(path string, file io.Reader) string {
	objectName := path + strconv.FormatInt(GenerateId(), 10) + ".mp4"
	err := conf.Bucket.PutObject(objectName, file)
	if err != nil {
		log.Println("file upload failed : ", err)
		return ""
	}

	return conf.Oss_url_Suffix + objectName
}
