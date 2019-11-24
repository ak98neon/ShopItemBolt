package pagination

import "akudria/appleShop/db"

type PageRequest struct {
	Items       []*db.Item `json:"items"`
	TotalCount  int        `json:"total_count"`
	CountOfPage int        `json:"count_of_page"`
	CurrentPage int        `json:"current_page"`
}
