package main

import (
	"byteknot-api/configs"
	"byteknot-api/internal/db"
	"byteknot-api/internal/db/repository"
	"byteknot-api/internal/handler"
	"byteknot-api/internal/routes"
	"byteknot-api/internal/service"
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Starting application setup...")

	// 1. Load Config
	log.Println("Loading Config from .env file...")
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// 2. DB Connection
	dbConn, err := db.NewMySQLConnection(config)
	if err != nil {
		log.Fatalf("Database connection failed: %v\n", err)
	}
	defer dbConn.Close()

	// 3. Inject Dependency setup layers
	articleRepo := repository.NewArticleRepository(dbConn)
	commentsRepo := repository.NewCommentsRepository(dbConn)

	articleService := service.NewArticleService(articleRepo)
	commentsService := service.NewCommentsService(commentsRepo)

	articleHandler := handler.NewArticleHandler(articleService)
	commentsHandler := handler.NewCommentsHandler(commentsService)

	// 4. Setup Router
	r := routes.ByteKnotRouter(articleHandler, commentsHandler)
	http.Handle("/", r)
	log.Println("Gorilla/mux: routes registered")

	// 5. Start Server
	listenAddr := fmt.Sprintf(":%d", config.Server.Port)
	log.Printf("Server listening on http://localhost%s", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
