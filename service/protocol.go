package service

type PostStats struct {
	Counts []HourStat
}

type HourStat struct {
	HoursAgo int
	Count    int
}