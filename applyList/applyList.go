package applyList

import(
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    "GoApp/student"
    db "GoApp/database"
)

type Apply struct{
    Id int `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
    Sid string `gorm:"column:sid" json:"sid"`
    Lid string `gorm:"column:lid" json:"lid"`
    Payed string `gorm:"column:payed;default:0" json:"payed"`
}


type ApplyMember struct{
    Id int `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
    Sid string `gorm:"column:sid" json:"sid"`
    Lid string `gorm:"column:lid" json:"lid"`
    Payed string `gorm:"column:payed" json:"payed"`
    LessonName string `gorm:"column:lessonName" json:"lessonName"`
    TuitionFee string `gorm:"column:tuitionFee" json:"tuitionFee"`
    Email string `gorm:"column:email" json:"email"`
}



var ApplyList Apply

func(st Apply) Create(c *gin.Context){

    if err := c.ShouldBindJSON(&st); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


    sid, _ := strconv.Atoi(st.Sid)
    _ , status := student.FindOne(student.Student{Id: sid})

    if status == false {
        c.JSON(http.StatusBadRequest, gin.H{"error": "student not exist"})
        return
    }

    var applylist Apply

    result := db.MariaDB.Table("apply_lists").Where("sid = ? AND lid = ?", st.Sid, st.Lid).First(&applylist)

    if result.RowsAffected > 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "student applied"})
        return
    }

    result = db.MariaDB.Table("apply_lists").Create(&st)

    if result.Error != nil{
        c.JSON(http.StatusBadRequest, gin.H{"error":result.Error})
    } else {
        c.JSON(http.StatusOK, gin.H{"data": "1"})
    }
}


func(st Apply) Update(c *gin.Context){

    if err := c.ShouldBindJSON(&st); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    sid, _ := strconv.Atoi(st.Sid)
    _ , status := student.FindOne(student.Student{Id: sid})

    if status == false {
        c.JSON(http.StatusBadRequest, gin.H{"error": "student not exist"})
        return
    }

    var applylist Apply

    result := db.MariaDB.Table("apply_lists").Where("sid = ? AND lid = ?", st.Sid, st.Lid).First(&applylist)

    if result.RowsAffected == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "student not apply lesson"})
        return
    }

    result = db.MariaDB.Table("apply_lists").Save(&st)

    if result.Error != nil{
        c.JSON(http.StatusBadRequest, gin.H{"error":result.Error})
    } else {
        c.JSON(http.StatusOK, gin.H{"data": "1"})
    }
}

// 查詢某會員參加的課程列表
func(st Apply) SearchMember(c *gin.Context){

    if err := c.ShouldBindJSON(&st); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    var result []ApplyMember


    db.MariaDB.Table("apply_lists").Select("lessons.lessonName, lessons.email, lessons.tuitionFee, apply_lists.sid, apply_lists.payed").Joins("left join lessons on lessons.id = apply_lists.lid").Where("apply_lists.sid = 1").Scan(&result)

    c.JSON(http.StatusOK, gin.H{"data": result})
}


// 查詢某課程的參加者
func(st Apply) SearchLesson(c *gin.Context){

    if err := c.ShouldBindJSON(&st); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    var result []ApplyMember

    db.MariaDB.Table("apply_lists").Select("students.studentName, students.email, apply_lists.*").Joins("left join students on students.id = apply_lists.lid").Where("apply_lists.Lid = ?", st.Lid).Scan(&result)
    c.JSON(http.StatusOK, gin.H{"data": result})
}


func(st Apply) Read(c *gin.Context){

    if err := c.ShouldBindJSON(&st); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    c.JSON(http.StatusBadRequest, gin.H{"error": "DB Failed"})
}
