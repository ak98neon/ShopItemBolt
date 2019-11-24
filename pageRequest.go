package main

type PageRequest struct {
	Items       []*Item `json:"items"`
	TotalCount  int     `json:"total_count"`
	CountOfPage int     `json:"count_of_page"`
	CurrentPage int     `json:"current_page"`
}
