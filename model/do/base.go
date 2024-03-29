package do

import "time"

type BaseModel struct {
	ID       uint64    `gorm:"primarykey" json:"ID"`
	CreateAt time.Time `gorm:"autoCreateTime"`
	UpdateAt time.Time `gorm:"autoUpdateTime"`
	DeleteAt time.Time `gorm:"index;default:null" json:"_"`
}
