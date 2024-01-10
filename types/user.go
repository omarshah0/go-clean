package types

type UserType string

const (
	Admin    UserType = "admin"
	Customer UserType = "customer"
)

type User struct {
	BaseModel
	Name  string   `json:"name" gorm:"not null"`
	Email string   `json:"email" gorm:"unique;not null"`
	Type  UserType `json:"type" gorm:"not null"`
}
