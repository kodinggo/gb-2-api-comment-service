package usecase

import (
	"context"
	"fmt"

	"github.com/kodinggo/gb-2-api-story-service/internal/model"
)

type commentUseCase struct {
	commentRepository model.CommentRepository
}

func InitCommentUsecase(commentRepository model.CommentRepository) model.CommentRepository {
	return &commentUseCase{commentRepository: commentRepository}
}
func (u *commentUseCase) Create(ctx context.Context, data *model.Comment) (newComment model.Comment, err error) {
	return u.commentRepository.Create(ctx, &model.Comment{Comment: data.Comment, StoryID: data.StoryID, UserID: data.UserID})
}
func (u *commentUseCase) FindById(ctx context.Context, id int64) (*model.Comment, error) {
	return u.commentRepository.FindById(ctx, id)
}
func (u *commentUseCase) Update(ctx context.Context, id int64, data *model.Comment) (*model.Comment, error) {

	existingComment, err := u.FindById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find comment with id %d: %w", id, err)
	}
	if existingComment == nil {
		return nil, fmt.Errorf("comment with id %d not found",id)
	}	

	if data.Comment != ""{
		existingComment.Comment = data.Comment
	}
	if data.StoryID != 0 {
		existingComment.StoryID = data.StoryID
	}
	if data.UserID != 0 {
		existingComment.UserID = data.UserID
	}

	commentUpdate, err := u.commentRepository.Update(ctx, id,existingComment)
	if err !=nil{
		return nil,fmt.Errorf("failed to update comment with id %d :%w",id,err)
	}
	return commentUpdate,nil
}
