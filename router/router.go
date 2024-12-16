package router

import(
    "GoApp/student"
    "GoApp/teacher"
    "GoApp/lesson"
    "GoApp/applyList"
    "GoApp/login"
    "GoApp/middleware"
    "github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine {
	router := gin.New()

	studentRouter := router.Group("/student", middleware.CheckToken())
	{
        studentRouter.GET("/", student.All) //讀取所有學生
        studentRouter.GET("/:id", student.StudentInfo.Read) //讀取一位學生
		studentRouter.PUT("/", student.StudentInfo.Update) //更新學生資料
	}


	teacherRouter := router.Group("/teacher", middleware.CheckToken())
	{
        teacherRouter.GET("/", teacher.All)  //讀取所有教師
        teacherRouter.GET("/:id", teacher.TeacherInfo.Read) //讀取一位教師
		teacherRouter.PUT("/", teacher.TeacherInfo.Update) //更新教師資料
	}

	lessonRouter := router.Group("/lesson", middleware.CheckToken())
	{
        lessonRouter.GET("/", lesson.LessonInfo.All) //讀取所有課程
        lessonRouter.GET("/:id", lesson.LessonInfo.Read) //讀取某一課程
		lessonRouter.POST("/", lesson.LessonInfo.Create) //建立新課程
		lessonRouter.PUT("/", lesson.LessonInfo.Update)  //更新課程資料
	}

	applyRouter := router.Group("/apply", middleware.CheckToken())
	{
		applyRouter.POST("/", applyList.ApplyList.Create) //報名
		applyRouter.PUT("/", applyList.ApplyList.Update)  //付款
		applyRouter.GET("/searchMember", applyList.ApplyList.SearchMember) //查詢某會員參加的課程列表
		applyRouter.GET("/searchLesson", applyList.ApplyList.SearchLesson) // 查詢某課程的參加者

		applyRouter.GET("/", applyList.ApplyList.Read)
	}


    router.POST("/teacherRegister", teacher.TeacherInfo.Create) //註冊教師
    router.POST("/register", student.StudentInfo.Create) //註冊學生
//     router.GET("/login", login.LoginInfo.UnSignIn)

    router.POST("/teacherLogin", login.LoginInfo.TeacherSignIn) //老師登入
    router.POST("/login", login.LoginInfo.SignIn)  //學生登入

    return router
}
