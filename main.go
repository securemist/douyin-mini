package main

import (
	"github.com/gin-gonic/gin"
	conf "github.com/securemist/douyin-mini/config"
)

func main() {
	r := gin.Default()

	r.MaxMultipartMemory = 100 << 20
	initRouter(r)

	r.Run(":" + conf.Port)
}
