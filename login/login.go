package login

import(
    "github.com/gin-gonic/gin"
    "net/http"
    "GoApp/pkg/session"
    "GoApp/pkg/hashing"
    "GoApp/teacher"
    "GoApp/student"
    db "GoApp/database"
    "strconv"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "context"
)


type Login struct{
    Account string `json:"account"`
    Pwd string `json:"password"`
}

var LoginInfo Login
var session_student = db.Mongo.Database.Collection("session_student")
var session_teacher = db.Mongo.Database.Collection("session_teacher")

func init(){
    go session.DeleteSession(session_student)
}


func(st Login) SignOut(c *gin.Context){

    cookieStr, _ := c.Cookie("session_id")

    cookie, _ := primitive.ObjectIDFromHex(cookieStr)

    _, err := session_student.DeleteOne(context.TODO(), bson.D{{"_id", cookie}})

    if err == nil{
        c.SetCookie("session_id", "", -1, "/", "localhost", false, true)
    }else{
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }

}

func(st Login) SignIn(c *gin.Context){

    if err := c.ShouldBindJSON(&st); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"request error": err.Error()})
		return
	}

    st.Pwd = hashing.HashPassword(st.Pwd)

    if studentInfo, ok := student.FindOne(student.Student{Account: st.Account, PW : st.Pwd }); ok {

        doc := session.InitSession(strconv.Itoa(studentInfo.Id))

        if _, err := session_student.InsertOne(context.TODO(), doc) ; err == nil {

            c.SetCookie("session_id", doc.ID.Hex(), 3600, "/", "localhost", false, true)

            c.JSON(200, gin.H{"message": "login successful"})
        }else{
            c.JSON(403, gin.H{"message": "DB Failed", "err": err, "doc": doc})
            return
        }

    } else{
        c.JSON(403, gin.H{"message": "Login Failed"})
        return
    }
}

func(st Login) TeacherSignIn(c *gin.Context){

    if err := c.ShouldBindJSON(&st); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"request error": err.Error()})
		return
	}

    st.Pwd = hashing.HashPassword(st.Pwd)

    if teacherInfo, ok := teacher.FindOne(teacher.Teacher{Account: st.Account, PW : st.Pwd }); ok {

        doc := session.InitSession(strconv.Itoa(teacherInfo.Id))

        if _, err := session_teacher.InsertOne(context.TODO(), doc) ; err == nil {

            c.SetCookie("session_id", doc.ID.Hex(), 3600, "/", "localhost", false, true)

            c.JSON(200, gin.H{"message": "login successful"})
        }else{
            c.JSON(403, gin.H{"message": "DB Failed", "err": err, "doc": doc})
            return
        }

    } else{
        c.JSON(403, gin.H{"message": "Login Failed"})
        return
    }
}