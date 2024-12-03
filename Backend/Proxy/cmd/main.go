package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

type Route struct {
	Prefix string
	Target *url.URL
	Proxy  *httputil.ReverseProxy
}

// NewReverseProxy создает новый ReverseProxy для заданного URL
func NewReverseProxy(target *url.URL) *httputil.ReverseProxy {
	proxy := httputil.NewSingleHostReverseProxy(target)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		// Здесь можно модифицировать запрос при необходимости
	}

	return proxy
}

func main() {
	// Получаем адреса бэкенд-сервисов из переменных окружения
	notificationURLStr := os.Getenv("NOTIFICATION_SERVICE_URL")
	if notificationURLStr == "" {
		notificationURLStr = "http://localhost:8081" // Протокол добавлен
	}
	coursesURLStr := os.Getenv("COURSES_SERVICE_URL")
	if coursesURLStr == "" {
		coursesURLStr = "http://localhost:8080/" // Протокол добавлен
	}
	authURLStr := os.Getenv("USERS_SERVICE_URL")
	if authURLStr == "" {
		authURLStr = "http://localhost:8082"
	}

	// Парсим URL-адреса
	notificationURL, err := url.Parse(notificationURLStr)
	if err != nil {
		log.Fatalf("Ошибка при разборе URL Notification: %v", err)
	}

	coursesURL, err := url.Parse(coursesURLStr)
	if err != nil {
		log.Fatalf("Ошибка при разборе URL Courses: %v", err)
	}

	authURL, err := url.Parse(authURLStr)
	if err != nil {
		log.Fatalf("Ошибка при разборе URL Auth: %v", err)
	}

	// Создаем маршруты
	routes := []Route{
		{
			Prefix: "/notification/",
			Target: notificationURL,
			Proxy:  NewReverseProxy(notificationURL),
		},
		{
			Prefix: "/courses/",
			Target: coursesURL,
			Proxy:  NewReverseProxy(coursesURL),
		},
		{
			Prefix: "/auth/",
			Target: authURL,
			Proxy:  NewReverseProxy(authURL),
		},
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		originalPath := r.URL.Path
		log.Printf("Запрос: %s", originalPath)

		// Попытка сопоставить исходный путь и путь с удаленным слэшем
		for _, route := range routes {
			if strings.HasPrefix(originalPath, route.Prefix) {
				log.Printf("Перенаправление %s на %s", originalPath, route.Target)
				r.URL.Path = strings.TrimPrefix(originalPath, route.Prefix)
				route.Proxy.ServeHTTP(w, r)
				return
			}
		}

		// Если путь не соответствует ни одному префиксу
		log.Printf("Маршрут не найден для %s", originalPath)
		http.NotFound(w, r)
	})

	// Запускаем HTTP-сервер
	serverAddress := ":8080"
	log.Printf("Прокси-сервер запущен на %s", serverAddress)
	if err := http.ListenAndServe(serverAddress, handler); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
