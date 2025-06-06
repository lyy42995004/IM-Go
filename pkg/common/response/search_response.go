package response

import "github.com/lyy42995004/IM-Go/internal/model"

type SearchResponse struct {
	User model.User `json:"user"`
	Group model.Group `json:"group"`
}
