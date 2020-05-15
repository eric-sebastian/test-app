package user_model

//User :
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//UserForm :
type UserForm struct {
	Name string `json:"name"`
}
