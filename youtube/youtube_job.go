package youtube

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	apihelpers "video_search/apiHelper"
	"video_search/constants"
	"video_search/helpers"
	"video_search/model"
)

var nextPageToken string

func WrapFetchVideos() {

	log.Print("WrapFetchVideos Started Dumper")

	nextPageToken = ""
	for i := 0; i < constants.FetchPageCount; i++ {
		nextPageToken = FetchVideos()
	}

}

func FetchVideos() string {

	urlBase := constants.YoutubeSearchUrl
	url := createCompleteUrl(urlBase)

	// //make empty payload
	payload := new(bytes.Buffer)

	//call api
	// var apiRes apihelpers.APIRes
	res, err := apihelpers.CallApi(http.MethodGet, url, payload)
	defer res.Body.Close()
	if err != nil {
		log.Println("FetchVideos CallApi error =", err)
		return ""
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	searchYoutubeErrorRes := model.SearchYoutubeErrorRes{}
	err = json.Unmarshal([]byte(string(body)), &searchYoutubeErrorRes)
	if err == nil && searchYoutubeErrorRes.Error.Code == 403 {
		log.Println("FetchVideos CallApi youtube search api gave error response err: ", err)
		return ""
	}

	var searchYoutubeRes model.SearchYoutubeRes
	json.Unmarshal([]byte(string(body)), &searchYoutubeRes)

	var allMongoStore model.AllMongoStore

	for i := 0; i < len(searchYoutubeRes.Items); i++ {
		var mongoStore model.MongoStore
		mongoStore.VideoID = searchYoutubeRes.Items[i].ID.VideoID
		mongoStore.Title = searchYoutubeRes.Items[i].Snippet.Title
		mongoStore.Description = searchYoutubeRes.Items[i].Snippet.Description
		mongoStore.PublishedAt = searchYoutubeRes.Items[i].Snippet.PublishedAt
		mongoStore.PublishTime = searchYoutubeRes.Items[i].Snippet.PublishTime
		mongoStore.Thumbnails = searchYoutubeRes.Items[i].Snippet.Thumbnails

		helpers.MongoDumper(mongoStore)

		allMongoStore.MongoStoreData = append(allMongoStore.MongoStoreData, mongoStore)
	}

	return searchYoutubeRes.NextPageToken

}

// pageToken
func createCompleteUrl(urlBase string) string {
	url := urlBase + "?part=" + constants.Part + "&maxResults=" + strconv.Itoa(constants.MaxResults) + "&order=" + constants.Order + "&q=" + constants.Topic + "&type=" + constants.ContentType + "&key=" + constants.ApiKey
	// log.Print("createCompleteUrl nextPageToken: ", nextPageToken)
	if nextPageToken != "" {
		url += "&pageToken=" + nextPageToken
	}
	return url
}
