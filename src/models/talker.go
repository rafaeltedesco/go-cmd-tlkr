package models

type Talk struct {
	WatchedAt string `json:"watchedAt"`
	Rate      int    `json:"rate"`
}

type Talker struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Id   int    `json:"id"`
	Talk Talk   `json:"talk"`
}
