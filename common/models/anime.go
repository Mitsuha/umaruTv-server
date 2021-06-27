package models

import "time"

type Anime struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Alias string `json:"alias"`
	Cover string `json:"cover"`
	Watch int `json:"watch"`
	Follow int `json:"follow"`
	Danmaku int `json:"danmaku"`
	ReleaseTime *time.Time `json:"release_time"`
	Episodes int `json:"episodes"`
	Status int `json:"status"`
	UpdateTime int `json:"update_time"`
	SeasonId int `json:"season_id"`
	Score float32 `json:"score"`
	SeasonName string `json:"season_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
