package handler

import (
	"CoursesBack/internal/models"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary GetLesson
// @Security ApiKeyAuth
// @Tags lessons
// @Description get lesson by id
// @ID get lesson by id
// @Accept json
// @Produce json
// @Param lessonId path int true "lessonId"
// @Param courseId path int true "courseId"
// @Success 200 {object} models.LessonFull
// @Failure default {object} models.Error
// @Router /courses/api/course/{courseId}/lesson/{lessonId} [get]
func (h *Handler) GetLesson(c *gin.Context) {
	lessonIdString := c.Param("lessonId")

	lessonId, err := strconv.Atoi(lessonIdString)
	if err != nil {
		slog.Error("error with parsing lessonId", "error", err.Error())
		c.JSON(http.StatusBadRequest, models.Error{
			Error: "Incorrect lesson id",
		})
		return
	}

	userId := c.GetInt("user")

	lesson, err := h.service.GetLesson(lessonId, userId)

	if err != nil {
		slog.Error("error with getting lesson", "error", err.Error())
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, lesson)
}

// @Summary AnswerTest
// @Security ApiKeyAuth
// @Tags lessons
// @Description answer test in lesson
// @ID answer test in lesson
// @Accept json
// @Produce json
// @Param lessonId path int true "lessonId"
// @Param courseId path int true "courseId"
// @Param answer body models.Answer true "answer"
// @Success 200 {object} models.ResultOfAnswering
// @Failure default {object} models.Error
// @Router /courses/api/course/{courseId}/lesson/{lessonId}/answer [post]
func (h *Handler) AnswerTest(c *gin.Context) {
	lessonIdString := c.Param("lessonId")

	lessonId, err := strconv.Atoi(lessonIdString)
	if err != nil {
		slog.Error("error with parsing lessonId", "error", err.Error())
		c.JSON(http.StatusBadRequest, models.Error{
			Error: "Incorrect lesson id",
		})

		return
	}

	userId := c.GetInt("user")

	var answer models.Answer

	err = c.ShouldBindJSON(&answer)

	if err != nil {
		slog.Error("error with parsing answer", "error", err.Error())
		c.JSON(http.StatusBadRequest, models.Error{
			Error: "Incorrect answer body",
		})
		return
	}

	result, err := h.service.AnswerTest(lessonId, userId, answer.Answer)

	if err != nil {
		slog.Error("error with answering test", "error", err.Error())
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
