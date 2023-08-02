package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/securemist/douyin-mini/model/constant"
	"github.com/securemist/douyin-mini/model/db"
	"github.com/securemist/douyin-mini/model/resp"
	"github.com/securemist/douyin-mini/service"
	"github.com/securemist/douyin-mini/util"
	"log"
	"net/http"
	"strconv"
)

type UserInfoResponse struct {
	resp.Response
	*resp.User `json:"user"`
}

type UserLoginResponse struct {
	resp.Response
	UserId *int64  `json:"user_id"`
	Token  *string `json:"token"`
}

type UserResponse struct {
	resp.Response
	*resp.User `json:"user"`
}

type UserRegisterResponse struct {
	resp.Response
	UserId *int64  `json:"user_id"`
	Token  *string `json:"token"`
}

func Register(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")

	user := db.User{
		Username: username,
		Password: password,
	}

	// 校验用户名是否存在 前端文档没有明确是否有这一步 TODO
	//userNameExist := service.checkUserName(username)
	//if userNameExist {
	//	c.JSON(http.StatusOK, UserRegisterResponse{
	//		Response: constant.USERNAME_EXISTED,
	//		userId:   0,
	//		token:    token,
	//	})
	//}

	// 创建用户，签发token
	id := service.AddUser(user)

	token, err := util.GenerateToken(id)
	if err != nil { // 生成token异常
		log.Println("token generate error", err)
		c.JSON(http.StatusOK, UserRegisterResponse{
			Response: constant.GENERAL_ERROR,
			UserId:   nil,
			Token:    nil,
		})
	}

	c.JSON(http.StatusOK, UserRegisterResponse{
		Response: constant.USER_REGISTER_SUCCESS,
		UserId:   &id,
		Token:    &token,
	})
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := ""
	// 验证用户名密码
	userId, err := service.GetUserId(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			constant.USER_LOGIN_FAILED,
			nil,
			nil,
		})
		return
	}

	// 每次登录更新token
	token, err = util.GenerateToken(userId)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			constant.GENERAL_ERROR,
			nil,
			nil,
		})
		return
	}

	c.JSON(http.StatusOK, UserLoginResponse{
		constant.GENERAL_SUCCESS,
		&userId,
		&token,
	})
}

func UserInfo(c *gin.Context) {
	id, _ := c.Get("userId")
	currentUserId := id.(int64)
	// 要查看的用户id
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		// 请求参数异常
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: constant.REQUEST_PARAM_ERROR,
		})
		return
	}

	user := service.GetUserInfo(currentUserId, userId)
	if user.Id == 0 { // 用户不存在
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: constant.REQUEST_PARAM_ERROR,
		})
	} else {
		// 返回成功结果
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: constant.GENERAL_SUCCESS,
			User:     &user,
		})
	}

}
