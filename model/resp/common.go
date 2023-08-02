package resp

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title,omitempty"`
	CreateTime    string `json:"-"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	Id              int64  `db:"id" json:"id"`
	Name            string `db:"name" json:"name"`
	Avatar          string `db:"avatar" json:"avatar"`
	BackgroundImage string `db:"background_image" json:"background_image"`
	Signature       string `db:"signature" json:"signature"`

	FollowCount    int    `db:"follow_count" json:"follow_count"`
	FollowerCount  int    `db:"follower_count" json:"follower_count"`
	IsFollow       bool   `db:"is_follow" json:"is_follow"`
	TotalFavorited string `db:"total_favorited" json:"total_favorited"`
	WorkCount      int    `db:"work_count" json:"work_count"`
	FavoriteCount  int    `db:"favorite_count" json:"favorite_count"`
}

type Message struct {
	Id         int64  `json:"id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}

type MessageSendEvent struct {
	UserId     int64  `json:"user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId int64  `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}
