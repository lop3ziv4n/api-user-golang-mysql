package models

import "strconv"

// User entity
type (
	User struct {
		ID       int64  `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
		Name     string `gorm:"column:name;type:varchar(45);not null" json:"name"`
		LastName string `gorm:"column:lastname;type:varchar(45);not null" json:"lastname"`
	}
)

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (u *User) GetID() string {
	return strconv.FormatInt(u.ID, 10)
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (u *User) SetID(id string) error {
	u.ID, _ = strconv.ParseInt(id, 10, 64)
	return nil
}
