package main

import (
	"github.com/matheusgb/cyclists/src/config"
	"github.com/matheusgb/cyclists/src/layers"
)

// @title						Cyclists API
// @version					1.0
// @description				API REST em Go para gestão de eventos de ciclismo: usuários, inscrições e autenticação com JWT.
// @license.name				MIT
// @license.url				https://github.com/matheusgb/cyclists/blob/main/LICENSE
// @BasePath					/api/v1
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
// @description				Informe o token JWT no formato: Bearer {token}
func main() {
	config := config.Init()
	config.MountConfigs()
	app := layers.Setup(config)
	app.Listen(config.Api.Port)
}
