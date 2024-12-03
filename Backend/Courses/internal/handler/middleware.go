package handler

import (
	"CoursesBack/internal/models"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Logger(c *gin.Context) {
	logger := slog.Default()
	c.Request = c.Request.Clone(context.WithValue(c.Request.Context(), "logger", logger))
}

func (h *Handler) CheckAuth(c *gin.Context) {
	log := c.Request.Context().Value("logger").(*slog.Logger)
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		log.Error("missing token")
		c.AbortWithStatus(http.StatusUnauthorized)
		c.JSON(http.StatusBadRequest, models.Error{Error: "Missing token"})
		return
	}

	token, err := jwt.Parse(strings.Split(tokenString, " ")[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			log.Error("unexpected signing method")
			c.JSON(http.StatusBadRequest, models.Error{Error: "Unexpected signing method"})
		}
		return []byte(h.salt), nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		log.Error("error with token", "error", err)
		c.JSON(http.StatusBadRequest, models.Error{Error: "Error with token"})
		return
	}
	idFloat, ok := token.Claims.(jwt.MapClaims)["id"].(float64)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		log.Error("id is not a valid float64")
		c.JSON(http.StatusBadRequest, models.Error{Error: "Invalid id format"})
		return
	}
	id := int(idFloat)
	isConfirmed, err := h.service.GetIsConfirmed(id)

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		log.Error("failed to check confirmation")
		c.JSON(http.StatusInternalServerError, models.Error{Error: "Internal Server Error"})
		return
	}

	if !isConfirmed {
		c.AbortWithStatus(http.StatusUnauthorized)
		log.Error("failed to check confirmation")
		c.JSON(http.StatusInternalServerError, models.Error{Error: "Your account is not confirmed"})
		return
	}
	c.Set("user", id)
	log.Info("user with id signed", "id", token.Claims.(jwt.MapClaims)["id"])
}

func (h *Handler) CheckAuthAndCreator(c *gin.Context) {
	log := c.Request.Context().Value("logger").(*slog.Logger)
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		log.Error("missing token")
		c.AbortWithStatus(http.StatusUnauthorized)
		c.JSON(http.StatusBadRequest, models.Error{Error: "Missing token"})
		return
	}

	token, err := jwt.Parse(strings.Split(tokenString, " ")[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Error("unexpected signing method")
			c.AbortWithStatus(http.StatusUnauthorized)
			c.JSON(http.StatusBadRequest, models.Error{Error: "Unexpected signing method"})
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(h.salt), nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		log.Error("error with token", "error", err)
		c.JSON(http.StatusBadRequest, models.Error{Error: "Error with token"})
		return
	}

	idFloat, ok := token.Claims.(jwt.MapClaims)["id"].(float64)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		log.Error("id is not a valid float64")
		c.JSON(http.StatusBadRequest, models.Error{Error: "Invalid id format"})
		return
	}
	id := int(idFloat)
	isConfirmed, err := h.service.GetIsConfirmed(id)

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		log.Error("failed to check confirmation")
		c.JSON(http.StatusInternalServerError, models.Error{Error: "Internal Server Error"})
		return
	}

	if !isConfirmed {
		c.AbortWithStatus(http.StatusUnauthorized)
		log.Error("failed to check confirmation")
		c.JSON(http.StatusInternalServerError, models.Error{Error: "Your account is not confirmed"})
		return
	}
	isCreator, err := h.service.CheckCreator(id)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			c.AbortWithStatus(http.StatusUnauthorized)
			c.JSON(http.StatusBadRequest, models.Error{Error: "User not found"})
			return
		}
		c.AbortWithStatus(http.StatusUnauthorized)
		log.Error("error with checking creator role", "error", err)
		c.JSON(http.StatusBadRequest, models.Error{Error: "Error with checking creator role"})
		return
	}

	if !isCreator {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.JSON(http.StatusBadRequest, models.Error{Error: "You are not creator"})
		return
	}

	c.Set("user", id)
}
