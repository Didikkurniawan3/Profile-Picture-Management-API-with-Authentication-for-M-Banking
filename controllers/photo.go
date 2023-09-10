package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	photoRes "github.com/Didik2584/task-5-pbi-btpns-Didik_kurniawan/app/photo"
	"github.com/Didik2584/task-5-pbi-btpns-Didik_kurniawan/helpers"
	"github.com/Didik2584/task-5-pbi-btpns-Didik_kurniawan/models"
	"gorm.io/gorm"
)

type PhotoController struct {
	db *gorm.DB
}

func NewPhotoController(db *gorm.DB) *PhotoController {
	return &PhotoController{db}
}

func (h *PhotoController) Get(c *gin.Context) {
	var userPhoto models.Photo
	err := h.db.Preload("User").First(&userPhoto).Error

	if err != nil {
		response := helpers.ApiResponse(http.StatusBadRequest, "error", nil, "Failed to Get Your Photo")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if userPhoto.PhotoURL == "" {
		response := helpers.ApiResponse(http.StatusBadRequest, "error", nil, "Please Upload Your Photo")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := photoRes.FormatPhoto(&userPhoto, "")
	response := helpers.ApiResponse(http.StatusOK, "success", formatter, "Successfully Fetch User Photo")
	c.JSON(http.StatusOK, response)
}

func (h *PhotoController) Create(c *gin.Context) {
	var userPhoto models.Photo
	var countPhoto int64
	currentUser := c.MustGet("currentUser").(models.User)

	h.db.Model(&userPhoto).Where("user_id = ?", currentUser.ID).Count(&countPhoto)
	if countPhoto > 0 {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helpers.ApiResponse(http.StatusBadRequest, "error", data, "You Already Have a Photo")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input models.Photo
	err := c.ShouldBind(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessages := gin.H{"errors": errors}

		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", errorMessages, "Failed to Upload User Photo")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := c.FormFile("photo_profile")
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessages := gin.H{"errors": errors}

		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", errorMessages, "Failed to Upload User Photo")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	extension := file.Filename
	path := "static/images/" + uuid.New().String() + extension

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", data, "Failed to Upload User Photo")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	h.InsertPhoto(input, path, currentUser.ID)

	data := gin.H{"is_uploaded": true}
	response := helpers.ApiResponse(http.StatusOK, "success", data, "Photo Profile Successfully Uploaded")
	c.JSON(http.StatusOK, response)
}

func (h *PhotoController) InsertPhoto(userPhoto models.Photo, fileLocation string, currUserID int) {
	savePhoto := models.Photo{
		UserID:   currUserID,
		Title:    userPhoto.Title,
		Caption:  userPhoto.Caption,
		PhotoURL: fileLocation,
	}

	h.db.Debug().Create(&savePhoto)
}

func (h *PhotoController) Update(c *gin.Context) {
	var userPhoto models.Photo
	currentUser := c.MustGet("currentUser").(models.User)

	err := h.db.Where("user_id = ?", currentUser.ID).First(&userPhoto).Error
	if err != nil {
		response := helpers.ApiResponse(http.StatusBadRequest, "error", err, "Photo Profile Failed to Update")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input models.Photo
	err = c.ShouldBind(&input)
	if err != nil {
		response := helpers.ApiResponse(http.StatusBadRequest, "error", err, "Photo Profile Failed to Update")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := c.FormFile("update_profile")
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", data, "Failed to Update User Photo")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	extension := file.Filename
	path := "static/images/" + uuid.New().String() + extension

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		response := helpers.ApiResponse(http.StatusBadRequest, "error", err, "Photo Profile Failed to Upload")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	h.UpdatePhoto(input, &userPhoto, path)

	data := photoRes.FormatPhoto(&userPhoto, "regular")
	response := helpers.ApiResponse(http.StatusOK, "success", data, "Photo Profile Successfully Updated")
	c.JSON(http.StatusOK, response)
}

func (h *PhotoController) UpdatePhoto(oldPhoto models.Photo, newPhoto *models.Photo, path string) {
	newPhoto.Title = oldPhoto.Title
	newPhoto.Caption = oldPhoto.Caption
	newPhoto.PhotoURL = path

	h.db.Save(newPhoto)
}

func (h *PhotoController) Delete(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(models.User)
	var userPhoto models.Photo

	err := h.db.Where("user_id = ?", currentUser.ID).Delete(&userPhoto).Error
	if err != nil {
		data := gin.H{
			"is_deleted": false,
		}

		response := helpers.ApiResponse(http.StatusBadRequest, "error", data, "Failed to delete user photo")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"is_deleted": true,
	}

	response := helpers.ApiResponse(http.StatusOK, "success", data, "User Photo Successfully Deleted")
	c.JSON(http.StatusOK, response)
}
