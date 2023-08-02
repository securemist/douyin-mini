package service

import (
	"github.com/securemist/douyin-mini/model/db"
	"github.com/securemist/douyin-mini/model/resp"
	"github.com/securemist/douyin-mini/util"
	"sync"
)

var chatConnMap = sync.Map{}

func HandleMessageAction(userId, toUserId int64, content string) {
	Db := util.GetDbConnection()
	defer Db.Close()

	_, err := Db.Exec("INSERT INTO user_message VALUES (?, ?, ?, ?, ?, ?)",
		int64(0), userId, toUserId, content, util.TimeNow(), 0)
	HandleSqlError(err)
}

// 这个接口有问题，前端每s都会发一次请求，直接鬼畜
func GetMessageChat(fromUserId, toUserId int64) []resp.Message {
	Db := util.GetDbConnection()
	defer Db.Close()

	var list []db.Message
	err := Db.Select(&list, "SELECT id, from_user_id, to_user_id, content, create_time FROM user_message  WHERE id in (SELECT id FROM user_message WHERE from_user_id = ? AND to_user_id = ? AND deleted = 0 UNION SELECT id FROM user_message WHERE from_user_id = ? AND to_user_id = ? AND deleted = 0) ORDER BY create_time DESC", fromUserId, toUserId, toUserId, fromUserId)
	HandleSqlError(err)

	var messageList []resp.Message
	for _, message0 := range list {
		message := resp.Message{
			message0.Id,
			message0.FromUserId,
			message0.ToUserId,
			message0.Content,
			util.TimeStringToUnix(message0.CreateTime),
			//util.TimeFormat(util.UnixToTime(util.TimeStringToUnix(message0.CreateTime))),
		}
		messageList = append(messageList, message)
	}
	return messageList
}
