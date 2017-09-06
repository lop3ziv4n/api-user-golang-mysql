package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/lop3ziv4n/api-user-golang-mysql/common"
	"github.com/lop3ziv4n/api-user-golang-mysql/data"

	"github.com/gorilla/mux"
)

// Handler for HTTP Get - "/users"
// Returns all User documents
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Create new context
	context := NewContext()
	defer context.Close()
	c := context.Db()
	repo := &data.UserRepository{c}
	// Get all users form repository
	users := repo.GetAll()
	j, err := json.Marshal(UsersResource{Data: users})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler for HTTP Get - "/users"
// Return User document by id
func GetUserById(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]
	// Create new context
	context := NewContext()
	defer context.Close()
	c := context.Db()
	repo := &data.UserRepository{c}
	// Get users form repository
	user, err_db := repo.GetByID(id)
	if err_db != nil {
		common.DisplayAppError(w, err_db, "An unexpected error has occurred", 500)
		return
	}
	// Create response data
	j, err := json.Marshal(UserResource{Data: user})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler for HTTP Get - "/users"
// Return User document by id
func GetUserByName(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	name := vars["name"]
	// Create new context
	context := NewContext()
	defer context.Close()
	c := context.Db()
	repo := &data.UserRepository{c}
	// Get all users form repository
	users := repo.GetAllByName(name)
	// Create response data
	j, err := json.Marshal(UsersResource{Data: users})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler for HTTP Post - "/users"
// Create a new User document
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var dataResource UserResource
	// Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid User data", 500)
		return
	}
	user := &dataResource.Data
	// Create new context
	context := NewContext()
	defer context.Close()
	c := context.Db()
	repo := &data.UserRepository{c}
	// Create User
	repo.Create(user)
	// Create response data
	j, err := json.Marshal(dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler for HTTP Put - "/users"
// Update a User document by id
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]
	var dataResource UserResource
	// Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid User data", 500)
		return
	}
	user := &dataResource.Data
	// Create new context
	context := NewContext()
	defer context.Close()
	c := context.Db()
	repo := &data.UserRepository{c}
	// Update user by id
	err_db := repo.Update(id, user)
	if err_db != nil {
		common.DisplayAppError(w, err_db, "An unexpected error has occurred", 500)
		return
	}
	// Create response data
	j, err := json.Marshal(dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler for HTTP Delete - "/users/{id}"
// Delete a User document by id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]
	// Create new context
	context := NewContext()
	defer context.Close()
	c := context.Db()
	repo := &data.UserRepository{c}
	// Remove user by id
	err := repo.Delete(id)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	// Send response back
	w.WriteHeader(http.StatusNoContent)
}
