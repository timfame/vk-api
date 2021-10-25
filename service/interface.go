package service

import "context"

type IService interface {
	GetPostStatsForLastNHoursByQuery(ctx context.Context, query string, lastNHours int) (*PostStats, error)
}
