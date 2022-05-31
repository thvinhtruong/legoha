package repository

import (
	"github.com/jinzhu/gorm"
	entity "github.com/thvinhtruong/legoha/app/domain/entities"
)

func NewUser(user *entity.User) *entity.User {
	u := entity.User{}
	u.ID = user.ID
	u.Name = user.Name
	u.Username = user.Username
	u.Password = user.Password
	u.Role = user.Role
	u.CreatedAt = user.CreatedAt
	return &u
}

func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

// add a user
func (r *Repository) CreateUser(entityUser *entity.User) error {
	entityUser.Role = "user"
	u := NewUser(entityUser)

	err := r.DB.Create(&u).Error
	if err != nil {
		return err
	}

	return nil
}

// get all user
func (r *Repository) ListUsers() ([]*entity.User, error) {
	var users []entity.User

	err := r.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}

	result := make([]*entity.User, 0, len(users))
	for _, user := range users {
		result = append(result, NewUser(&user))
	}

	return result, nil
}

// get user by id
func (r *Repository) GetUserByID(id int) (*entity.User, error) {
	var user entity.User

	err := r.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return NewUser(&user), nil
}

// get user by username
func (r *Repository) GetUserByUsername(username string) (*entity.User, error) {
	var user entity.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return NewUser(&user), nil
}

// update user infor
func (r *Repository) PatchUser(id int, u *entity.User) error {
	var user entity.User
	err := r.DB.First(&user, id).Error
	if err != nil {
		return err
	}
	err = r.DB.Model(&user).Updates(u).Error
	if err != nil {
		return err
	}
	return nil
}

// delete user by id
func (r *Repository) DeleteUser(u *entity.User) error {

	err := r.DB.First(&u, u.ID).Error
	if err != nil {
		return err
	}

	err = r.DB.Delete(&u).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) LoginUser(u *entity.User) (*entity.User, error) {
	var user entity.User
	err := r.DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	if err != nil {
		return nil, err
	}

	return NewUser(&user), nil
}