package do

type UserTreeRoot struct {
	BaseModel
	UserId     uint64 `json:"userId"`
	TreeRootId uint64 `json:"treeRootId"`
}

func (utr UserTreeRoot) TableName() string {
	return "r_user_file_root"
}
