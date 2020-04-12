package rest

type Month struct {
	Persons     map[string]string `json:"persons"`
	Items       []Item            `json:"items"`
	Adjustments []Amount          `json:"adjustments"`
	Totals      []Amount          `json:"totals"`
}

type Item struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Amount   int    `json:"amount"`
	PersonID string `json:"personId"`
	Date     string `json:"date"`
}

type Amount struct {
	PersonID string `json:"personId"`
	Amount   int    `json:"amount"`
}
