package models

type AlbumImage struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement"`
	AlbumID   int64  `gorm:"column:album_id;not null"`
	CreatedAt int64  `gorm:"column:created_at;not null;default:0"`
	TgFile    string `gorm:"column:tg_file;not null"`
}

func (AlbumImage) TableName() string {
	return "album_images"
}
