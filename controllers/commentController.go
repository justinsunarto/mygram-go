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

func CreateComments(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	fmt.Println(userID)

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Comment)
}

func GetComment(c *gin.Context) {
	db := database.GetDB()
	comment := models.Comment{}

	db.Find(&comment)
	comment.User = &models.User{}

	c.JSON(http.StatusOK, comment)
}

func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	comment := models.Comment{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&comment)
	} else {
		c.ShouldBind(&comment)
	}

	comment.UserID = userID
	comment.ID = uint(photoId)

	err := db.Model(&comment).Where("id = ?", photoId).Updates(models.Comment{Message: comment.Message}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, comment)
}

func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("commentId")
	selectedID, _ := strconv.Atoi(id)

	comment := models.Comment{}

	db.Debug().Delete(&comment, selectedID)

	c.JSON(http.StatusOK, gin.H{
		// "err":     "Bad Request",
		"message": "Your comment has been succesfully deleted",
	})
}
