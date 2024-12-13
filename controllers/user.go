package controllers

import (
	"net/http"
	"strconv"

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

// Register user
func (h *UserController) Register(c *gin.Context) {
	var user models.User

	// Bind JSON input ke struct user
	if err := c.ShouldBindJSON(&user); err != nil {
		errors := helpers.FormatValidationError(err)
		response := helpers.ApiResponse(http.StatusBadRequest, "error", gin.H{"errors": errors}, "Invalid input")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Hash password
	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		response := helpers.ApiResponse(http.StatusInternalServerError, "error", nil, "Failed to hash password")
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	user.Password = hashedPassword

	// Simpan user ke database
	if err := h.db.Create(&user).Error; err != nil {
		if helpers.IsDuplicateError(err) {
			response := helpers.ApiResponse(http.StatusConflict, "error", nil, "Email already exists")
			c.JSON(http.StatusConflict, response)
			return
		}
		response := helpers.ApiResponse(http.StatusInternalServerError, "error", nil, "Failed to create user")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// Format respons
	formatter := userRes.FormatUserResponse(user, "")
	response := helpers.ApiResponse(http.StatusOK, "success", formatter, "User registered successfully")
	c.JSON(http.StatusOK, response)
}

// Login user
func (h *UserController) Login(c *gin.Context) {
	var user models.User
	var userInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind input JSON ke struct userInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		response := helpers.ApiResponse(http.StatusBadRequest, "error", nil, "Invalid input")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Cari user berdasarkan email
	if err := h.db.Debug().Where("email = ?", userInput.Email).First(&user).Error; err != nil {
		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", nil, "Login Failed")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Verifikasi password
	comparePass := helpers.ComparePassword(user.Password, userInput.Password)
	if !comparePass {
		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", nil, "Invalid email or password")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Generate token
	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", nil, "Login Failed")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Format user response dengan token
	formatter := userRes.FormatUserResponse(user, token)
	response := helpers.ApiResponse(http.StatusOK, "success", formatter, "User Login Successfully")
	c.JSON(http.StatusOK, response)
}


// Update user
func (h *UserController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		response := helpers.ApiResponse(http.StatusBadRequest, "error", nil, "Invalid user ID")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var oldUser models.User
	if err := h.db.First(&oldUser, id).Error; err != nil {
		response := helpers.ApiResponse(http.StatusNotFound, "error", nil, "User not found")
		c.JSON(http.StatusNotFound, response)
		return
	}

	var input struct {
		Name     string `json:"name,omitempty"`
		Email    string `json:"email,omitempty"`
		Password string `json:"password,omitempty"`
	}

	// Bind JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		response := helpers.ApiResponse(http.StatusBadRequest, "error", nil, "Invalid input")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Update hanya atribut yang diisi
	if input.Password != "" {
		hashedPassword, _ := helpers.HashPassword(input.Password)
		input.Password = hashedPassword
	}

	if err := h.db.Model(&oldUser).Updates(input).Error; err != nil {
		response := helpers.ApiResponse(http.StatusInternalServerError, "error", nil, "Failed to update user")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helpers.ApiResponse(http.StatusOK, "success", nil, "User updated successfully")
	c.JSON(http.StatusOK, response)
}

// Delete user
func (h *UserController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		response := helpers.ApiResponse(http.StatusBadRequest, "error", nil, "Invalid user ID")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var user models.User
	if err := h.db.First(&user, id).Error; err != nil {
		response := helpers.ApiResponse(http.StatusNotFound, "error", nil, "User not found")
		c.JSON(http.StatusNotFound, response)
		return
	}

	if err := h.db.Delete(&user).Error; err != nil {
		response := helpers.ApiResponse(http.StatusInternalServerError, "error", nil, "Failed to delete user")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helpers.ApiResponse(http.StatusOK, "success", nil, "User deleted successfully")
	c.JSON(http.StatusOK, response)
}
