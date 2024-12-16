package teacher

import(
    "github.com/gin-gonic/gin"
    "net/http"
    "GoApp/pkg/hashing"
    db "GoApp/database"
)

type Teacher struct{
    Id int `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
    Account string `gorm:"column:account" json:"account"`
    PW string `gorm:"column:password" json:"password"`
    TeacherName string `gorm:"column:teacherName" json:"teacherName"`
    Email string `gorm:"column:email;index" json:"email"`
    Phone string `gorm:"column:phone" json:"phone"`
}

var TeacherInfo Teacher


func(st Teacher) Create(c *gin.Context){

    if err := c.ShouldBindJSON(&st); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    st.PW = hashing.HashPassword(st.PW)
    result := db.MariaDB.Create(&st)

    if result.Error != nil{
        c.JSON(http.StatusBadRequest, gin.H{"error":result.Error})
    } else {
        c.JSON(http.StatusOK, gin.H{"data": "1"})
    }
}

func(st Teacher) Update(c *gin.Context){

    if err := c.ShouldBindJSON(&st); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"request error => ": err.Error()})
		return
	}

    st.PW = hashing.HashPassword(st.PW)
    result := db.MariaDB.Save(&st)

    if result.Error != nil{
        c.JSON(http.StatusBadRequest, gin.H{"database update error => ":result.Error})
    } else {
        c.JSON(http.StatusOK, gin.H{"data": st})
    }
}

func(st Teacher) Read(c *gin.Context){

    id := c.Query("id")

    result := db.MariaDB.First(&st, id)

    if result.Error != nil{
        c.JSON(http.StatusBadRequest, gin.H{"error":result.Error})
    } else {
        c.JSON(http.StatusOK, gin.H{"data": st})
    }

}

func All(c *gin.Context){

    var data []Teacher

    result := db.MariaDB.Find(&data)

    if result.Error != nil{
        c.JSON(http.StatusBadRequest, gin.H{"error":result.Error})
    } else if result.RowsAffected == 0 {
        c.JSON(http.StatusOK, gin.H{"data": "null"})
    } else {
        c.JSON(http.StatusOK, gin.H{"data": data})
    }

}


func FindOne(st Teacher)(Teacher, bool){

    reData := Teacher{}

    result := db.MariaDB.Where(&st).First(&reData)

    if result.RowsAffected > 0 {
        return reData, true
    }else{
        return reData, false
    }
}