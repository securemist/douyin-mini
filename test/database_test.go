package test

type User struct {
	Id              int64
	Name            string
	Avatar          string
	BackgroundImage string
	Signature       string
	Username        string
	Password        string
}

func (User) TableName() string {
	return "sys_user"
}
