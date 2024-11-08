package model

import (
	"context"
	"time"
)

type CommentsRepository interface {
	Create(ctx context.Context, user_id int64, story_id int64, comment string) (Comment, error)
}
type CommentUseCase interface {
	Create(ctx context.Context, user_id int64, story_id int64, comment string) (Comment, error)
}
type Comment struct {
	ID         int64      `json:"id"`
	Comment    string     `json:"comment"`
	Story_id   int64      `json:"story_id"`
	User_id    int64      `json:"user_id"`
	Created_at time.Time  `json:"created_at"`
	Updated_at *time.Time `json:"updated_at"`
}
