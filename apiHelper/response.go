package apihelper

import (
	"bytes"
	"net/http"
	"video_search/constants"

	"github.com/gin-gonic/gin"
)

type APIRes struct {
	Status    bool        `json:"status"`
	Message   string      `json:"message"`
	ErrorCode string      `json:"errorcode"`
	Data      interface{} `json:"data"`
}

func CustomResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}

func CallApi(methodType, url string, payload *bytes.Buffer) (*http.Response, error) {
	req, _ := http.NewRequest(methodType, url, payload)

	authToken := constants.ApiKey

	req.Header.Add("Authorization", authToken)
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, err
}

func SendInternalServerError() (int, APIRes) {
	var apiRes APIRes
	apiRes.Status = false
	apiRes.Message = constants.InternalServerError
	apiRes.ErrorCode = "500"
	return http.StatusInternalServerError, apiRes
}
