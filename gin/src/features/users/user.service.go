package users

import models "myapp/src/model"

type UserService struct {
	userRepo *UserRepository
}

func NewUserService(userRepository *UserRepository) *UserService {
	return &UserService{
		userRepo: userRepository,
	}
}

func (user *UserService) GetUserById(userId string) (string, error){
	// This is the place to put the business logic
	// for the user service.
	// This is a good place to put the database connection
	// and other dependencies that are needed for the application.
	// This is also a good place to put the configuration
	// for the application.

	println("user service")

	return userId, nil;
}


func (user *UserService) Create(newUser *UserRequestDto){
	var usr = models.User{
		ID: newUser.ID,
		Name: newUser.Name,
		Email: newUser.Email,
	};
	user.userRepo.Create(&usr);
}

func (user *UserService) Get(offset int, limit int) *[]models.User{
	var usr = user.userRepo.Find(offset, limit)
	return usr;
}