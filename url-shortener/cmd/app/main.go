package main

import "url-shortener/internal/app"

const configsDir = "configs"

func main() {
	app.Run(configsDir)
}
