package handler

import (
	"CoursesBack/internal/models"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
) //todo документацию поменять немного

// @Summary CreateCourse
// @Security ApiKeyAuth
// @Tags creator
// @Description create course
// @ID create course
// @Accept json
// @Produce json
// @Param input body models.Course true "Course"
// @Success 200 {string} string
// @Failure 404 {object} models.Error
// @Failure default {object} models.Error
// @Router /courses/api/creator/ [post]
func (h *Handler) CreateCourse(c *gin.Context) {
	log := c.Request.Context().Value("logger").(*slog.Logger) //todo убрать, slog сам задает дефолтный логгер
	var request *models.Course

	userId, exists := c.Get("user")

	if !exists {
		log.Error("error with getting userId")
		c.JSON(http.StatusBadRequest, models.Error{
			Error: "Invalid header",
		})
		return
	}

	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Error("error with binding", "error", err.Error())
		c.JSON(http.StatusBadRequest, models.Error{
			Error: "Incorrect Body Request",
		})
		return
	}

	id, err := h.service.CreateCourse(request, userId.(int))
	if err != nil {
		log.Error("error with insertign result", "error", err.Error())
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

// @Summary GetCreatorCourses
// @Security ApiKeyAuth
// @Tags creator
// @Description get creator courses
// @ID get creator courses
// @Accept json
// @Produce json
// @Param search query string false "search"
// @Success 200 {object} []models.Course
// @Failure 404 {object} models.Error
// @Failure default {object} models.Error
// @Router /courses/api/creator/ [get]
func (h *Handler) GetCreatorCourses(c *gin.Context) {
	userId, exists := c.Get("user")
	if !exists {
		slog.Error("error with getting userId")
		c.JSON(http.StatusBadRequest, models.Error{
			Error: "Invalid header",
		})
		return
	}

	searchCourse := c.Query("search")

	courses, err := h.service.GetCreatorCourses(userId.(int), searchCourse)

	if err != nil {
		slog.Error("error with getting course", "error", err.Error())
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, courses)
}

// @Summary DeleteCourse
// @Security ApiKeyAuth
// @Tags creator
// @Description delete course
// @ID delete course
// @Accept json
// @Produce json
// @Param input query string false "search"
// @Success 200 {string} string
// @Failure 404 {object} models.Error
// @Failure default {object} models.Error
// @Router /courses/api/creator/ [delete]
func (h *Handler) DeleteCourse(c *gin.Context) {
	userId, exists := c.Get("user")
	if !exists {
		slog.Error("error with getting userId")
		c.JSON(http.StatusBadRequest, models.Error{
			Error: "Invalid header",
		})
		return
	}

	courseIdString := c.Param("courseId")

	courseId, err := strconv.Atoi(courseIdString)

	if err != nil {
		slog.Error("error with parsing courseId", "error", err.Error())
		c.JSON(http.StatusBadRequest, models.Error{
			Error: "Invalid courseId",
		})
		return
	}

	ok, err := h.service.DeleteCourse(userId.(int), courseId)

	if err != nil {
		slog.Error("error with deleting", "error", err.Error())
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": ok,
	})
}
