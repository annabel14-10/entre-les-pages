package repository

import (
	"auth-service/domain"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	result := r.db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, domain.ErrUserNotFound
		}
		return nil, result.Error
	}

	return &user, nil
}

func (r *userRepo) Create(user *domain.User) error {
	result := r.db.Create(user)

	if result.Error != nil {
		return domain.ErrUserAlreadyExists
	}

	return nil
}

func (r *userRepo) GetByID(id uint) (*domain.User, error) {
    var user domain.User
    result := r.db.First(&user, id)
    if result.Error != nil {
        return nil, domain.ErrUserNotFound
    }
    return &user, nil
}