package models

type Pagination struct {
	Limit    int `json:"limit"`
	Page     int `json:"page"`
	Previous int `json:"previous"`
	Next     int `json:"next"`
}
