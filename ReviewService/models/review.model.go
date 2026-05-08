package models

import "time"

type Review struct {
	Id        int64     `json:"id"`
	BookingId int64     `json:"BookingId"`
	UserId    int64     `json:"userId"`
	Comment   string    `json:"review"`
	Rating    int64     `json:"rating"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
