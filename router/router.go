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
        studentRouter.GET("/", student.All)
        studentRouter.GET("/:id", student.StudentInfo.Read)
		studentRouter.PUT("/", student.StudentInfo.Update)
	}


	teacherRouter := router.Group("/teacher", middleware.CheckToken())
	{
        teacherRouter.GET("/", teacher.All)
        teacherRouter.GET("/:id", teacher.TeacherInfo.Read)
		teacherRouter.PUT("/", teacher.TeacherInfo.Update)
	}

	lessonRouter := router.Group("/lesson", middleware.CheckToken())
	{
        lessonRouter.GET("/", lesson.LessonInfo.All)
        lessonRouter.GET("/:id", lesson.LessonInfo.Read)
		lessonRouter.POST("/", lesson.LessonInfo.Create)
		lessonRouter.PUT("/", lesson.LessonInfo.Update)
	}

	applyRouter := router.Group("/apply", middleware.CheckToken())
	{
		applyRouter.POST("/", applyList.ApplyList.Create)
		applyRouter.PUT("/", applyList.ApplyList.Update)
		applyRouter.GET("/searchMember", applyList.ApplyList.SearchMember)
		applyRouter.GET("/searchLesson", applyList.ApplyList.SearchLesson)

		applyRouter.GET("/", applyList.ApplyList.Read)
	}


    router.POST("/teacherRegister", teacher.TeacherInfo.Create)
    router.POST("/register", student.StudentInfo.Create)
//     router.GET("/login", login.LoginInfo.UnSignIn)

    router.POST("/teacherLogin", login.LoginInfo.TeacherSignIn)
    router.POST("/login", login.LoginInfo.SignIn)

    return router
}
