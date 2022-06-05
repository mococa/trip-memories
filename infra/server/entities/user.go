package entities

type User struct {
	PartitionKey string `json:"partition_key"`
	SortKey      string `json:"sort_key"`
	GoogleId     string `json:"google_id"`
	CreatedAt    int64  `json:"created_at"`
	Picture      string `json:"picture"`
	Name         string `json:"name"`
	Email        string `json:"email"`
}
