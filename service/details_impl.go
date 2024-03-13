package service

import (
	"context"
	"log"
	"net/http"
	apihelper "video_search/apiHelper"
	"video_search/constants"
	"video_search/database"
	"video_search/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DetailsObj struct{}

func InitDetails() DetailsObj {
	detailsObj := DetailsObj{}

	return detailsObj
}

func (obj DetailsObj) GetStoredVideoDetails(pageNumber int, pageSize int) (int, apihelper.APIRes) {
	var apiRes apihelper.APIRes

	var allMongoStore model.AllMongoStore

	// Calculate the number of documents to skip based on the page number and page size
	skip := (pageNumber - 1) * pageSize

	findOptions := options.Find()
	// Sort by `publishedat` field descending
	findOptions.SetSort(bson.D{{"publishedat", -1}})

	// Set the number of documents to skip
	findOptions.SetSkip(int64(skip))

	// Set the number of documents to return for this page
	findOptions.SetLimit(int64(pageSize))

	allDoc, err := database.FindAllMongo(constants.VideoDetails, bson.M{}, findOptions)
	if err != nil && err.Error() != constants.MongoNoDocError {
		log.Printf("GetStoredVideoDetails Error in mongo err: ", err)
		return apihelper.SendInternalServerError()
	}

	for allDoc.Next(context.Background()) {
		var mongoStore model.MongoStore
		err := allDoc.Decode(&mongoStore)
		if err != nil {
			log.Printf("GetStoredVideoDetails Mongo parsing failed error =", err)
		}
		allMongoStore.MongoStoreData = append(allMongoStore.MongoStoreData, mongoStore)
	}

	apiRes.Data = allMongoStore
	apiRes.Message = "SUCCESS"
	apiRes.Status = true
	return http.StatusOK, apiRes
}

func (obj DetailsObj) SearchQueryBasedVideo(req model.SearchQueryBasedVideoReq) (int, apihelper.APIRes) {
	var apiRes apihelper.APIRes

	var combineMongoStore []model.MongoStore

	// Check if title or description is empty
	if req.Title == "" && req.Description == "" {
		apiRes.Message = "Both title and description are empty"
		apiRes.Status = false
		return http.StatusBadRequest, apiRes
	}

	// Construct regular expression patterns for title and description if they are not empty
	var titleRegex, descriptionRegex bson.M
	if req.Title != "" {
		titleRegex = bson.M{"$regex": primitive.Regex{Pattern: req.Title, Options: "i"}}
	}
	if req.Description != "" {
		descriptionRegex = bson.M{"$regex": primitive.Regex{Pattern: req.Description, Options: "i"}}
	}

	// Construct the query using $or operator to match either title or description
	var query bson.M
	if req.Title != "" && req.Description != "" {
		query = bson.M{"$or": []bson.M{{"title": titleRegex}, {"description": descriptionRegex}}}
	} else if req.Title != "" {
		query = bson.M{"title": titleRegex}
	} else if req.Description != "" {
		query = bson.M{"description": descriptionRegex}
	}

	findOptions := options.Find()
	// Sort by `publishedat` field descending
	findOptions.SetSort(bson.D{{"publishedat", -1}})

	// Perform the search
	cursor, err := database.FindAllMongo(constants.VideoDetails, query, findOptions)
	if err != nil && err.Error() != constants.MongoNoDocError {
		log.Printf("SearchQueryBasedVideo Error in mongo err: ", err)
		return apihelper.SendInternalServerError()
	}

	// Iterate over the results and decode them
	for cursor.Next(context.Background()) {
		var result model.MongoStore
		err := cursor.Decode(&result)
		if err != nil {
			log.Printf("SearchQueryBasedVideo Mongo parsing failed error =", err)
			continue // Skip to the next iteration if decoding fails
		}
		combineMongoStore = append(combineMongoStore, result)
	}

	apiRes.Data = combineMongoStore
	apiRes.Message = "SUCCESS"
	apiRes.Status = true
	return http.StatusOK, apiRes
}
