package models

type Album struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement"`
	PublicID  *int64 `gorm:"column:public_id;null;default:null;index:idx_public_id,unique"`
	CreatedAt int64  `gorm:"column:created_at;not null;default:0"`
	UpdatedAt int64  `gorm:"column:updated_at;not null;default:0"`
	DeletedAt *int64 `gorm:"column:deleted_at"`
	OwnerID   int64  `gorm:"column:owner_id;not null"`
	Title     string `gorm:"column:title;not null;default:'';size:255"`
	Saved     bool   `gorm:"column:saved;not null;default:false"`
}

func (Album) TableName() string {
	return "albums"
}

type AlbumDomain struct {
	Album

	ImagesCount int64 `gorm:"column:images_count"`

	PublicLink string `gorm:"-"`
}
