package controllers

import (
	"assignment_4/database"
	"assignment_4/helpers"
	"assignment_4/models.go"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateSosmed(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	// sosmed := models.User{}
	contentType := helpers.GetContentType(c)

	sosmed := models.Sosmed{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&sosmed)
	} else {
		c.ShouldBind(&sosmed)
	}

	sosmed.UserID = userID

	err := db.Debug().Create(&sosmed).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, sosmed)
}

func GetSosmed(c *gin.Context) {
	var (
		result gin.H
	)
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	sosmed := models.Sosmed{}
	user := models.User{}
	userID := uint(userData["id"].(float64))
	db.Find(&user).Where("id = ?", userID)
	// userName := userData["full_name"].(string)

	fmt.Println(userID)
	db.Find(&sosmed)
	// sosmed.User.ID = userID
	// sosmed.User.FullName = userName
	sosmed.User = &user
	result = gin.H{
		"social_media": sosmed,
		// "user":         user,
	}
	c.JSON(http.StatusOK, result)
}

func UpdateSosmed(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	sosmed := models.Sosmed{}

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&sosmed)
	} else {
		c.ShouldBind(&sosmed)
	}

	sosmed.UserID = userID
	sosmed.ID = uint(socialMediaId)

	err := db.Model(&sosmed).Where("id = ?", socialMediaId).Updates(models.Sosmed{Nama: sosmed.Nama, SosmedURL: sosmed.SosmedURL}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, sosmed)
}

func Deletesosmed(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("socialMediaId")
	selectedID, _ := strconv.Atoi(id)

	sosmed := models.Sosmed{}

	db.Debug().Delete(&sosmed, selectedID)

	c.JSON(http.StatusOK, gin.H{
		// "err":     "Bad Request",
		"message": "Your social media has been succesfully deleted",
	})
}
