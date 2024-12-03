package handler

import (
	"Notifiation/internal/models"
	"Notifiation/internal/store"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	repo *store.Store
	salt string
}

func NewHandler(repo *store.Store, salt string) *Handler {
	return &Handler{repo: repo, salt: salt}
}

func (h *Handler) InitRoutes() http.Handler {
	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/notifications", h.GetNotifications)
	return router
}

func (h *Handler) GetNotifications(c *gin.Context) {
	println(1)
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
	c.Set("user", id)
	log.Info("user with id signed", "id", token.Claims.(jwt.MapClaims)["id"])
}
