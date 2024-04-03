package do

type FileNode struct {
	Name     string `json:"name"`
	DataId   uint64 `json:"dataId"`
	Parent   uint64 `json:"parent" gorm:"comment: 0 root else other"`
	Kind     byte   `json:"kind" gorm:"default 1;comment: 0 file 1 directory"`
	ShareUrl string `json:"shareUrl"`
	BaseModel
}

func (fn FileNode) TableName() string {
	return "t_file_tree"
}
