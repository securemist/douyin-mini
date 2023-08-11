package service

import (
	"fmt"
	"github.com/securemist/douyin-mini/model/db"
	"github.com/securemist/douyin-mini/model/resp"
	"github.com/securemist/douyin-mini/util"
	"log"
	"strconv"
)

func AddUser(user db.User) int64 {
	Db := util.GetDbConnection()

	defer Db.Close()

	r, err := Db.Exec("INSERT INTO sys_user (username, password) VALUES (?, ?)", user.Username, user.Password)
	if err != nil {
		fmt.Println(err)
	}

	id, err := r.LastInsertId()

	if err != nil {
		fmt.Println(err)
	}

	return id
}

func GetUserId(username string, password string) (int64, error) {

	Db := util.GetDbConnection()

	defer Db.Close()

	var id int64 = 0
	err := Db.QueryRow("SELECT id FROM sys_user WHERE username=? AND password=?", username, password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetUserById(userId int64) db.User {
	Db := util.GetDbConnection()

	defer Db.Close()

	var user db.User
	err := Db.Get(&user, "SELECT * FROM sys_user WHERE id = ?", userId)
	if err != nil {
		log.Println("sql exec failed in GetUserById : ", err)
	}
	return user
}

// @Description:  获取用户的详细信息
// @param currentUserId 当前的用户id
// @param userId 查看的用户id
// @return resp.User  查询不到用户信息会返回一个id为0的空用户
func GetUserInfo(currentUserId int64, userId int64) resp.User {
	Db := util.GetDbConnection()
	defer Db.Close()

	// 获取用户信息
	var user db.User
	_ = Db.Get(&user, "SELECT * FROM sys_user WHERE id = ?", userId)

	// 关注总数
	var followCount int
	_ = Db.QueryRow("SELECT COUNT(id) FROM user_follow WHERE user_id = ?", userId).Scan(&followCount)

	// 粉丝总数
	var followerCount int
	_ = Db.QueryRow("SELECT Count(id) FROM user_follow WHERE follow_user_id = ?", userId).Scan(&followerCount)

	// 判断当前用户有没有关注指定用户
	var isFollow bool
	id := 0
	_ = Db.QueryRow("SELECT COUNT(id) FROM user_follow WHERE user_id=? AND follow_user_id=?", currentUserId, userId).Scan(&id)
	isFollow = id != 0

	// 获赞数量
	var totalFavorited int
	_ = Db.QueryRow("SELECT COUNT(id) FROM user_favorite WHERE work_id IN (SELECT id FROM user_work WHERE user_id = ?)", userId).Scan(&totalFavorited)

	// 作品数量
	var workCount int
	_ = Db.QueryRow("SELECT COUNT(id) FROM user_work WHERE user_id = ?", userId).Scan(&workCount)

	// 喜欢数
	var favoriteCount int
	_ = Db.QueryRow("SELECT COUNT(id) FROM user_favorite WHERE user_id = ?", userId).Scan(&favoriteCount)

	user0 := resp.User{
		Id:              user.Id,
		Name:            user.Name,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		Signature:       user.Signature,

		IsFollow:       isFollow,
		FollowCount:    followCount,
		FollowerCount:  followerCount,
		TotalFavorited: strconv.Itoa(totalFavorited),
		WorkCount:      workCount,
		FavoriteCount:  favoriteCount,
	}
	return user0
}
