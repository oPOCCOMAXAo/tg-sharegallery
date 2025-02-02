package models

type User struct {
	ID        int64 `gorm:"column:id;primaryKey;autoIncrement"`
	CreatedAt int64 `gorm:"column:created_at;not null;default:0"`
	UpdatedAt int64 `gorm:"column:updated_at;not null;default:0"`
	TgID      int64 `gorm:"column:tg_id;not null;index:idx_tg_id,unique"`
}

func (User) TableName() string {
	return "users"
}
