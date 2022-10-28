package middlewares

import (
	"assignment_4/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unathorized",
				"message": err.Error(),
			})
			return
		}
		c.Set("userData", verifyToken)
		// fmt.Printf("%s", userData)
		c.Next()
	}
}

// func UserAunthorization() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		db := database.GetDB()
// 		userId, err := strconv.Atoi(c.Param("userId"))
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 				"error":   "Bad Request",
// 				"message": "invalid parameter",
// 			})
// 			return
// 		}
// 		// userData := c.MustGet("userData").(jwt.MapClaims)
// 		userID := uint(userData["id"].(float64))
// 		User := models.User{}

// 		err = db.Select("id").First(&User, uint(userId)).Error

// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
// 				"error":   "Data not found",
// 				"message": "data doesn't exist",
// 			})
// 			return
// 		}

// 		if User.ID != userID {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"error":   "Unauthorized",
// 				"message": "you are not allowed to access this data",
// 			})
// 			return
// 		}

// 		c.Next()
// 	}
// }
