package request

type RegisterInfo struct {
	Username string
	Password string
	Email    string
	Phone    string
}

type LoginInfo struct {
	Username string
	Password string
}
