package config

import (
	"time"

	"github.com/SergeyBogomolovv/milutin-jewelry/pkg/lib/env"
)

type Config struct {
	Addr          string
	PostgresUrl   string
	RedisUrl      string
	CORSOrigin    string
	Mail          MailConfig
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
	return Config{
		Addr:        env.MustString("ADDR"),
		PostgresUrl: env.MustString("POSTGRES_URL"),
		RedisUrl:    env.MustString("REDIS_URL"),
		CORSOrigin:  env.MustString("CORS_ORIGIN"),
		Mail: MailConfig{
			Host:  env.MustString("MAIL_HOST"),
			Port:  env.MustInt("MAIL_PORT"),
			User:  env.MustString("MAIL_USER"),
			Pass:  env.MustString("MAIL_PASS"),
			Admin: env.MustString("ADMIN_EMAIL"),
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
