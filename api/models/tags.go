package models

type (
	NewTag struct {
		Name string `json:"name"`
	}

	Tag struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)
