package dto

import "time"

type PasscodeFinishRequest struct {
	Id   string `json:"id"`
	Code string `json:"code"`
}

type PasscodeInitRequest struct {
	UserId string `json:"user_id"`
}

type PasscodeReturn struct {
	Id        string    `json:"id"`
	TTL       int       `json:"ttl"`
	CreatedAt time.Time `json:"created_at"`
}
