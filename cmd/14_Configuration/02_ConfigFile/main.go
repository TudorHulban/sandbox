package main

import (
	"embed"
	"os"
)

//go:embed configs/*.json
var _fs embed.FS

func main() {
	config, _ := NewConfigurationFromFS("configs/config.json")

	config.WriteTo(os.Stdout)
}
