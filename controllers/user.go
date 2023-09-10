package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	userRes "github.com/Didik2584/task-5-pbi-btpns-Didik_kurniawan/app/user"
	"github.com/Didik2584/task-5-pbi-btpns-Didik_kurniawan/helpers"
	"github.com/Didik2584/task-5-pbi-btpns-Didik_kurniawan/models"
	"gorm.io/gorm"
)

type UserController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{db}
}

func (h *UserController) Register(c *gin.Context) {
	var user models.User
	c.ShouldBindJSON(&user)

	user.Password = helpers.HashPassword(user.Password)

	if err := h.db.Debug().Create(&user).Error; err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", errorMessage, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := userRes.FormatUserResponse(user, "")
	response := helpers.ApiResponse(http.StatusOK, "success", formatter, "User Registered Successfully")
	c.JSON(http.StatusOK, response)
}

func (h *UserController) Login(c *gin.Context) {
	var user models.User

	c.ShouldBindJSON(&user)

	Inputpassword := user.Password
	if err := h.db.Debug().Where("email = ?", user.Email).First(&user).Error; err != nil {
		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", nil, "Login Failed")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	comparePass := helpers.ComparePassword(user.Password, Inputpassword)
	if !comparePass {
		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", nil, "Login Failed")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", nil, "Login Failed")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := userRes.FormatUserResponse(user, token)
	response := helpers.ApiResponse(http.StatusOK, "success", formatter, "User Login Successfully")
	c.JSON(http.StatusOK, response)
}

func (h *UserController) Update(c *gin.Context) {
	var oldUser models.User
	var newUser models.User

	id := c.Param("userId")

	if err := h.db.First(&oldUser, id).Error; err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", errorMessage, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err := json.NewDecoder(c.Request.Body).Decode(&newUser); err != nil {
		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", nil, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err := h.db.Model(&oldUser).Updates(newUser).Error; err != nil {
		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", nil, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.ApiResponse(http.StatusOK, "success", nil, "User Updated Successfully")
	c.JSON(http.StatusOK, response)
}

func (h *UserController) Delete(c *gin.Context) {
	var user models.User

	id := c.Param("userId")
	if err := h.db.First(&user, id).Error; err != nil {
		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", nil, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err := h.db.Delete(&user).Error; err != nil {
		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", nil, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.ApiResponse(http.StatusOK, "success", nil, "User Deleted Successfully")
	c.JSON(http.StatusOK, response)
}
