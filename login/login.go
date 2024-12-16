package login

import(
    "github.com/gin-gonic/gin"
    "net/http"
    "GoApp/pkg/jwt"
    "GoApp/pkg/hashing"
    "GoApp/teacher"
    "GoApp/student"
)


type Login struct{
    Account string `json:"account"`
    Pwd string `json:"password"`
}

var LoginInfo Login

func(st Login) UnSignIn(c *gin.Context){
    s := hashing.HashPassword("4654a9gr6ag")
    c.JSON(200, gin.H{"token": s})
}

func(st Login) SignIn(c *gin.Context){

    if err := c.ShouldBindJSON(&st); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"request error": err.Error()})
		return
	}

    st.Pwd = hashing.HashPassword(st.Pwd)

    if studentInfo, ok := student.FindOne(student.Student{Account: st.Account, PW : st.Pwd }); ok {
        if token, _ := jwt.Build(studentInfo.Id, studentInfo.Account); token != ""{
            c.Set("sid", studentInfo.Id)
            c.Set("account", studentInfo.Account)

            c.JSON(200, gin.H{"token": token})
        }else{
            c.JSON(403, gin.H{"error": "Login Failed(token error)"})
            return
        }

    } else{
        c.JSON(403, gin.H{"error": "Login Failed"})
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
        if token, _ := jwt.Build(teacherInfo.Id, teacherInfo.Account); token != ""{
            c.Set("tid", teacherInfo.Id)
            c.Set("account", teacherInfo.Account)

            c.JSON(200, gin.H{"token": token})
        }else{
            c.JSON(403, gin.H{"error": "Login Failed(token error)"})
            return
        }

    } else{
        c.JSON(403, gin.H{"error": "Login Failed"})
        return
    }
}