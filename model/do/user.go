package do

type UserModel struct {
	BaseModel
	Username string `json:"userName" gorm:"index"`
	Password string `json:"-" `
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Enable   int8   `json:"enable" gorm:"comment: 0 开启 1 冻结"`
}

func (um *UserModel) TableName() string {
	return "t_user"
}
