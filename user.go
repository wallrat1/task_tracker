package todo

type User struct {
	Id   int    `json:"-"`
	Name string `json:"name"`
	User string `json:"user"`
	Pass string `json:"password"`
}
