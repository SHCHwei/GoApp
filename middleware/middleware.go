package middleware

import(
    "github.com/gin-gonic/gin"
    "net/http"
    "context"

    db "GoApp/database"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"

)



func CheckSLogin() gin.HandlerFunc {
	return func(c *gin.Context) {

        var session_user = db.Mongo.Database.Collection("session_student")

        cookie, err := c.Cookie("session_id")

        if err != nil {
          c.JSON(http.StatusBadRequest, gin.H{"error": "Please login again"})
          c.Abort()
        }

        objID, _ := primitive.ObjectIDFromHex(cookie)
        filter := bson.D{{"_id", objID}}

        if count, _ := session_user.CountDocuments(context.TODO(), filter); count == 0 {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Please login again"})
            c.Abort()
        }

        c.Next()

    }
}


func CheckTLogin() gin.HandlerFunc {
	return func(c *gin.Context) {

        var session_user = db.Mongo.Database.Collection("session_teacher")

        cookie, err := c.Cookie("session_id")

        if err != nil {
          c.JSON(http.StatusBadRequest, gin.H{"error": "Please login again s"})
          c.Abort()
        }

        objID, _ := primitive.ObjectIDFromHex(cookie)
        filter := bson.D{{"_id", objID}}

        if count, _ := session_user.CountDocuments(context.TODO(), filter); count == 0 {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Please login again"})
            c.Abort()
        }

        c.Next()

    }
}