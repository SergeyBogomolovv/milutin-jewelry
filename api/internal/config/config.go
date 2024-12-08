package config

import (
	"github.com/SergeyBogomolovv/milutin-jewelry/pkg/utils"
	"github.com/joho/godotenv"
)

type Config struct {
	Addr          string
	PostgresUrl   string
	RedisUrl      string
	JWTSecret     string
	Mail          MailConfig
	ObjectStorage ObjectStorageConfig
}

type MailConfig struct {
	Host string
	Port int
	User string
	Pass string
}

type ObjectStorageConfig struct {
	AccessKey string
	SecretKey string
	Endpoint  string
	Bucket    string
	Region    string
}

func New() *Config {
	godotenv.Load()

	return &Config{
		Addr:        utils.GetEnvString("ADDR", ":8080"),
		PostgresUrl: utils.GetEnvString("POSTGRES_URL", "postgres://admin:admin@localhost:5432/jewelry?sslmode=disable"),
		RedisUrl:    utils.GetEnvString("REDIS_URL", "redis://localhost:6379/0"),
		JWTSecret:   utils.GetEnvString("JWT_SECRET", "secret"),
		Mail: MailConfig{
			Host: utils.GetEnvString("MAIL_HOST", "smtp.gmail.com"),
			Port: utils.GetEnvInt("MAIL_PORT", 587),
			User: utils.GetEnvString("MAIL_USER", "example@example.com"),
			Pass: utils.GetEnvString("MAIL_PASS", "password"),
		},
		ObjectStorage: ObjectStorageConfig{
			AccessKey: utils.GetEnvString("OBJECT_STORAGE_ACCESS", "secret"),
			SecretKey: utils.GetEnvString("OBJECT_STORAGE_SECRET", "secret"),
			Endpoint:  utils.GetEnvString("OBJECT_STORAGE_ENDPOINT", "https://storage.yandexcloud.net"),
			Bucket:    utils.GetEnvString("OBJECT_STORAGE_BUCKET", "milutin-jewelry"),
			Region:    utils.GetEnvString("OBJECT_STORAGE_REGION", "ru-central1"),
		},
	}
}
