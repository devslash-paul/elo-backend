package models

import (
	"errors"
	"time"
)

type User struct {
	ID        uint
	DeletedAt *time.Time `xml:"-" json:"-"`
	CreatedAt *time.Time `xml:"-" json:"-"`
	UpdatedAt *time.Time `xml:"-" json:"-"`
	FullName  string
}

type UserController struct {
	db *ExportDB
}

func NewUserController(s *ExportDB) *UserController {
	return &UserController{s}
}

func (u *User) Validate() *ModelError {
	err := new(ModelError)
	err.Errors = []ModelValidation{}

	if u.FullName == "" {
		err.Errors = append(err.Errors, ModelValidation{
			Field: "FullName",
			Error: "This field is required",
		})
	}

	if len(err.Errors) > 0 {
		return err
	}
	return nil
}

// GetUserByID will get a user by a specified user id
func (u *UserController) GetUserByID(id uint64) (*User, error) {
	user := &User{}
	u.db.First(&user, id)

	if user.ID == 0 {
		return nil, errors.New("No user has been found with that ID")
	}

	return user, nil
}

func (u *UserController) Delete(user *User) error {
	u.db.Delete(user)
	return nil
}

func (u *UserController) Create(user *User) error {
	u.db.Create(user)
	return nil
}
