package model

import (
	"context"
	"time"
)

type CommentRepository interface {
	Create(ctx context.Context, data *Comment) (Comment, error)
	Update(ctx context.Context, data *[]Comment) (Comment, error)
}
type CommentUseCase interface {
	Create(ctx context.Context, data *Comment) (Comment, error)
	Update(ctx context.Context, data *[]Comment) (Comment, error)
}

type Author struct {
	ID         int64  `json:"id"`
	FullName   string `json:"fullname"`
	SortBio    string `json:"sort_bio"`
	Gender     string `json:"gender"`
	PictureUrl string `json:"picture_url"`
	Username   string `json:"username"`
	Email      string `json:"email"`
}
type Comment struct {
	ID         int64      `json:"id"`
	Comment    string     `json:"comment"`
	Author     Author     `json:"author"`
	StoryId    int64      `json:"-"`
	UserId     int64      `json:"-"`
	Created_at time.Time  `json:"created_at"`
	Updated_at *time.Time `json:"updated_at"`
}
type Stories struct {
	ID           int64  `json:"id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	ThumbnailUrl string `json:"thumbnail_url"`
	CategoryID   string `json:"categoryID"`
	UserId       string `json:"user_id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    string `json:"deleted_at"`
}
