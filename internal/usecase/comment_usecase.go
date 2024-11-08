package usecase

import (
	"context"

	"github.com/kodinggo/gb-2-api-story-service/internal/model"
)


type commentUseCase struct{
	commentRepository model.CommentsRepository
}

func InitCommentUsecase(commentRepository model.CommentsRepository)model.CommentsRepository{
	return &commentUseCase{commentRepository: commentRepository}
}

func (u *commentUseCase) Create(ctx context.Context, user_id int64 , story_id int64,comment string) (newComment model.Comment, err error) {
	return u.commentRepository.Create(ctx, user_id,story_id,comment)
}