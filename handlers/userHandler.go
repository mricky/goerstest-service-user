package handlers

import (
	"goers_service_user/helpers"
	"goers_service_user/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// tidak perlu interface, tidak ada dependecy

type usersHandler struct {
	userService user.Service // ini panggil interface service
}

func NewUserHandler(userService user.Service) *usersHandler {
	return &usersHandler{userService}
}

func (h *usersHandler) UserById(c *gin.Context) {

	var input user.UserInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		response := helpers.APIResponse("Fetch Failed", http.StatusUnprocessableEntity, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := h.userService.GetUserById(input)

	if err != nil {
		response := helpers.APIResponse("Fetch Failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse("User Detail", http.StatusOK, "success", user)
	c.JSON(http.StatusOK, response)

}
func (h *usersHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helpers.FormatError(err)
		errorMessage := gin.H{"errors": errors} // map

		response := helpers.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUserInput(input)

	if err != nil {
		response := helpers.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormaterUser(newUser, "token")

	response := helpers.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *usersHandler) Login(c *gin.Context) {

	var input user.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helpers.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	logginUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormaterUser(logginUser, "tokentoken")

	response := helpers.APIResponse("Success fully login", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}
