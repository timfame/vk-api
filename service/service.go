package service

import (
	"context"
	"github.com/timfame/vk-api.git/vk"
	"time"
)

type service struct {
	vkClient vk.IClient
}

func New(vkClient vk.IClient) *service {
	return &service{
		vkClient: vkClient,
	}
}

func (s *service) GetPostStatsForLastNHoursByQuery(ctx context.Context, query string, lastNHours int) (*PostStats, error) {
	now := time.Now()
	startTime := now.Add(-time.Hour * time.Duration(lastNHours))
	request := &vk.NewsfeedSearchRequest{
		Query:     query,
		StartTime: startTime,
	}
	response, err := s.vkClient.NewsfeedSearch(ctx, request)
	if err != nil {
		return nil, err
	}

	stats := &PostStats{Counts: make([]HourStat, 0, lastNHours)}
	end := now
	for hoursAgo := 1; hoursAgo <= lastNHours; hoursAgo++ {
		start := end.Add(-time.Hour)
		count := countItemsInInterval(response.Response.Items, start, end)
		stats.Counts = append(stats.Counts, HourStat{
			HoursAgo: hoursAgo,
			Count:    count,
		})
		end = start
	}

	return stats, nil
}

func countItemsInInterval(items []vk.Item, start, end time.Time) int {
	count := 0
	for _, item := range items {
		date := time.Unix(item.Date, 0)
		if date.After(start) && date.Before(end) {
			count++
		}
	}
	return count
}
