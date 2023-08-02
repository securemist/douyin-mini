package service

import (
	"github.com/securemist/douyin-mini/model/resp"
	"github.com/securemist/douyin-mini/util"
	"sort"
)

// @Description: 关注操作
// @param userId 关注者id
// @param toUserId 被关注者id
func RelationAction(userId, toUserId int64) {
	Db := util.GetDbConnection()
	r := 0
	_ = Db.QueryRow("SELECT COUNT(id) FROM user_follow WHERE user_id = ? AND follow_user_id = ?",
		userId, toUserId).Scan(&r)

	// 数据库没有记录，是第一次关注
	var err error
	if r == 0 {
		_, err = Db.Exec("INSERT INTO user_follow VALUES (?, ?, ?, ?, ?)",
			int64(0), toUserId, userId, util.TimeNow(), 0)
	} else { // 后续的关注或者取消关注
		_, err = Db.Exec("UPDATE user_follow SET deleted = CASE WHEN deleted = 0 THEN 1 WHEN deleted = 1 THEN 0 END WHERE user_id = ? AND follow_user_id = ?",
			userId, toUserId)
	}

	HandleSqlError(err)
}

func GetFollowList(currentUserId, userId int64) []resp.User {
	Db := util.GetDbConnection()
	defer Db.Close()

	var idList []int64
	err := Db.Select(&idList, "SELECT `follow_user_id` FROM user_follow WHERE user_id = ? AND deleted = 0 ORDER BY create_time DESC", userId)
	HandleSqlError(err)

	var userList []resp.User
	for _, id := range idList {
		user := GetUserInfo(currentUserId, id)
		userList = append(userList, user)
	}

	return userList
}

func GetFollowerList(currentUserId, userId int64) []resp.User {
	Db := util.GetDbConnection()
	defer Db.Close()

	var idList []int64
	err := Db.Select(&idList, "SELECT `user_id` FROM user_follow WHERE follow_user_id = ? AND deleted = 0 ORDER BY create_time DESC", userId)
	HandleSqlError(err)

	var userList []resp.User
	for _, id := range idList {
		user := GetUserInfo(currentUserId, id)
		userList = append(userList, user)
	}

	return userList
}

// @Description: 查找朋友列表 同时查出一个人的关注列表和粉丝列表，取交集即可
// @param currentUserId 当前用户id
// @param userId 查找的用户id
// @return []resp.User
func GetFriendList(currentUserId, userId int64) []resp.User {
	Db := util.GetDbConnection()
	defer Db.Close()

	var idList1 []int64
	err := Db.Select(&idList1, "SELECT `follow_user_id` FROM user_follow WHERE user_id = ? AND deleted = 0 ORDER BY create_time DESC", userId)
	HandleSqlError(err)

	var idList2 []int64
	err = Db.Select(&idList2, "SELECT `user_id` FROM user_follow WHERE follow_user_id = ? AND deleted = 0 ORDER BY create_time DESC", userId)
	HandleSqlError(err)

	// 取两个数组的交集 不去重
	idList := findIntersection(idList1, idList2)

	var userList []resp.User

	for _, id := range idList {
		user := GetUserInfo(currentUserId, id)
		userList = append(userList, user)

	}
	return userList
}

func findIntersection(list1, list2 []int64) []int64 {
	// 先对两个切片进行排序
	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})
	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})

	var intersection []int64
	i, j := 0, 0

	// 遍历两个排序后的切片，找到交集
	for i < len(list1) && j < len(list2) {
		if list1[i] < list2[j] {
			i++
		} else if list1[i] > list2[j] {
			j++
		} else {
			// 找到交集元素，加入结果切片
			intersection = append(intersection, list1[i])
			i++
			j++
		}
	}

	return intersection
}
