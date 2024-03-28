package db

type User struct {
	Id       uint   `gorm:"primaryKey"`
	Username string `gorm:"column:v0"`
	Password string `gorm:"column:v0"`
}
