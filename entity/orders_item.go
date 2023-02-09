package entity

import "time"

type OrderItem struct {
	Id        int        `gorm:"primaryKey;autoIncrement;column:id;->;<-:create"`
	Name      string     `gorm:"column:name;size:256;->;<-:create"`
	Price     uint       `gorm:"column:price;->;<-:create"`
	ExpiredAt time.Time  `gorm:"column:expired_at;->;<-:create"`
	CreatedAt time.Time  `gorm:"column:created_at;autoCreateTime;->;<-:create"`
	UpdatedAt time.Time  `gorm:"column:updated_at;autoUpdateTime;->;<-:create"`
	DeletedAt *time.Time `gorm:"column:deleted_at;->;<-:create"`
}
