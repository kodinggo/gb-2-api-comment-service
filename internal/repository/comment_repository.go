package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/kodinggo/gb-2-api-story-service/internal/model"
	"github.com/sirupsen/logrus"
)
type commentRepository struct{
	db *sql.DB
}

func InitCommentRepository(db *sql.DB)model.CommentRepository{
	return &commentRepository{db:db}
}

func  (s *commentRepository)Create(ctx context.Context,data *model.Comment)(newComment model.Comment,err error){
	timeNow := time.Now().UTC()
	results,err :=sq.Insert("comments").Columns("user_id","story_id","comment","created_at").
	Values(data.UserId,data.StoryId,data.Comment,timeNow).RunWith(s.db).ExecContext(ctx)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"user_id":data.UserId,
			"story_id":data.StoryId,
			"comment":data.Comment,
		})
		return model.Comment{},err
	}
	lastInsertId,_:= results.LastInsertId()
	newComment = model.Comment{ID:lastInsertId ,UserId:data.UserId,StoryId:data.StoryId,Comment:data.Comment,Created_at:timeNow,}
	return
	
}

func (s *commentRepository) Update(ctx context.Context,data *[]model.Comment)(model.Comment,error){
	panic("implement me!")
}