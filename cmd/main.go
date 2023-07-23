package main

import (
	apimodel "fire_agent/apicaller"
	"fire_agent/app"
	"fire_agent/config"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	conf := config.GetDeviceInfo()
	agentconfig := new(config.AgentConfig)

	request := apimodel.New()
	request.Url = fmt.Sprintf("http://%s/enrollment", os.Getenv("SERVICE"))
	request.Body = conf
	request.Header = []apimodel.Header{
		{
			Key:   "Content-Type",
			Value: "application/json",
		},
	}

	request.Method = "POST"
	request.SetRequest()

	apimodel.ToStruct(request.Response, agentconfig)

	app.UrlMapping(r, agentconfig)

	r.Run(":8081")
}
