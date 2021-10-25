package vk

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type client struct {
	httpClient *http.Client
	config     *Config
}

func NewClient(config *Config) *client {
	return &client{
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
		config: config,
	}
}

func (c *client) NewsfeedSearch(ctx context.Context, request *NewsfeedSearchRequest) (*NewsfeedSearchResponse, error) {
	startTimeValue := strconv.FormatInt(request.StartTime.Unix(), 10)
	count := request.Count
	if count == 0 {
		count = c.config.DefaultCount
	}
	countValue := strconv.Itoa(count)
	url := c.getDefaultURL(c.config.NewsFeedSearchMethod) + "&" +
		c.config.QueryParam + "=" + request.Query + "&" +
		c.config.StartTimeParam + "=" + startTimeValue + "&" +
		c.config.CountParam + "=" + countValue

	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	httpResponse, err := c.httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	var response NewsfeedSearchResponse
	if err := json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *client) getDefaultURL(method string) string {
	return c.config.APIScheme + "://" + c.config.APIUrl + "/" +
		c.config.MethodPath + "/" + method + "?" +
		c.config.VersionParam + "=" + c.config.Version + "&" +
		c.config.AccessTokenParam + "=" + c.config.AccessToken
}
