package server

import (
	"fmt"
	"log"
	"os"
)

func Start() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := run(port); err != nil {
		log.Printf("Error: %v", err)
	}

	fmt.Println("Server is running on port " + port)
	// http.ListenAndServe(":"+port, nil)
}

func run(port string) error {

	router := mapRoutes()
	return router.Run(":" + port)
}
