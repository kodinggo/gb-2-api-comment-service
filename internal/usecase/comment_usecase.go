package usecase

import (
	"context"
	"fmt"
	"log"

	pb "github.com/kodinggo/gb-2-api-account-service/pb/account"
	"github.com/kodinggo/gb-2-api-comment-service/internal/model"
)

type commentUseCase struct {
	commentRepository    model.CommentRepository
	accountServiceClient pb.AccountServiceClient
}

func InitCommentUsecase(
	commentRepository model.CommentRepository,
	accountService pb.AccountServiceClient,
) model.CommentRepository {
	return &commentUseCase{
		commentRepository:    commentRepository,
		accountServiceClient: accountService}
}

func (u *commentUseCase) Create(ctx context.Context, data *model.Comment) (newComment model.Comment, err error) {
	return u.commentRepository.Create(ctx, data)
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
		return nil, fmt.Errorf("comment with id %d not found", id)
	}

	existingComment.Comment = data.Comment
	existingComment.StoryID = data.StoryID
	existingComment.UserID = data.UserID

	commentUpdate, err := u.commentRepository.Update(ctx, id, existingComment)
	if err != nil {
		return nil, fmt.Errorf("failed to update comment with id %d :%w", id, err)
	}
	return commentUpdate, nil
}

func (u *commentUseCase) Delete(ctx context.Context, id int64) error {
	return u.commentRepository.Delete(ctx, id)
}

func (u *commentUseCase) FindByStoryId(ctx context.Context, id int64) ([]*model.Comment, error) {
	comments, err := u.commentRepository.FindByStoryId(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch comments: %w", err)
	}

	for idx, comment := range comments {
		resp, err := u.accountServiceClient.FindByID(ctx, &pb.FindByIDRequest{Id: comment.UserID})
		if err != nil {
			log.Printf("failed to fetch author for comment ID %d: %s", comment.ID, err)
			return nil, err
		}

		comments[idx].Author = model.Author{
			ID:         resp.Id,
			Fullname:   resp.Fullname,
			SortBio:    resp.SortBio,
			Gender:     resp.Gender.String(),
			PictureURL: resp.PictureUrl,
			Username:   resp.Username,
			Email:      resp.Email,
		}

	}

	return comments, nil

}
func (u *commentUseCase)FindByStoryIds(ctx context.Context,ids []int64) ([]*model.Comment,error){
	if len(ids) == 0 {
		return nil, fmt.Errorf("story_id cannot be zero")
	}
	comments, err := u.commentRepository.FindByStoryIds(ctx, ids)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch comments: %w", err)
	}
	return comments,nil
}
