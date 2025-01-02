package config

import (
	"github.com/SergeyBogomolovv/milutin-jewelry/pkg/lib/env"
	"github.com/joho/godotenv"
)

type Config struct {
	Addr          string
	PostgresUrl   string
	RedisUrl      string
	JWTSecret     string
	AdminEmail    string
	Mail          MailConfig
	ObjectStorage ObjectStorageConfig
	CORSOrigin    string
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
		Addr:        env.GetEnvString("ADDR", ":8080"),
		PostgresUrl: env.GetEnvString("POSTGRES_URL", "postgres://admin:admin@localhost:5432/jewelry?sslmode=disable"),
		RedisUrl:    env.GetEnvString("REDIS_URL", "redis://localhost:6379/0"),
		JWTSecret:   env.GetEnvString("JWT_SECRET", "secret"),
		AdminEmail:  env.GetEnvString("ADMIN_EMAIL", "email@email.com"),
		CORSOrigin:  env.GetEnvString("CORS_ORIGIN", "http://localhost:3000"),
		Mail: MailConfig{
			Host: env.GetEnvString("MAIL_HOST", "smtp.gmail.com"),
			Port: env.GetEnvInt("MAIL_PORT", 587),
			User: env.GetEnvString("MAIL_USER", "example@example.com"),
			Pass: env.GetEnvString("MAIL_PASS", "password"),
		},
		ObjectStorage: ObjectStorageConfig{
			AccessKey: env.GetEnvString("OBJECT_STORAGE_ACCESS", "secret"),
			SecretKey: env.GetEnvString("OBJECT_STORAGE_SECRET", "secret"),
			Endpoint:  env.GetEnvString("OBJECT_STORAGE_ENDPOINT", "https://storage.yandexcloud.net"),
			Bucket:    env.GetEnvString("OBJECT_STORAGE_BUCKET", "milutin-jewelry"),
			Region:    env.GetEnvString("OBJECT_STORAGE_REGION", "ru-central1"),
		},
	}
}
