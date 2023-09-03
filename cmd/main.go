package main

import (
	apimodel "fire_agent/apicaller"
	"fire_agent/app"
	"fire_agent/config"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	conf := config.GetDeviceInfo()
	agentconfig := new(config.LoginResponse)

	request := apimodel.New()
	request.Url = fmt.Sprintf("http://%s/login", os.Getenv("SERVICE"))
	request.Body = conf
	request.Header = []apimodel.Header{
		{
			Key:   "Content-Type",
			Value: "application/json",
		},
	}

	request.Method = "POST"
	request.SetRequest()

	if request.Status != http.StatusOK {
		return
	}

	fmt.Println("Status :", request.Status)

	fmt.Println("Body :", request.Response)

	apimodel.ToStruct(request.Response, agentconfig)

	app.UrlMapping(r, agentconfig)

	r.Run(":8081")
}
