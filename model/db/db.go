package db

type User struct {
	Id              int64  `db:"id"`
	Name            string `db:"name"`
	Avatar          string `db:"avatar"`
	BackgroundImage string `db:"background_image"`
	Signature       string `db:"signature"`
	Username        string `db:"username"`
	Password        string `db:"password"`
}

type UserFollow struct {
	Id           int64 `db:"id"`
	UserId       int64 `db:"user_id"`
	FollowUserId int64 `db:"follow_user_id"`
}

type Work struct {
	Id         int64  `db:"id"`
	UserId     int64  `db:"user_id"`
	PlayUrl    string `db:"play_url"`
	CoverUrl   string `db:"cover_url"`
	Title      string `db:"title"`
	CreateTime string `db:"create_time"`
}

type WorkFavorite struct {
	Id     int64 `db:"id"`
	UserId int64 `db:"user_id"`
	WorkId int64 `db:"work_id"`
}

type WorkComment struct {
	Id      int64  `db:"id"`
	UserId  int64  `db:"user_id"`
	WorkId  int64  `db:"work_id"`
	Content string `db:"content"`
}
