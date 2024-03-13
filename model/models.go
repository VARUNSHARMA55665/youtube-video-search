package model

import "time"

type SearchYoutubeRes struct {
	Kind          string `json:"kind"`
	Etag          string `json:"etag"`
	NextPageToken string `json:"nextPageToken"`
	RegionCode    string `json:"regionCode"`
	PageInfo      struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
	Items []struct {
		Kind string `json:"kind"`
		Etag string `json:"etag"`
		ID   struct {
			Kind    string `json:"kind"`
			VideoID string `json:"videoId"`
		} `json:"id"`
		Snippet struct {
			PublishedAt time.Time `json:"publishedAt"`
			ChannelID   string    `json:"channelId"`
			Title       string    `json:"title"`
			Description string    `json:"description"`
			Thumbnails  struct {
				Default struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"default"`
				Medium struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"medium"`
				High struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"high"`
			} `json:"thumbnails"`
			ChannelTitle         string    `json:"channelTitle"`
			LiveBroadcastContent string    `json:"liveBroadcastContent"`
			PublishTime          time.Time `json:"publishTime"`
		} `json:"snippet"`
	} `json:"items"`
}

type AllMongoStore struct {
	MongoStoreData []MongoStore `json:"mongoStore"`
}

type MongoStore struct {
	VideoID     string    `json:"videoId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"publishedAt"`
	PublishTime time.Time `json:"publishTime"`
	Thumbnails  struct {
		Default struct {
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"default"`
		Medium struct {
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"medium"`
		High struct {
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"high"`
	} `json:"thumbnails"`
}

type SearchYoutubeErrorRes struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Errors  []struct {
			Message      string `json:"message"`
			Domain       string `json:"domain"`
			Reason       string `json:"reason"`
			Location     string `json:"location"`
			LocationType string `json:"locationType"`
		} `json:"errors"`
		Status string `json:"status"`
	} `json:"error"`
}

type SearchQueryBasedVideoReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
