package kafka

import (
	"Notifiation/internal/config"
	"Notifiation/internal/models"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/jmoiron/sqlx"
	"github.com/segmentio/kafka-go"
)

type KafkaReader struct {
	Conn       *kafka.Reader
	DB         *sqlx.DB
	MailConfig config.MailAccount
}

func NewKafkaReader(kafkaConfig config.KafkaConfig, db *sqlx.DB, mailConfig config.MailAccount) *KafkaReader {
	conn := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{fmt.Sprintf("%s:%s", kafkaConfig.HostName, kafkaConfig.Port)},
		Topic:     kafkaConfig.Topic,
		Partition: kafkaConfig.Partition,
		MaxBytes:  10e6,
		GroupID:   "group",
	})
	return &KafkaReader{
		Conn:       conn,
		DB:         db,
		MailConfig: mailConfig,
	}
}

func (r *KafkaReader) CheckKafka() {
	var kafkaMessage models.KafkaMessage

	for {
		message, err := r.Conn.ReadMessage(context.Background())
		if err != nil {
			slog.Error("error with getting kafka message", "error", err.Error())

		}

		jsonDecoder := json.NewDecoder(bytes.NewReader(message.Value))

		err = jsonDecoder.Decode(&kafkaMessage)

		if err != nil {
			slog.Error("error with marshalling kafka message to struct", "error", err.Error())
		}

		err = r.processMessage(kafkaMessage)
		if err != nil {
			slog.Error("error with processing kafka message", "error", err.Error())
		}

		// Коммитим offset вручную после успешной обработки
		err = r.Conn.CommitMessages(context.Background(), message)
		if err != nil {
			slog.Error("Error committing offset:", "error", err.Error())
		} else {
			slog.Info("Offset committed successfully!")
		}
	}
}

func (r *KafkaReader) processMessage(message models.KafkaMessage) error {
	switch message.MessageType {
	case "mail":
		var mail models.MailMessage
		println()
		jsonDecoder := json.NewDecoder(bytes.NewReader([]byte(message.Message)))

		err := jsonDecoder.Decode(&mail)

		if err != nil {
			slog.Error("error with marshalling mail message", "error", err.Error())
			return err
		}

		err = r.SendMail(mail.Mail, mail.CheckInt)

		if err != nil {
			slog.Error("error with sending check", "error", err.Error())
			return err
		}
		return nil

	case "subscription":
		println(2)
	}
	return nil
}
