package rest

import "time"

type Month struct {
	Persons     map[string]string `json:"persons"`
	Items       []Item            `json:"items"`
	Adjustments []Amount          `json:"adjustments"`
	Totals      []Amount          `json:"totals"`
}

type Item struct {
	Name     string    `json:"name"`
	Amount   int       `json:"amount"`
	PersonID string    `json:"personId"`
	Date     time.Time `json:"date"`
}

type Amount struct {
	PersonID string `json:"personId"`
	Amount   int    `json:"amount"`
}
