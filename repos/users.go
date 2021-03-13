package repos

import (
	"errors"

	"github.com/userq11/grpc-test/models"
	"xorm.io/xorm"
)

// UsersRepo describes users repo interface
type UsersRepo interface {
	Create(*models.User) error
	FindByID(int64) (*models.User, error)
	FindByEmail(string) (*models.User, error)
	Update(*models.User) (err error)
	Delete(*models.User) (err error)
}

// NewUsersRepo returns new users repo
func NewUsersRepo(db *xorm.Engine) UsersRepo {
	return &usersRepo{db: db}
}

type usersRepo struct {
	db *xorm.Engine
}

// Create creates user, returns error if needed
func (u usersRepo) Create(user *models.User) error {
	if err := models.Validate(user); err != nil {
		return err
	}

	if _, err := u.db.Insert(user); err != nil {
		return err
	}

	return nil
}

// FindByID returns user found by id, error if not found
func (u usersRepo) FindByID(id int64) (*models.User, error) {
	if id <= 0 {
		return nil, errors.New("bad ID")
	}

	user := new(models.User)
	exists, err := u.db.ID(id).Get(user)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("could not find user")
	}

	return user, nil
}

// FindByEmail returns user found by email, error if not found
func (u usersRepo) FindByEmail(email string) (*models.User, error) {
	if len(email) == 0 {
		return nil, errors.New("bad email")
	}

	user := new(models.User)
	user.Email = email
	exists, err := u.db.Get(user)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("could not find user")
	}

	return user, nil
}

func (u usersRepo) Update(user *models.User) error {
	if user == nil || user.ID <= 0 {
		return errors.New("user values invalid")
	}

	_, err := u.db.ID(user.ID).Update(user)
	if err != nil {
		return err
	}

	return nil
}

func (u usersRepo) Delete(user *models.User) error {
	if user == nil || user.ID <= 0 {
		return errors.New("user values invalid")
	}

	_, err := u.db.ID(user.ID).Delete(user)
	if err != nil {
		return err
	}

	return nil
}
