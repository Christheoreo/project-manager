package dtos

type (
	StandardResponseDto struct {
		Status   int      `json:"status"`
		Messages []string `json:"messages"`
	}
)
