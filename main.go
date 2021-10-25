package main

import (
	"context"
	"fmt"
	"github.com/timfame/vk-api.git/service"
	"github.com/timfame/vk-api.git/vk"
)

const (
	accessToken = ""

	apiScheme            = "https"
	apiURL               = "api.vk.com"
	methodPath           = "method"
	newsFeedSearchMethod = "newsfeed.search"
	accessTokenParam     = "access_token"
	versionParam         = "v"
	version              = "5.131"
	queryParam           = "q"
	startTimeParam       = "start_time"
	countParam           = "count"
	defaultCount         = 200
)

func main() {
	background := context.Background()

	config := &vk.Config{
		APIScheme:            apiScheme,
		APIUrl:               apiURL,
		MethodPath:           methodPath,
		NewsFeedSearchMethod: newsFeedSearchMethod,
		AccessTokenParam:     accessTokenParam,
		AccessToken:          accessToken,
		VersionParam:         versionParam,
		Version:              version,
		QueryParam:           queryParam,
		StartTimeParam:       startTimeParam,
		CountParam:           countParam,
		DefaultCount:         defaultCount,
	}

	vkClient := vk.NewClient(config)

	srv := service.New(vkClient)

	stats, err := srv.GetPostStatsForLastNHoursByQuery(background, "president", 12)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(stats)
}
