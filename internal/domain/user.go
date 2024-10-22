package domain

type User struct {
	ID          string `json:"user_id" redis:"user_id"`
	Email       string `json:"email" redis:"email"`
	Username    string `json:"username" redis:"username"`
	Password    string `json:"password" redis:"password"`
	Phonenumber string `json:"phonenumber" redis:"phonenumber"`
}
