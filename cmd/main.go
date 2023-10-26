package main

import (
	"codebin/internal/app"
	_ "codebin/swagger_ui"
)

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
