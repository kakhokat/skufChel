package main

import (
	"CoursesBack/internal/config"
	"CoursesBack/internal/handler"
	"CoursesBack/internal/migrate"
	"CoursesBack/internal/server"
	"CoursesBack/internal/service"
	"CoursesBack/internal/store"
	"context"
	"log/slog"
	"os"

	"github.com/segmentio/kafka-go"
)

// @title Courses API
// @description API Server 4 Courses

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	slog.SetDefault(log)

	baseConfig, err := config.CreateConfig()

	if err != nil {
		slog.Error("error with reading config", "error", err.Error())
		return
	}

	dbConfig := store.NewConfig(
		baseConfig.Database.Hostname,
		baseConfig.Database.Port,
		baseConfig.Database.Username,
		baseConfig.Database.Password,
		baseConfig.Database.Database,
		baseConfig.Database.Sslmode,
	)
	db, err := store.InitPostgres(dbConfig)
	if err != nil {
		slog.Error("Failed to initialize database", "error", err.Error())
		return
		//todo логи сделать другими, отдельный пакет, в котором будут логи, потом мб подрубить кибану, эластик че-нить такое
	}

	conn, err := kafka.DialLeader(context.Background(), "tcp", baseConfig.Kafka.HostName+":"+baseConfig.Kafka.Port, baseConfig.Kafka.Topic, baseConfig.Kafka.Partition)
	if err != nil {
		panic(err)
	}

	storeLevel := store.NewStore(db)
	serviceLevel := service.NewService(storeLevel, baseConfig.Salt, conn)
	handlerLevel := handler.NewHandler(serviceLevel, baseConfig.Salt)
	err = migrate.Migrat(baseConfig.Database.Hostname)
	if err != nil {
		slog.Error("Failed to migrate database", "error", err.Error())
	}
	baseServer := new(server.Server)

	if err := baseServer.InitServer("8080", handlerLevel.InitRoutes()); err != nil {
		log.Error("error with init server", "error", err.Error())
	}

	//todo graceful shutdown
}
