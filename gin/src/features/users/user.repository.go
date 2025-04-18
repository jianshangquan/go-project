package users

import (
	models "myapp/src/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}


func (userRepo *UserRepository) FindById(id uint){
	userRepo.db.Where("id = ?", id).Find(models.User{ID: id})
}

func (userRepo *UserRepository) Create(user *models.User){
	userRepo.db.Create(user);
}

func (userRepo *UserRepository) Find(offset int, limit int) *[]models.User{
	var users []models.User;
	userRepo.db.Offset(offset).Limit(limit).Find(&users);
	return &users;
}