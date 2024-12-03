package handler

import (
	"CoursesBack/internal/models"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary GetCourses
// @Security ApiKeyAuth
// @Tags courses
// @Description get courses by search
// @ID get courses by search
// @Accept json
// @Produce json
// @Param search query string false "search"
// @Success 200 {object} []models.Course
// @Failure default {object} models.Error
// @Router /courses/api/course/ [get]
func (h *Handler) GetCourses(c *gin.Context) {
	searchCourse := c.Query("search")
	courses, err := h.service.GetCourses(searchCourse)
	if err != nil {
		slog.Error("error with getting course", "error", err.Error())

		c.JSON(http.StatusInternalServerError, models.Error{
			Error: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, courses)
}

// @Summary GetFullCourse
// @Security ApiKeyAuth
// @Tags courses
// @Description get full course
// @ID get full course
// @Accept json
// @Produce json
// @Param courseId path int true "courseId"
// @Success 200 {object} models.Course
// @Failure default {object} models.Error
// @Router /courses/api/course/{courseId} [get]
func (h *Handler) GetFullCourse(c *gin.Context) {
	courseIdString := c.Param("courseId")

	courseId, err := strconv.Atoi(courseIdString)
	if err != nil {
		slog.Error("error with arsing courseId", "error", err.Error())
		c.JSON(http.StatusBadRequest, models.Error{
			Error: "Incorrect course id",
		})
		return
	}

	course, err := h.service.GetFullCourse(courseId)

	if err != nil {
		slog.Error("error with getting full course", "error", err.Error())
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, course)
}

// @Summary LikeCourse
// @Security ApiKeyAuth
// @Tags courses
// @Description like course or dislike
// @ID like course
// @Accept json
// @Produce json
// @Param courseId path int true "courseId"
// @Success 200 {bool} true
// @Failure default {object} models.Error
// @Router /courses/api/course/{courseId}/like [post]
func (h *Handler) LikeCourse(c *gin.Context) {
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
		slog.Error("error with arsing courseId", "error", err.Error())
		c.JSON(http.StatusBadRequest, models.Error{
			Error: "Incorrect course id",
		})
		return
	}

	ok, err := h.service.LikeCourse(userId.(int), courseId)

	if err != nil {
		slog.Error("error with like course", "error", err.Error())
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, ok)
}

// @Summary SubscribeCourse
// @Security ApiKeyAuth
// @Tags courses
// @Description subscribe course
// @ID subscribe course
// @Accept json
// @Produce json
// @Param courseId path int true "courseId"
// @Success 200 {bool} true
// @Failure default {object} models.Error
// @Router /courses/api/course/{courseId}/subscribe [post]
func (h *Handler) SubscribeCourse(c *gin.Context) {
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
		slog.Error("error with arsing courseId", "error", err.Error())
		c.JSON(http.StatusBadRequest, models.Error{
			Error: "Incorrect course id",
		})
		return
	}

	ok, err := h.service.SubscribeCourse(userId.(int), courseId)
	//todo закидывать в нотификашки создателю о подписке на курс
	if err != nil {
		slog.Error("error with subscribe course", "error", err.Error())
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, ok)
}

//todo протестить лайк
