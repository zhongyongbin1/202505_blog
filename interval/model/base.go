package model

type Base struct {
	ID        uint  `gorm:"primarykey"`
	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
	DeletedAt int64 `gorm:"index"`
}
