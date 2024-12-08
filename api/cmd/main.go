package main

import (
	"fmt"

	"github.com/SergeyBogomolovv/milutin-jewelry/internal/config"
)

func main() {
	cfg := config.New()

	fmt.Println(cfg)
}
