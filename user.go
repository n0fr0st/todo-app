package todo

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"-"`
	Username string `json:"username"`
	Password string `json:"password"`
}
