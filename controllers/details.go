package controllers

import (
	"log"
	"strconv"
	"video_search/model"

	"github.com/gin-gonic/gin"

	apihelper "video_search/apiHelper"
	"video_search/helpers"
)

var theDetialsProvider model.DetialsProvider

func InitDetialsProvider(provider model.DetialsProvider) {
	theDetialsProvider = provider
}

func GetStoredVideoDetails(c *gin.Context) {
	pageNumberStr := c.Query("pageNumber")
	pageSizeStr := c.Query("pageSize")

	if pageNumberStr == "" {
		pageNumberStr = "1"
	}
	if pageSizeStr == "" {
		pageSizeStr = "10"
	}

	pageNumber, err := strconv.Atoi(pageNumberStr)
	if err != nil || pageNumber <= 0 {
		pageNumber = 1 // Default to first page if page number is invalid
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 10 // default
	}

	log.Print("GetStoredVideoDetails Controller request pageNumber:", pageNumber, " pageSize: ", pageSize)

	code, resp := theDetialsProvider.GetStoredVideoDetails(pageNumber, pageSize)
	apihelper.CustomResponse(c, code, resp)
}

func SearchQueryBasedVideo(c *gin.Context) {
	title := c.Query("title")
	description := c.Query("description")

	var req model.SearchQueryBasedVideoReq
	req.Title = title
	req.Description = description

	log.Print("SearchQueryBasedVideo Controller request: ", helpers.LogStructAsJSON(req))

	code, resp := theDetialsProvider.SearchQueryBasedVideo(req)
	apihelper.CustomResponse(c, code, resp)
}
