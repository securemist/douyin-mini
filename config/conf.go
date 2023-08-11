package conf

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
)

var Port = "8080"
var ip = "127.0.0.1"

// 静态资源的url前缀，加载数据库url字段前面返回给前端就可以了
var Project_url_suffix = "http://" + ip + ":" + Port

// 视频流单词请求返回视频数
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

	// 读取配置文件
	loadConf()
}

// 读取配置文件，检验数据库与oss的连接
func loadConf() {

	core()

	mysql()

}

func core() {
	if conf["port"] != "" {
		Port = conf["port"]
	}

	if conf["default_feed_list_size"] != "" {
		Default_feed_list_size, _ = strconv.Atoi(conf["default_feed_list_size"])
	}
}

func mysql() {

	//url=root:root@tcp(127.0.0.1:3306)/douyin
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
