package types

type User struct {
	ID     string `json:"id" gorm:"column:u_userid"`
	Name   string `json:"name" gorm:"column:u_username"`
	Avatar string `json:"avatar" gorm:"column:u_avatar"`
	Email  string `json:"email" gorm:"column:u_email"`
}

// TableName returns the name of User's corresponding table in the database.
func (User) TableName() string {
	return "users"
}
