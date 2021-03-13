package models

import (
	"errors"

	grpc "github.com/userq11/grpc-test/protobufs"
	"golang.org/x/crypto/bcrypt"
)

// User describes User model
type User struct {
	ID        int64  `json:"id" xorm:"'id' pk autoincr" schema:"id"`
	Email     string `json:"email" xorm:"email" schema:"email" validate:"required,contains=@"`
	Password  string `json:"-" xorm:"password" schema:"password" validate:"required,gt=8"`
	FirstName string `json:"first_name" xorm:"first_name" schema:"first_name"`
	LastName  string `json:"last_name" xorm:"last_name" schema:"last_name"`
	Visible   bool   `json:"visible" xorm:"visible" schema:"visible"`
}

// TempUser describes temp user for creating new user
// just to deal with password confirmation
type TempUser struct {
	Email           string `json:"email" validate:"required,contains=@"`
	Password        string `json:"-" validate:"required,gt=8"`
	ConfirmPassword string `json:"-" validate:"required,gt=8"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
}

// NewUser creates new user
func NewUser(u *TempUser) (*User, error) {
	if u.Password != u.ConfirmPassword {
		return nil, errors.New("password and confirm password do not match")
	}

	user := &User{
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Visible:   true,
	}

	if err := user.SetPassword(u.Password); err != nil {
		return nil, err
	}

	return user, nil
}

// SetPassword hashes given password and sets it on User
func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)
	return nil
}

// ComparePassword compares a given password with stored one
func (u *User) ComparePassword(password string) error {
	if !u.Visible {
		return errors.New("User is inactive")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return errors.New("invalid email or password")
	}

	return nil
}

// TableName returns table name (for xorm)
func (u *User) TableName() string {
	return "users"
}

func (u *User) ToProtobuf() *grpc.User {
	nu := new(grpc.User)

	nu.Id = u.ID
	nu.Email = u.Email
	nu.FirstName = u.FirstName
	nu.LastName = u.LastName
	nu.Visible = u.Visible

	return nu
}
