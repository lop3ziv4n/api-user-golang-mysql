package data

import (
	"github.com/lop3ziv4n/api-user-golang-mysql/models"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	C *gorm.DB
}

func (r *UserRepository) Create(user *models.User) {
	r.C.Create(&user)
}

func (r *UserRepository) GetAll() []models.User {
	var users []models.User
	r.C.Find(users)
	return users
}

func (r *UserRepository) GetById(id string) models.User {
	var user models.User
	r.C.First(user, id)
	return user
}

func (r *UserRepository) GetAllByName(name string) []models.User {
	var users []models.User
	r.C.Where("name = ?", name).Find(users)
	return users
}

func (r *UserRepository) Delete(id string) {
	var user models.User
	user.ID = id
	r.C.Delete(&user)
}

func (r *UserRepository) Update(id string, user *models.User) {
	r.C.Where("id = ?", id).Save(&user)
}
