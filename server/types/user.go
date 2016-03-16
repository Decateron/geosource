package types

type User struct {
	ID     string `json:"id" gorm:"column:u_userid"`
	Name   string `json:"name" gorm:"column:u_username"`
	Avatar string `json:"avatar" gorm:"column:u_avatar"`
	Email  string `json:"email" gorm:"column:u_email"`
}

// Returns the name of the user's corresponding table in the database.
func (user User) TableName() string {
	return "users"
}
