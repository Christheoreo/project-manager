package dtos

type (
	NewProjectComponentDto struct {
		Title       string      `json:"title"`
		Description string      `json:"description"`
		Data        interface{} `json:"data"`
	}

	ProjectComponentDto struct {
		ID          int         `json:"id"`
		Title       string      `json:"title"`
		Description string      `json:"description"`
		Data        interface{} `json:"data"`
	}

	NewProjectDto struct {
		Title       string                   `json:"title"`
		Description string                   `json:"description"`
		Tags        []string                 `json:"tags"`
		Priority    string                   `json:"priority"`
		Components  []NewProjectComponentDto `json:"components"`
	}

	ProjectDto struct {
		ID          int                   `json:"id"`
		Title       string                `json:"title"`
		Description string                `json:"description"`
		Tags        []string              `json:"tags"`
		Priority    string                `json:"priority"`
		Components  []ProjectComponentDto `json:"components"`
	}
)
