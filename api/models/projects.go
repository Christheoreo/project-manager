package models

type (
	NewProjectComponent struct {
		Title       string            `json:"title"`
		Description string            `json:"description"`
		Data        map[string]string `json:"data"`
	}

	ProjectComponent struct {
		ID          int             `json:"id"`
		Title       string          `json:"title"`
		Description string          `json:"description"`
		Data        []ComponentData `json:"data"`
	}

	ComponentData struct {
		ID    int    `json:"id"`
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	NewProject struct {
		Title       string                `json:"title"`
		Description string                `json:"description"`
		TagIDs      []int                 `json:"tagIds"`
		PriorityID  int                   `json:"priorityId"`
		Components  []NewProjectComponent `json:"components"`
	}

	Project struct {
		ID          int                `json:"id"`
		Title       string             `json:"title"`
		Description string             `json:"description"`
		Tags        []string           `json:"tags"`
		Priority    string             `json:"priority"`
		Components  []ProjectComponent `json:"components"`
	}
)
