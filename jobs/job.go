package job

import (
	"log"
	"video_search/youtube"

	"github.com/robfig/cron"
)

func CronJobSetup() {
	c := cron.New()

	// Schedule the to run job every half an hour
	err := c.AddFunc("0 */30 * * * *", youtube.WrapFetchVideos)
	if err != nil {
		log.Print("CronJobSetup Error scheduling job:", err)
	}

	// Start the Cron job
	c.Start()
}
