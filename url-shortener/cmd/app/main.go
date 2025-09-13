package main

import "url-shortener/internal/app"

const configPath = "./configs/local.yml"

func main() {
	app.Run(configPath)
}
