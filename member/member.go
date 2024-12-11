package member

import(
    "github.com/gin-gonic/gin"
    "net/http"
    "GoApp/pkg/hashing"
    db "GoApp/database"
)

type Member struct{
    Id int `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
    Account string `gorm:"column:account" json:"account"`
    FirstName string `gorm:"column:firstName" json:"firstName"`
    LastName string `gorm:"column:lastName" json:"lastName"`
    Email string `gorm:"column:email;index" json:"email"`
    Phone string `gorm:"column:phone" json:"phone"`
    PW string `gorm:"column:password" json:"password"`
    MType string `gorm:"column:memberType" json:"memberType"`
}

var MemberInfo Member

func(m Member) Create(c *gin.Context){

    if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    m.PW = hashing.HashPassword(m.PW)
    result := db.MariaDB.Create(&m)

    if result.Error != nil{
        c.JSON(http.StatusBadRequest, gin.H{"error":result.Error})
    } else {
        c.JSON(http.StatusOK, gin.H{"data": "1"})
    }
}

func(m Member) Update(c *gin.Context){

    if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"request error => ": err.Error()})
		return
	}

    m.PW = hashing.HashPassword(m.PW)
    result := db.MariaDB.Save(&m)

    if result.Error != nil{
        c.JSON(http.StatusBadRequest, gin.H{"database update error => ":result.Error})
    } else {
        c.JSON(http.StatusOK, gin.H{"data": m})
    }
}

func(m Member) Read(c *gin.Context){

    id := c.Query("id")

    result := db.MariaDB.First(&m, id)

    if result.Error != nil{
        c.JSON(http.StatusBadRequest, gin.H{"error":result.Error})
    } else {
        c.JSON(http.StatusOK, gin.H{"data": m})
    }

}

func All(c *gin.Context){

    var data []Member

    result := db.MariaDB.Find(&data)

    if result.Error != nil{
        c.JSON(http.StatusBadRequest, gin.H{"error":result.Error})
    } else if result.RowsAffected == 0 {
        c.JSON(http.StatusOK, gin.H{"data": "null"})
    } else {
        c.JSON(http.StatusOK, gin.H{"data": data})
    }

}


func FindOne(st Member)(Member, bool){

    reData := Member{}

    result := db.MariaDB.Where(&st).First(&reData)

    if result.RowsAffected > 0 {
        return reData, true
    }else{
        return reData, false
    }


}