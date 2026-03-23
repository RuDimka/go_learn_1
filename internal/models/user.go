package models

type User struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Name       string `gorm:"not null" json:"name"`
	Age        int    `json:"age"`
	Password   string `gorm:"not null" json:"password"`
	Email      string `json:"email"`
	Profession string `json:"profession"`
	IsAlive    bool   `gorm:"default:true" json:"is_alive"`
}
