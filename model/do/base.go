package do

import "time"

type BaseModel struct {
	ID       uint64    `gorm:"primarykey" json:"id"`
	CreateAt time.Time `gorm:"autoCreateTime" json:"createAt"`
	UpdateAt time.Time `gorm:"autoUpdateTime" json:"updateAt"`
	DeleteAt time.Time `gorm:"index;default:null" json:"-"`
}
