package entities

type Trip struct {
	PartitionKey string   `json:"partition_key"`
	SortKey      string   `json:"sort_key"`
	Picture      string   `json:"picture"`
	CreatedAt    int      `json:"created_at"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	DoneAt       int      `json:"done_at"`
	Images       []string `json:"images"`
}
