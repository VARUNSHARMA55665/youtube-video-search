package helpers

import (
	"log"
	"video_search/constants"
	"video_search/database"
	"video_search/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDumper(mongoStore model.MongoStore) {

	var getMongoStore model.MongoStore
	collection := constants.VideoDetails
	filter := bson.D{{"videoid", mongoStore.VideoID}}
	err := database.FindOneMongo(collection, filter, &getMongoStore)
	if err != nil && err.Error() != constants.MongoNoDocError {
		log.Print("MongoDumper Error in finding to mongo with videoId:", mongoStore.VideoID, " err: ", err)
	}

	if getMongoStore.VideoID == "" {
		update := bson.D{{"$set", mongoStore}}
		opts := options.Update().SetUpsert(true)
		err = database.UpdateOneMongo(collection, filter, update, opts)
		if err != nil {
			log.Print("MongoDumper Error in updating to mongo with videoId:", mongoStore.VideoID, " err: ", err)
		}
	}

}
