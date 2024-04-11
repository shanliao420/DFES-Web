package request

const (
	OpCreate = int8(0)
	OpDelete = int8(1)
	OpUpdate = int8(2)
)

type FileSystemOperateInfo struct {
	ID           uint64
	Name         string
	Parent       uint64
	Kind         byte
	OperatorCode int8
}

type UploadFileInfo struct {
}
