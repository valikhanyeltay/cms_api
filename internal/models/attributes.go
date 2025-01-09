package models

type Attributes struct {
	ID        int64  `json:"id"`
	ContentID int64  `json:"content_id"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}
