package dtos

type (
	NewTagDto struct {
		Name string `json:"name"`
	}

	TagDto struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)
