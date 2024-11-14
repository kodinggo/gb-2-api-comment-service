package usecase

import (
	"context"

	"github.com/kodinggo/gb-2-api-story-service/internal/model"
)

type commentUseCase struct {
	commentRepository model.CommentRepository
}

func InitCommentUsecase(commentRepository model.CommentRepository) model.CommentRepository {
	return &commentUseCase{commentRepository: commentRepository}
}

func (u *commentUseCase) Create(ctx context.Context, data *model.Comment) (newComment model.Comment, err error) {
	return u.commentRepository.Create(ctx, data)
}
func (u *commentUseCase) Update(ctx context.Context ,data *[]model.Comment)(model.Comment,error){
	panic("Implement me!")
}