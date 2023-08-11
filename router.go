package main

import (
	"github.com/gin-gonic/gin"
	"github.com/securemist/douyin-mini/controller"
	"github.com/securemist/douyin-mini/middleware"
	"net/http"
)

func initRouter(r *gin.Engine) {

	r.Static("/static", "./public")

	r.LoadHTMLGlob("templates/**")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "说明"})
	})

	// 不需要鉴权的接口
	registerAndLogin := r.Group("/douyin")
	registerAndLogin.POST("/user/register/", controller.Register)
	registerAndLogin.POST("/user/login/", controller.Login)

	apiRouter := r.Group("/douyin")
	apiRouter.Use(middleware.JWT())
	// 基础接口
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/user/", controller.UserInfo)
	apiRouter.POST("/publish/action/", controller.Publish)
	apiRouter.GET("/publish/list/", controller.PublishList)

	// 互动接口
	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
	apiRouter.POST("/comment/action/", controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.HandleCommentList)

	// 社交接口
	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
	apiRouter.GET("/relation/friend/list/", controller.FriendList)
	apiRouter.GET("/message/chat/", controller.MessageChat)
	apiRouter.POST("/message/action/", controller.MessageAction)
}
