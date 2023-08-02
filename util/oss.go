package util

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func handleError(err error) {
	if err != nil {
		log.Fatal("oss client error:", err)
		os.Exit(-1)
	}
}

var bucket *oss.Bucket

var endpoint string
var acessKeyId string
var accessKeySecret string
var bucketName string

func init() {
	data := readProperties("./oss.properties")

	endpoint = data["endpoint"]
	acessKeyId = data["acessKeyId"]
	accessKeySecret = data["accessKeySecret"]
	bucketName = data["bucketName"]

	client, err := oss.New(endpoint, acessKeyId, accessKeySecret)
	handleError(err)

	bucket, err = client.Bucket(bucketName)
	handleError(err)
}

// @Description: 上传文件
// @param path 不同的类型对应了不同的上传路径，具体参见 model/constant/oss.go
// @param fileName 文件名
// @param file 文件流
// @param userId 用户id
// @return string 返回上传之后的URL，上传失败返回 ""
func Upload(path string, fileName string, file io.Reader) string {
	// 文件路径 + snowflake id + 扩展名
	objectName := path + strconv.FormatInt(GenerateId(), 10) + filepath.Ext(fileName)
	err := bucket.PutObject(objectName, file)

	if err != nil {
		log.Println("file upload failed : ", err)
		return ""
	}

	return "https://" + bucketName + "." + endpoint + "/" + objectName
}

// 从网络下载的文件流，直接上传，不考虑文件扩展名，全部为mp4
func UploadIO(path string, file io.Reader) string {
	objectName := path + strconv.FormatInt(GenerateId(), 10) + ".mp4"
	err := bucket.PutObject(objectName, file)
	if err != nil {
		log.Println("file upload failed : ", err)
		return ""
	}

	return "https://" + bucketName + "." + endpoint + "/" + objectName
}

func Test() {
	fmt.Println(bucket)
}
