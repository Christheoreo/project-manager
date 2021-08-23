package models

type (
	StandardResponse struct {
		Status   int      `json:"status"`
		Messages []string `json:"messages"`
	}
)
