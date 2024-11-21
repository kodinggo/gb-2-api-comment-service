package model

import (
	"context"
	"time"
)

type CommentRepository interface {
	Create(ctx context.Context, data *Comment) (Comment, error)
	Update(ctx context.Context, id int64, data *Comment) (*Comment, error)
	FindById(ctx context.Context, id int64) (*Comment, error)
	Delete(ctx context.Context,id int64) error
}
type CommentUseCase interface {
	Create(ctx context.Context, data *Comment) (Comment, error)
	Update(ctx context.Context, id int64, data *Comment) (*Comment, error)
	FindById(ctx context.Context, id int64) (*Comment, error)
	Delete(ctx context.Context,id int64) error
}
type Comment struct {
	ID         int64      `json:"id"`
	Comment    string     `json:"comment" validate:"required"`
	StoryID    int64      `json:"story_id" validate:"required"`
	UserID     int64      `json:"user_id" validate:"required"`
	Author     Author     `json:"-"`
	Story      Story      `json:"-"`
	Created_at time.Time  `json:"created_at"`
	Updated_at *time.Time `json:"updated_at"`
}
type Author struct {
	ID         int64  `json:"id"`
	Fullname   string `json:"fullname"`
	SortBio    string `json:"sort_bio"`
	Gender     string `json:"gender"`
	PictureUrl string `json:"picture_url"`
	Username   string `json:"username"`
	Email      string `json:"email"`
}

type Story struct {
	ID          int64          `json:"id"`
	Title       string         `json:"title"`
	Content     string         `json:"content"`
	ThumbnaiUrl string         `json:"thumbnail_url"`
	Category    map[string]int `json:"category"`
	UserId      int64          `json:"user_id"`
	Created_at  time.Time      `json:"created_at"`
	Updated_at  time.Time      `json:"updated_at"`
	Deleted_at  time.Time      `json:"deleted_at"`
}
