package models

import "time"

type Episode struct {
	ID int `json:"id"`
	AnimeID int `json:"anime_id"`
	Rank int `json:"rank"`
	Name string `json:"name"`
	Cover string `json:"cover"`
	Info string `json:"info"`
	Coin int `json:"coin"`
	Watch int `json:"watch"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
