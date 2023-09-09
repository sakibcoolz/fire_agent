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
	fmt.Println("Started")
	conf := config.GetDeviceInfo()
	fmt.Println("conf", conf)
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
	fmt.Println("request.Status:===> ", request.Status)
	if request.Status != http.StatusOK {

		return
	}
	apimodel.ToStruct(request.Response, agentconfig)
	r := gin.Default()
	app.UrlMapping(r, agentconfig)
	r.Run(":8081")
}
