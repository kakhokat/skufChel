package main

import (
	"Notifiation/internal/config"
	"Notifiation/internal/handler"
	"Notifiation/internal/kafka"
	"Notifiation/internal/migrate"
	"Notifiation/internal/server"
	"Notifiation/internal/store"
	"log/slog"
	"os"
)

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
	if err != nil {
		slog.Error("error with connection to kafka", "error", err.Error())
	}

	db, err := store.InitPostgres(dbConfig)
	if err != nil {
		slog.Error("Failed to initialize database", "error", err.Error())
		return
		//todo логи сделать другими, отдельный пакет, в котором будут логи, потом мб подрубить кибану, эластик че-нить такое
	}

	kafkaReader := kafka.NewKafkaReader(baseConfig.Kafka, db, baseConfig.Mail)

	go kafkaReader.CheckKafka()

	storeLevel := store.NewStore(db)
	handlerLevel := handler.NewHandler(storeLevel, baseConfig.Salt)
	err = migrate.Migrat(baseConfig.Database.Hostname)
	if err != nil {
		slog.Error("Failed to migrate database", "error", err.Error())
	}
	baseServer := new(server.Server)

	if err := baseServer.InitServer(baseConfig.HttpServer.Port, handlerLevel.InitRoutes()); err != nil {
		log.Error("error with init server", "error", err.Error())
	}

}
