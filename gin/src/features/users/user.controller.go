package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


type UserController struct{
	userService *UserService;
}


func NewUserController(userService *UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}



func (controller *UserController) getUserById(c *gin.Context){
	// This is the place to put the business logic
	// for the user controller.
	// This is a good place to put the database connection
	// and other dependencies that are needed for the application.
	// This is also a good place to put the configuration
	// for the application.

	var userId = c.Param("id")
	controller.userService.GetUserById(userId);
	println("user controller")
	c.JSON(200, gin.H{
		"userId": userId,
	})
}


func (controller *UserController) create(c *gin.Context){
	var user UserRequestDto
	if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	controller.userService.Create(&user)
}


func (controller *UserController) getUsers(c *gin.Context){
	var offset, _ = strconv.ParseInt(c.DefaultQuery("offset", "0"), 10, 0);
	var limit, _ = strconv.ParseInt(c.DefaultQuery("limit", "10"), 10, 0);
	var users = controller.userService.Get(int(offset), int(limit))
	c.JSON(200, users);
}


func (controller *UserController) RegisterRoute(route *gin.RouterGroup){
	route.GET("/user/:id", controller.getUserById)
	route.GET("/user", controller.getUsers)
	route.POST("/user", controller.create)
}