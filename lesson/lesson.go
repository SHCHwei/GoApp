package lesson

import(
    "github.com/gin-gonic/gin"
    "GoApp/member"
    "strconv"
    "net/http"
    db "GoApp/database"
)

type Lesson struct{
    Id int `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
    LessonName string `gorm:"column:lessonName" json:"lessonName"`
    LessonDescribe string `gorm:"column:lessonDescribe" json:"lessonDescribe"`
    Mid string `gorm:"column:mid" json:"mid"`
    Teacher string `gorm:"column:teacher;index" json:"teacher"`
    LessonTime string `gorm:"column:lessonTime" json:"lessonTime"`
    LessonAddress string `gorm:"column:lessonAddress" json:"lessonAddress"`
    TuitionFee string `gorm:"column:tuitionFee" json:"tuitionFee"`
    Email string `gorm:"column:email" json:"email"`
}

var LessonInfo Lesson

func(st Lesson) Create(c *gin.Context){

    if err := c.ShouldBindJSON(&st); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"request error": err.Error()})
		return
	}

    if checkResult := validMemberType(st.Mid); checkResult {

        result := db.MariaDB.Create(&st)

        if result.Error != nil{
            c.JSON(http.StatusBadRequest, gin.H{"database error ":result.Error})
        } else {
            c.JSON(http.StatusOK, gin.H{"data": "1"})
        }
    }else{
        c.JSON(http.StatusBadRequest, gin.H{"request error ": "no Permissions"})
    }
}

func(st Lesson) Update(c *gin.Context){

    if err := c.ShouldBindJSON(&st); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"request error => ": err.Error()})
		return
	}

    if checkResult := validMemberType(st.Mid); checkResult {

        result := db.MariaDB.Save(&st)

        if result.Error != nil{
            c.JSON(http.StatusBadRequest, gin.H{"database error ":result.Error})
        } else {
            c.JSON(http.StatusOK, gin.H{"data": "1"})
        }
    }else{
        c.JSON(http.StatusBadRequest, gin.H{"request error ": "no Permissions"})
    }


}

func(st Lesson) Read(c *gin.Context){

    id := c.Query("id")

    result := db.MariaDB.First(&st, id)

    if result.Error != nil{
        c.JSON(http.StatusBadRequest, gin.H{"error":result.Error})
    } else {
        c.JSON(http.StatusOK, gin.H{"data": st})
    }

}

func(st Lesson) All(c *gin.Context){

    var list []Lesson

    result := db.MariaDB.Find(&list)

    if result.Error != nil{
        c.JSON(http.StatusBadRequest, gin.H{"error":result.Error})
    } else {
        c.JSON(http.StatusOK, gin.H{"data": list})
    }

}

func validMemberType(id string)(bool){

    mid, _ := strconv.Atoi(id)
    data, status := member.FindOne(member.Member{Id: mid})
    return status && data.MType == "school"
}