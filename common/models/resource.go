package models

import "time"

type Resource struct {
	ID int `json:"id"`
	EpisodeId int `json:"episode_id"`
	Source string `json:"source"`
	Type int `json:"type"`
	Resolution int `json:"resolution"`
	Ranking int `json:"ranking"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
