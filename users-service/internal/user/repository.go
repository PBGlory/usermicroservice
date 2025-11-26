package user

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *Repository) GetByID(id uint) (*User, error) {
	var u User
	err := r.db.First(&u, id).Error
	return &u, err
}

func (r *Repository) List() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *Repository) Update(id uint, email, password string) (*User, error) {
	var u User
	if err := r.db.First(&u, id).Error; err != nil {
		return nil, err
	}

	u.Email = email
	u.Password = password
	err := r.db.Save(&u).Error
	return &u, err
}

func (r *Repository) Delete(id uint) error {
	return r.db.Delete(&User{}, id).Error
}
