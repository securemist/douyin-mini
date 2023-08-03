package config

import (
	"bufio"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var Port = 8080
var Default_feed_list_size = 25
var Db_url = ""

var endpoint string
var acessKeyId string
var accessKeySecret string
var bucketName string

// 这里的oss连接全局共用一个，不知道会不会有问题 TODO
var Bucket *oss.Bucket
var Oss_url_Suffix string

var conf = make(map[string]string, 100)

func init() {
	err := readProperties("./conf")
	if err != nil {
		log.Fatal("conf read error : ", err)
	}

	loadConf()
}

// 读取配置文件，检验数据库与oss的连接
func loadConf() {

	core()

	mysql()

	ossConf()
}

func core() {
	if conf["port"] != "" {
		Port, _ = strconv.Atoi(conf["port"])
	}

	if conf["default_feed_list_size"] != "" {
		Default_feed_list_size, _ = strconv.Atoi(conf["default_feed_list_size"])
	}
}

func mysql() {

	//url=root:root@tcp(127.0.0.1:3306)/douyim
	Db_url = conf["username"] + ":" + conf["password"] + "@tcp(" + conf["addr"] + ")/" +
		"" + conf["database"] + "?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"

	db, err := sqlx.Open("mysql", Db_url)
	if err != nil {
		log.Fatal("database connection failed")
	}
	defer db.Close()

	// 校验数据库是否可连接
	err = db.Ping()
	if err != nil {
		log.Fatal("database connection failed")
	}

}

func ossConf() {

	endpoint = conf["endpoint"]
	acessKeyId = conf["acessKeyId"]
	accessKeySecret = conf["accessKeySecret"]
	bucketName = conf["bucketName"]

	client, err := oss.New(endpoint, acessKeyId, accessKeySecret)
	if err != nil {
		log.Fatal("oss client error :", err)
	}

	Bucket, err = client.Bucket(bucketName)
	if err != nil {
		log.Fatal("oss client error :", err)
	}

	Oss_url_Suffix = "https://" + bucketName + "." + endpoint + "/"
}

func readProperties(file string) error {

	// 1. 读取文件，得到文件句柄
	open, err := os.Open(file)
	defer open.Close() // 关闭关文件

	if err != nil {
		return err
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
			return err
		}

		if len(line) == 0 || string(line[0]) == "#" {
			continue
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

		conf[key] = value // 添加到map中，key为map的key，value为map的value
	}

	return nil
}
