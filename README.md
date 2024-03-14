
# Video Query

1 Dump data from YouTube by cron jobs with interval of 30 minute from YouTube api.

2 Api to fetch Data from db and display it in paginated form sorted in reverse chronological order.

3 Api to fetch Data based on Title and Description sorted in reverse chronological order.


## Tech Stack

Backend - Golang

Database - NoSql(MongoDB)


## Client Api
Their are 2 apis both are GET apis
1. GetStoredVideoDetails - Fetch vidoes details from db sorted in reverse chronological order and display it.
Query Param - pageNumber, pageSize

2. searchQueryBasedVideo - Fetch Data based on Title and Description sorted in reverse chronological order and display it.
Query Param - title, description


## Project Setup
Their can be 2 ways to start the Project:

1. Setup Manually
Prerequisite - Go

Step to run project:
- Clone project - (https://github.com/VARUNSHARMA55665/youtube-video-search.git)
- Run go mod tidy
- Run go mod vendor
- go build

2 Docker
- Clone project - (https://github.com/VARUNSHARMA55665/youtube-video-search.git)
- Run :- docker build -t youtube-video .
- Run :- docker run -it --name youtube -p 8080:8080 youtube-video

Note - Run above commands on the same directory where Dockerfile is present


## Curl
1. GetStoredVideoDetails :
curl --location 'localhost:8080/details/getStoredVideoDetails?pageNumber=2&pageSize=15'

2. SearchQueryBasedVideo :
curl --location 'localhost:8080/details/searchQueryBasedVideo?title=virat%20kohli&description=India%20win%20the%20world%20cup'
