package handler

import (
	"database/sql"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"CoursesBack/internal/models"
)

// @Summary SignUp
// @Tags auth
// @Description register
// @ID register
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param username formData string true "username"
// @Param password formData string true "password"
// @Param email formData string true "email"
// @Param birthday formData string true "birthday"
// @Param image formData file true "image"
// @Success 200 {string} string
// @Failure 404 {object} models.Error
// @Failure default {object} models.Error
// @Router /courses/auth/signup [post]
func (h *Handler) SignUp(c *gin.Context) {
	log := c.Request.Context().Value("logger").(*slog.Logger)

	var req models.SignUpRequest

	req.Username = c.PostForm("username")
	req.Password = c.PostForm("password")
	req.Email = c.PostForm("email")
	req.Birthday = c.PostForm("birthday")

	if !strings.Contains(req.Email, "@") {
		log.Error("incorrect email")
		c.JSON(http.StatusBadRequest, models.Error{
			Error: "Некорректная почта",
		})
		return
	}

	file, err := c.FormFile("image")
	if file != nil {
		if err != nil {
			log.Error("error with binding photo", "error", err.Error())
			c.JSON(http.StatusBadRequest, models.Error{
				Error: "Проблема с изображением",
			})
			return
		}

		photoData, err := readFile(file)

		if err != nil {
			log.Error("error with binding photo", "error", err.Error())
			c.JSON(http.StatusBadRequest, models.Error{
				Error: "Проблема с изображением",
			})
			return
		}

		req.Photo = photoData
	}

	result, err := h.service.SignUp(req)

	if err != nil {
		log.Error("error with inserting to db", "error", err.Error())
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: "Внутренняя ошибка сервера",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})

}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Success 200 {string} string
// @Failure 404 {object} models.Error
// @Failure default {object} models.Error
// @Router /courses/auth/signin [post]
func (h *Handler) SignIn(c *gin.Context) {

	var req models.SignUpRequest

	req.Password = c.PostForm("password")
	req.Email = c.PostForm("email")
	//todo тут через формы делаю хотя мб жсон
	isConfirmed, token, err := h.service.SignIn(req)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, models.Error{
				Error: "Пользователь не существует",
			})
			return
		}
		slog.Error("error with insertign result", "error", err.Error())
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: "Internal Server Error",
		})
		return
	}

	if !isConfirmed {
		c.JSON(http.StatusBadRequest, models.Error{
			Error: "Account is not confirmned",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// @Summary CheckKey
// @Tags auth
// @Description checkkey
// @ID checkkey
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param input body models.CheckMessage true "CheckKey"
// @Success 200 {string} string
// @Failure 404 {object} models.Error
// @Failure default {object} models.Error
// @Router /courses/auth/checkkey [post]
func (h *Handler) CheckKey(c *gin.Context) {
	var req models.CheckMessage

	err := c.ShouldBindJSON(&req)
	if err != nil {
		slog.Error("error with marshalling body", "error", err.Error())
		c.JSON(http.StatusBadRequest, models.Error{
			Error: "Incorrect Body Request",
		})
		return
	}

	ok, err := h.service.CheckKey(req.CheckKey, req.Mail)

	if err != nil {
		slog.Error("error with checking key", "error", err.Error())
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: "Internal Server Error",
		})
		return
	}

	if !ok {
		slog.Error("invalid key value")
		c.JSON(http.StatusBadRequest, models.Error{
			Error: "Incorrect Check Key",
		})
		return
	}

	c.JSON(http.StatusOK, "ok")
}

// @Summary GetUsedById
// @Tags userinfo
// @Description get user by id
// @ID get-user-by-id
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} models.User
// @Failure 404 {object} models.Error
// @Failure default {object} models.Error
// @Router /courses/api/userinfo/{id} [get]
func (h *Handler) GetUserById(c *gin.Context) {

	idString := c.Param("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		slog.Error("error with parsing id", "error", err.Error())
		c.JSON(http.StatusBadRequest, models.Error{
			Error: "Некорректное значение id",
		})
		return
	}

	user, err := h.service.GetUserById(id)

	if err != nil {
		slog.Error("error with getting user", "error", err.Error())
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary SetCreator
// @Security ApiKeyAuth
// @Tags personal
// @Description set creator
// @ID set creator
// @Accept json
// @Produce json
// @Success 200 {boolean} false
// @Failure 404 {object} models.Error
// @Failure default {object} models.Error
// @Router /courses/api/personal/creator [patch]
func (h *Handler) SetCreator(c *gin.Context) { //todo тут доделать и сделать какие-нибудь проверки

	userId, exists := c.Get("user")
	if !exists {
		slog.Error("error with getting userId")
		c.JSON(http.StatusBadRequest, models.Error{
			Error: "Invalid header",
		})
		return
	}

	setted, err := h.service.SetCreator(userId.(int))

	if err != nil {
		slog.Error("error with getting user", "error", err.Error())
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, setted)
}

// todo перенести в helper
func readFile(fileHeader *multipart.FileHeader) ([]byte, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return io.ReadAll(file)
}
