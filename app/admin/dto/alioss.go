package dto

import "time"

type BucketsListItemOutput struct {
	Name         string    `json:"Name"`
	Location     string    `json:"Location"`
	StorageClass string    `json:"StorageClass"`
	CreationDate time.Time `json:"CreationDate"`
}

//对像返回值
type ObjectListItemOutput struct {
	Key         string    `json:"Key"`
	Type     string    `json:"Type"`
	ETag string    `json:"ETag"`
	StorageClass string    `json:"StorageClass"`
	LastModified time.Time `json:"LastModified"`
}
