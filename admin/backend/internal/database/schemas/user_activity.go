package schemas

import "time"

type UserActivity struct {
	ID     uint      `gorm:"primaryKey"`
	UserID uint      `gorm:"not null;index"`
	Date   time.Time `gorm:"not null;index"`
}
