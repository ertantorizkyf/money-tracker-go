package dto

type RegisterReq struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	DOB      string `json:"dob"`
	Password string `json:"password"`
}

type LoginReq struct {
	UsernameOrEmail string `json:"username_or_email"`
	Password        string `json:"password"`
}
