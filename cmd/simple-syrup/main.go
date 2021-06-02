package main

import (
	"os"

	"github.com/Zferg/simple-http/pkg/web"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	web.Serve(port)
}
