package entity

type Todo struct {
	UserId    string `json:"userId"`
	Title     string `json:"title"`
	Completed int    `json:"completed"`
}
