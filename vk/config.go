package vk

type Config struct {
	APIScheme            string `json:"api_scheme"`
	APIUrl               string `json:"api_url"`
	MethodPath           string `json:"method_path"`
	NewsFeedSearchMethod string `json:"news_feed_search_method"`
	AccessTokenParam     string `json:"access_token_param"`
	AccessToken          string `json:"access_token"`
	VersionParam         string `json:"version_param"`
	Version              string `json:"version"`
	QueryParam           string `json:"query_param"`
	StartTimeParam       string `json:"start_time_param"`
	CountParam           string `json:"count_param"`
	DefaultCount         int    `json:"default_count"`
}
