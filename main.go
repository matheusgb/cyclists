package main

import (
	"github.com/matheusgb/cyclists/src/config"
	"github.com/matheusgb/cyclists/src/layers"
)

func main() {
	config := config.Init()
	config.MountConfigs()
	app := layers.Setup(config)
	app.Listen(config.Api.Port)
}
