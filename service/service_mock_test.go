package service

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/timfame/vk-api.git/vk"
	"github.com/timfame/vk-api.git/vk/mocks"
	"math/rand"
	"testing"
	"time"
)

const timeAcceptableDifference = time.Microsecond

func TestPostStatsForLastHourWithMock(t *testing.T) {
	size := 200

	end := time.Now()
	start := end.Add(-time.Hour)
	startUnix, endUnix := start.Unix(), end.Unix()
	items := make([]vk.Item, 0, size)
	for i := 0; i < size; i++ {
		datetime := rand.Int63n(endUnix - startUnix) + startUnix
		items = append(items, vk.Item{Date: datetime})
	}

	request := &vk.NewsfeedSearchRequest{
		StartTime: start,
	}
	response := &vk.NewsfeedSearchResponse{Response: vk.ResponseBody{Items: items}}

	vkClient := &mocks.IClient{}
	vkClient.On("NewsfeedSearch", mock.Anything, mock.MatchedBy(func (req *vk.NewsfeedSearchRequest) bool {
		return req.StartTime.Sub(request.StartTime) < timeAcceptableDifference
	})).Return(response, nil)

	service := New(vkClient)

	stats, err := service.GetPostStatsForLastNHoursByQuery(context.Background(), "", 1)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(stats.Counts))
	assert.Equal(t, 1, stats.Counts[0].HoursAgo)
	assert.Equal(t, size, stats.Counts[0].Count)
}

func TestPostStatsForLast6HoursWithMock(t *testing.T) {
	maxSizePerHour := 200
	hours := 6
	now := time.Now()

	items := make([]vk.Item, 0, maxSizePerHour * hours)
	sizes := make([]int, 0, hours)
	end := now
	for h := 0; h < hours; h++ {
		start := end.Add(-time.Hour)
		startUnix, endUnix := start.Unix(), end.Unix()

		if h == hours - 1 {
			startUnix++
		}

		currentSize := rand.Intn(maxSizePerHour) + 1
		sizes = append(sizes, currentSize)

		for i := 0; i < currentSize; i++ {
			datetime := rand.Int63n(endUnix - startUnix) + startUnix
			items = append(items, vk.Item{Date: datetime})
		}

		end = start
	}

	request := &vk.NewsfeedSearchRequest{
		StartTime: end,
	}
	response := &vk.NewsfeedSearchResponse{Response: vk.ResponseBody{Items: items}}

	vkClient := &mocks.IClient{}
	vkClient.On("NewsfeedSearch", mock.Anything, mock.MatchedBy(func (req *vk.NewsfeedSearchRequest) bool {
		return req.StartTime.Sub(request.StartTime) < timeAcceptableDifference
	})).Return(response, nil)

	service := New(vkClient)

	stats, err := service.GetPostStatsForLastNHoursByQuery(context.Background(), "", hours)

	assert.Nil(t, err)
	assert.Equal(t, hours, len(stats.Counts))
	for i, count := range stats.Counts {
		assert.Equal(t, i + 1, count.HoursAgo, fmt.Sprintf("Stats for %d hours ago", i + 1))
		assert.Equal(t, sizes[i], count.Count, fmt.Sprintf("Stats for %d hours ago", i + 1))
	}
}
