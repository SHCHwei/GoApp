package middleware

import(
    "github.com/gin-gonic/gin"
    "net/http"
    "GoApp/pkg/jwt"
    "strings"
)

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
        auth := c.GetHeader("Authorization")
        token := strings.Split(auth, "Bearer ")

        if len(token) > 1 && jwt.Verify(token[1]) {
            c.Next()
        }else{
            c.JSON(http.StatusBadRequest, gin.H{"error": "JWT Verify Failed"})
            c.Abort()
        }
    }
}

