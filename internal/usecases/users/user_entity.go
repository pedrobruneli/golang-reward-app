package users

type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Username string `gorm:"unique" json:"username"`
}

type UserDTO struct {
	Username string `json:"username"`
}
