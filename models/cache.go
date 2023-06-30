package models

import "time"

type CachedPrice struct {
	TimeStamp time.Time `json:"time_stamp"`
	Price     float64   `json:"price"`
}
