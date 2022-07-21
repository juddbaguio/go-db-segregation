package model

type User struct {
	ID        interface{} `json:"id"`
	FirstName string      `json:"firstName"`
	LastName  string      `json:"lastName"`
	Age       int         `json:"age"`
}
