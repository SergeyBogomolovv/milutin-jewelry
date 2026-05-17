package config

import (
	"time"

	"github.com/SergeyBogomolovv/milutin-jewelry/pkg/lib/env"
)

type Config struct {
	Port          int
	PostgresUrl   string
	RedisUrl      string
	CORSOrigin    string
	Mail          MailConfig
	Admin         AdminConfig
	ObjectStorage ObjectStorageConfig
	Jwt           JwtConfig
}

type MailConfig struct {
	Host  string
	Port  int
	User  string
	Pass  string
	Admin string
}

type AdminConfig struct {
	Email    string
	Password string
}

type ObjectStorageConfig struct {
	AccessKey string
	SecretKey string
	Endpoint  string
	Bucket    string
	Region    string
}

type JwtConfig struct {
	TTL    time.Duration
	Secret []byte
}

func New() Config {
	adminEmail := env.MustString("ADMIN_EMAIL")

	return Config{
		Port:        env.MustInt("PORT"),
		PostgresUrl: env.MustString("POSTGRES_URL"),
		RedisUrl:    env.MustString("REDIS_URL"),
		CORSOrigin:  env.MustString("CORS_ORIGIN"),
		Mail: MailConfig{
			Host:  env.String("MAIL_HOST", ""),
			Port:  env.Int("MAIL_PORT", 0),
			User:  env.String("MAIL_USER", ""),
			Pass:  env.String("MAIL_PASS", ""),
			Admin: adminEmail,
		},
		Admin: AdminConfig{
			Email:    adminEmail,
			Password: env.MustString("ADMIN_PASSWORD"),
		},
		ObjectStorage: ObjectStorageConfig{
			AccessKey: env.MustString("OBJECT_STORAGE_ACCESS"),
			SecretKey: env.MustString("OBJECT_STORAGE_SECRET"),
			Endpoint:  env.MustString("OBJECT_STORAGE_ENDPOINT"),
			Bucket:    env.MustString("OBJECT_STORAGE_BUCKET"),
			Region:    env.MustString("OBJECT_STORAGE_REGION"),
		},
		Jwt: JwtConfig{
			TTL:    env.MustDuration("JWT_TTL"),
			Secret: []byte(env.MustString("JWT_SECRET")),
		},
	}
}
