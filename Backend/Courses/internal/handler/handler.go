package handler

import (
	"CoursesBack/internal/service"
	"net/http"

	_ "CoursesBack/docs" // swagger embed files

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type Handler struct {
	service *service.Service
	salt    string
}

func NewHandler(service *service.Service, salt string) *Handler {
	return &Handler{service: service, salt: salt}
}

func (h *Handler) InitRoutes() http.Handler {
	controller := gin.Default()
	controller.RedirectTrailingSlash = false
	controller.Use(gin.Recovery(), Logger)

	auth := controller.Group("/auth")
	{
		auth.POST("/signup", h.SignUp)
		auth.POST("/signin", h.SignIn)
		auth.POST("/checkkey", h.CheckKey)
	}
	controller.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := controller.Group("/api")
	{
		userInfo := api.Group("/userinfo")
		{
			userInfo.GET("/:id", h.GetUserById)
		}

		personal := api.Group("/personal")
		personal.Use(h.CheckAuth)
		{
			personal.PATCH("/creator", h.SetCreator)
		}

		creatorCourse := api.Group("/creator")
		creatorCourse.Use(h.CheckAuthAndCreator)
		{
			creatorCourse.POST("/", h.CreateCourse)
			creatorCourse.GET("/", h.GetCreatorCourses)
			//todo сделать update(сейчс не помню) :))
			creatorCourse.DELETE("/:courseId", h.DeleteCourse)
		}

		courses := api.Group("/course")
		courses.Use(h.CheckAuth)
		{
			courses.POST("/:courseId/like", h.LikeCourse)
			courses.GET("/", h.GetCourses)
			courses.GET("/:courseId", h.GetFullCourse)
			courses.POST("/:courseId/subscribe", h.SubscribeCourse)
			lesson := courses.Group("/:courseId/lesson")
			{
				lesson.GET("/:lessonId", h.GetLesson)
				lesson.POST("/:lessonId/answer", h.AnswerTest)
			}

		}
	}

	return controller
}
