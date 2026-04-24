package main

import (
	"fmt"
	"net/http"

	"github.com/larry-lw-chan/helloworld/pkg/config"
	"github.com/larry-lw-chan/helloworld/pkg/handlers"
	"github.com/larry-lw-chan/helloworld/pkg/render"
)

// Configuration
const PORT_NUMBER string = ":8080"

func main() {
	app := config.AppConfig{}

	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("Error creating template cache")
	}

	// Assign the template cache to the app
	app.TemplateCache = tc

	// Declare routes here
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	// Start the server
	fmt.Println("Server is running on port", PORT_NUMBER)
	http.ListenAndServe(PORT_NUMBER, nil)
}
