package util

import (
	"bufio"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"io"
	"log"
	"os"
	"strings"
)

var url string

func init() {
	data := readProperties("./mysql.properties")

	//url=root:root@tcp(127.0.0.1:3306)/douyim
	url = data["username"] + ":" + data["password"] + "@tcp(" + data["addr"] + ")/" +
		"" + data["database"] + "?charset=utf8&parseTime=true&loc=Asia%2FShanghai"
}

func readProperties(file string) map[string]string {
	allConfig := make(map[string]string)

	// 1. 读取文件，得到文件句柄
	open, err := os.Open(file)
	defer open.Close() // 关闭关文件

	if err != nil {
		fmt.Println("err1", err)
	}

	// 2. 读取文件内容
	content := bufio.NewReader(open)
	for {
		// 3. 按行读取文件内容
		line, _, err := content.ReadLine()
		if err != nil {
			if err == io.EOF { // 去读到结尾，就跳出循环读取
				break
			}
			return nil
		}
		// 4. 处理每一行读取到的文件内容
		s := strings.TrimSpace(string(line)) // 去掉左右空格
		index := strings.Index(s, "=")       // 因为配置是=，找到=的索引位置
		if index < 0 {
			continue
		}

		key := strings.TrimSpace(s[:index]) // 截取=左侧的值为key
		if len(key) == 0 {
			continue
		}

		value := strings.TrimSpace(s[index+1:]) // 截取=右侧的为value
		if len(value) == 0 {
			continue
		}

		allConfig[key] = value // 添加到map中，key为map的key，value为map的value
	}
	return allConfig
}

func GetDbConnection() *sqlx.DB {
	db, err := sqlx.Open("mysql", url)
	if err != nil {
		log.Fatal("database connection failed")
	}
	return db
}
