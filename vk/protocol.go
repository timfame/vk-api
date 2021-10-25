package vk

import "time"

type NewsfeedSearchRequest struct {
	Query     string
	StartTime time.Time
	Count     int
}

type NewsfeedSearchResponse struct {
	Response ResponseBody `json:"response"`
}

type ResponseBody struct {
	Items []Item `json:"items"`
}

type Item struct {
	Date int64 `json:"date"`
}
