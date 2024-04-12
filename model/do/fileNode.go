package do

const (
	FileKind      = byte(0)
	DirectoryKind = byte(1)
)

type FileNode struct {
	Name     string `json:"name"`
	DataId   string `json:"dataId" gorm:"default null"`
	FileSize uint64 `json:"fileSize" gorm:"default 0"`
	Parent   uint64 `json:"parent" gorm:"comment: 0 root else other"`
	Kind     byte   `json:"kind" gorm:"default 1;comment: 0 file 1 directory"`
	ShareUrl string `json:"shareUrl" gorm:"default null"`
	BaseModel
}

func (fn FileNode) TableName() string {
	return "t_file_tree"
}
