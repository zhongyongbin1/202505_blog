package model

type User struct {
	Base
	Username string `gorm:"type:varchar(255);unique;not null"`
	Password string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);unique;not null"`
	Phone    string `gorm:"type:varchar(255);unique;not null"`
	role     string `gorm:"type:varchar(255);not null"`
	status   int    `gorm:"type:int(11);default:1"`
}
