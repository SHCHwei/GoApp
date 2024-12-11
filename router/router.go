package router

import(
    "GoApp/member"
    "GoApp/lesson"
    "GoApp/applyList"
    "GoApp/login"
    "GoApp/middleware"
    "github.com/gin-gonic/gin"

)


func SetupRouter() *gin.Engine {
	router := gin.New()

	memberRouter := router.Group("/member", middleware.CheckToken())
	{
        memberRouter.GET("/", member.All)
        memberRouter.GET("/:id", member.MemberInfo.Read)
		memberRouter.PUT("/", member.MemberInfo.Update)
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



    router.POST("/register", member.MemberInfo.Create)
    router.GET("/login", login.LoginInfo.UnSignIn)
    router.POST("/login", login.LoginInfo.SignIn)



    return router
}
