package constants

const (
	Part             = "snippet"
	MaxResults       = 50
	Order            = "date"
	Topic            = "cricket"
	ContentType      = "video"
	ApiKey           = "AIzaSyC5zCddD2PXN8T0r8hfeZx3dsShsg4whBE"
	FetchPageCount   = 10
	YoutubeSearchUrl = "https://youtube.googleapis.com/youtube/v3/search"
)

const (
	MongoBase = "mongodb+srv://varun"
	MongoPass = "vmv2021"
	MongoUri  = "cluster0.uhwsn.mongodb.net/team_crew_db?retryWrites=true&w=majority"
)

// Collections name
const (
	YoutubeDbName = "youtube-store"
	VideoDetails  = "video-detail"
)

// Mongo Error message
const (
	MongoNoDocError     = "mongo: no documents in result"
	InternalServerError = "Internal Server Error"
)
