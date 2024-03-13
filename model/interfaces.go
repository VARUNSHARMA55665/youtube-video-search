package model

import (
	apihelper "video_search/apiHelper"
)

type DetialsProvider interface {
	GetStoredVideoDetails(page int, pageSize int) (int, apihelper.APIRes)
	SearchQueryBasedVideo(req SearchQueryBasedVideoReq) (int, apihelper.APIRes)
}
