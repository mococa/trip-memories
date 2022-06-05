package entities

type Trip struct {
	PartitionKey string `json:"partition_key"`
	SortKey      string `json:"sort_key"`
	CreatedAt    int    `json:"created_at"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	DoneAt       string `json:"done_at"`
}
