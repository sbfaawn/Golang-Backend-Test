package entity

import "time"

type OrderHistory struct {
	Id           int       `gorm:"primaryKey;autoIncrement;column:id;->;<-:create"`
	UserId       int       `gorm:"column:user_id;->;<-:create"`
	OrderItemId  int       `gorm:"column:order_item_id;->;<-:create"`
	Descriptions string    `gorm:"column:description;type:text;->;<-:create"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime;->;<-:create"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime;->;<-:create"`
}
