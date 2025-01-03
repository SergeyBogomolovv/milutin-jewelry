package env

import (
	"log"
	"os"
	"strconv"
	"time"
)

func String(key, fallback string) string {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return fallback
}

func Int(key string, fallback int) int {
	val, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	v, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return v
}

func MustString(key string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("missing env var: %s", key)
	}
	return val
}

func MustInt(key string) int {
	val, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("missing env var: %s", key)
	}
	v, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("invalid env var: %s", key)
	}
	return v
}

func MustDuration(key string) time.Duration {
	str, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("missing env var: %s", key)
	}
	val, err := time.ParseDuration(str)
	if err != nil {
		log.Fatalf("invalid env var: %s", key)
	}
	return val
}
