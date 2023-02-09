package entity

import "time"

type User struct {
	Id         int        `gorm:"primaryKey;autoIncrement;column:id;->;<-:create"`
	FullName   string     `gorm:"column:full_name;size:256;->;<-:create"`
	FirstOrder string     `gorm:"column:first_order;size:256;->;<-:create"`
	CreatedAt  time.Time  `gorm:"column:created_at;autoCreateTime;->;<-:create"`
	UpdatedAt  time.Time  `gorm:"column:updated_at;autoUpdateTime;->;<-:create"`
	DeletedAt  *time.Time `gorm:"column:deleted_at;->;<-:create"`
}
