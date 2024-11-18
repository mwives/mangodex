package main

import (
	"github.com/mwives/mangodex/internal/app"
	"github.com/mwives/mangodex/internal/app/config"
)

func main() {
	config.InitConfig()
	app.Run()
}
