package do

type UserTreeRoot struct {
	BaseModel
	UserId     uint64 `json:"userId"`
	TreeRootId uint64 `json:"treeRootId"`
}
