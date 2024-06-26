package middlewares

// import (
// 	"net/http"

// 	"sosservice/model"
// 	"sosservice/storage"
// 	"sosservice/utils"

// 	"github.com/gin-gonic/gin"
// )

// func AuthenticationMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		tokenString := c.GetHeader("Authentication")

// 		if tokenString == "" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
// 			c.Abort()
// 			return
// 		}

// 		claims, err := utils.ValidateToken(tokenString)
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
// 			c.Abort()
// 			return
// 		}

// 		c.Set("userID", claims.UserID)
// 		c.Next()
// 	}
// }

// func AuthorizationMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 		authHeader := c.GetHeader("Authorization")
// 		if authHeader == "" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not provided"})
// 			c.Abort()
// 			return
// 		}

// 		claims, err := utils.ValidateToken(authHeader)
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
// 			c.Abort()
// 			return
// 		}

// 		var user model.Users
// 		if err := storage.DB.Where("id = ?", claims.UserID).First(&user).Error; err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
// 			c.Abort()
// 			return
// 		}

// 		if !user.IsAdmin {
// 			c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to access this resource"})
// 			c.Abort()
// 			return
// 		}

// 		c.Next()
// 	}
// }
