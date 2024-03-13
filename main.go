package main

import (
	"database/sql"
	"log"
	"os"
	"video_search/controllers"
	"video_search/database"
	job "video_search/jobs"
	"video_search/routers"
	"video_search/service"

	"video_search/model"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {

	err := database.InitMongoClient()
	if err != nil {
		log.Print("main Error in initiating mongo err: ", err)
		return
	}

	// Start background task to fetch videos
	job.CronJobSetup()

	r := routers.SetupRouter()

	detailsProvider := BuildDetialsProvider()
	controllers.InitDetialsProvider(detailsProvider)

	// Start HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server listening on port %s", port)
	r.Run(":" + port)

}

func BuildDetialsProvider() model.DetialsProvider {
	return service.InitDetails()
}
