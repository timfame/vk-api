package vk

import "context"

type IClient interface {
	NewsfeedSearch(ctx context.Context, request *NewsfeedSearchRequest) (*NewsfeedSearchResponse, error)
}
