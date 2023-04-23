package main

import (
	"fmt"
	"log"

	"github.com/soulteary/nginx-formatter/internal/cmd"
	"github.com/soulteary/nginx-formatter/internal/formatter"
	"github.com/soulteary/nginx-formatter/internal/updater"
	"github.com/soulteary/nginx-formatter/internal/version"
)

func main() {
	fmt.Printf("Nginx Formatter v%s\n\n", version.Version)

	src, dest, indent, char := cmd.InitArgv()
	err := updater.UpdateConfInDir(src, dest, indent, char, formatter.Formatter)
	if err != nil {
		log.Fatal(err)
	}
}
