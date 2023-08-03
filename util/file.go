package util

import (
	"io"
	"mime/multipart"
	"os"
)

// @Description:  把文件保存到本地
// @param videoFileName 文件路径 + 文件名
// @param file 文件流
func Upload(videoFileName string, file multipart.File) error {
	videoFile, err := os.Create(videoFileName)
	if err != nil {
		return err
	}

	defer videoFile.Close()

	_, err = io.Copy(videoFile, file)
	if err != nil {
		return err
	}

	return nil
}
