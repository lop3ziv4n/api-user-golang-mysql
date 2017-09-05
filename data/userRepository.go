package data

import (
	"fmt"
	"strconv"

	"github.com/lop3ziv4n/api-user-golang-mysql/models"

	"github.com/jinzhu/gorm"
)

// UserRepository user repository
type UserRepository struct {
	C *gorm.DB
}

// GetAll users
func (r *UserRepository) GetAll() []models.User {
	var users []models.User
	r.C.Order("id").Find(&users)
	return users
}

// GetByID user by id
func (r *UserRepository) GetByID(id string) (models.User, error) {
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return models.User{}, fmt.Errorf("Id must be integer: %s", id)
	}
	var user models.User
	if r.C.First(&user, intID).RecordNotFound() {
		return models.User{}, fmt.Errorf("User for id %s not found", id)
	}
	return user, nil
}

// GetAllByName users by name
func (r *UserRepository) GetAllByName(name string) []models.User {
	var users []models.User
	r.C.Where("name = ?", name).Find(&users)
	return users
}

// Create user
func (r *UserRepository) Create(user *models.User) {
	r.C.Create(&user)
}

// Delete user
func (r *UserRepository) Delete(id string) error {
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return fmt.Errorf("Id must be integer: %s", id)
	}
	var user models.User
	if r.C.First(&user, intID).RecordNotFound() {
		return fmt.Errorf("User with id %s does not exist", id)
	}
	r.C.Delete(&user)
	return nil
}

// Update user
func (r *UserRepository) Update(id string, u *models.User) error {
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return fmt.Errorf("Id must be integer: %s", id)
	}
	var user models.User
	if r.C.First(&user, intID).RecordNotFound() {
		return fmt.Errorf("User with id %s does not exist", id)
	}
	user.Name = u.Name
	user.LastName = u.LastName
	r.C.Save(&user)
	return nil
}
