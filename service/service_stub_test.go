package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/timfame/vk-api.git/vk"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

const (
	accessToken = "password"

	apiScheme            = "http"
	methodPath           = "method"
	newsFeedSearchMethod = "search"
	accessTokenParam     = "access_token"
	versionParam         = "v"
	version              = "test"
	queryParam           = "q"
	startTimeParam       = "start_time"
	countParam           = "count"
	defaultCount         = 200
)

func TestPostStatsForLastHourWithStub(t *testing.T) {
	hours := 12

	httpServer := getStubServer()
	defer httpServer.Close()

	vkConfig := getVKStubConfig(httpServer.URL)

	vkClient := vk.NewClient(vkConfig)

	service := New(vkClient)

	stats, err := service.GetPostStatsForLastNHoursByQuery(context.Background(), "", hours)

	assert.Nil(t, err)
	assert.Equal(t, hours, len(stats.Counts))
	for i, count := range stats.Counts {
		assert.Equal(t, i + 1, count.HoursAgo)
		if i < 7 || i == 9 {
			assert.Equal(t, 1, count.Count, fmt.Sprintf("Stats for %d hours ago", i + 1))
		} else {
			assert.Equal(t, 0, count.Count, fmt.Sprintf("Stats for %d hours ago", i + 1))
		}
	}
}

func getVKStubConfig(url string) *vk.Config {
	urlWithoutScheme := url[7:]
	return &vk.Config{
		APIScheme:            apiScheme,
		APIUrl:               urlWithoutScheme,
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
}

func getStubServer() *httptest.Server {
	router := mux.NewRouter()
	router.HandleFunc("/" + methodPath + "/" + newsFeedSearchMethod, stubHandler).Methods(http.MethodGet)
	return httptest.NewServer(router)
}

func stubHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if token := query.Get(accessTokenParam); token != accessToken {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if ver := query.Get(versionParam); ver != version {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := strconv.Atoi(query.Get(countParam))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = strconv.ParseInt(query.Get(startTimeParam), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	now := time.Now().Add(-time.Second)
	response := testHTTPResponse{Response: testHTTPResponseBody{Items: []testHTTPItem{
		{
			ID:   rand.Int63(),
			Date: now.Unix(),
		},
		{
			ID:   rand.Int63(),
			Date: now.Add(-time.Hour).Unix(),
		},
		{
			ID:   rand.Int63(),
			Date: now.Add(-time.Hour * 2).Unix(),
		},
		{
			ID:   rand.Int63(),
			Date: now.Add(-time.Hour * 3).Unix(),
		},
		{
			ID:   rand.Int63(),
			Date: now.Add(-time.Hour * 4).Unix(),
		},
		{
			ID:   rand.Int63(),
			Date: now.Add(-time.Hour * 5).Unix(),
		},
		{
			ID:   rand.Int63(),
			Date: now.Add(-time.Hour * 6).Unix(),
		},
		{
			ID:   rand.Int63(),
			Date: now.Add(-time.Hour * 9).Unix(),
		},
	}}}
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type testHTTPResponse struct {
	Response testHTTPResponseBody `json:"response"`
}

type testHTTPResponseBody struct {
	Items []testHTTPItem `json:"items"`
}

type testHTTPItem struct {
	ID   int64 `json:"id"`
	Date int64 `json:"date"`
}
