package types

type User struct {
	Id     string `json:"id" gorm:"column:u_userid"`
	Name   string `json:"name" gorm:"column:u_username"`
	Avatar string `json:"avatar" gorm:"column:u_avatar"`
	Email  string `json:"email" gorm:"column:u_email"`
}

func (user User) TableName() string {
	return "users"
}
