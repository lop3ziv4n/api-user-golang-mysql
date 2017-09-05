package models

type (
	User struct {
		ID       string `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
		Name     string `gorm:"column:name;type:varchar(45);not null" json:"name"`
		LastName string `gorm:"column:lastname;type:varchar(45);not null" json:"lastname"`
	}
)
