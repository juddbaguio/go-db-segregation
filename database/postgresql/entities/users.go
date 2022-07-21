package entities

type User struct {
	ID        int    `json:"id" gorm:"primaryKey;column:id"`
	FirstName string `json:"firstName" gorm:"column:firstName"`
	LastName  string `json:"lastName" gorm:"column:lastName"`
	Age       int    `json:"age" gorm:"column:age"`
}
