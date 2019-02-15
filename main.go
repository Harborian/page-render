package main

import (
	"log"

	chromeRenderer "github.com/harborian/page-render/renderer/chrome"
)

func main() {
	r := chromeRenderer.New()

	_, err := r.Render("https://google.com")
	if err != nil {
		log.Printf("error - %s", err)
	}
}
