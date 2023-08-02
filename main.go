package main

import (
	"github.com/gin-gonic/gin"
	"github.com/securemist/douyin-mini/service"
)

func main() {
	go service.RunMessageServer()

	r := gin.Default()

	r.MaxMultipartMemory = 100 << 20

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
