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

func InitCommentRepository(db *sql.DB)model.CommentsRepository{
	return &commentRepository{db:db}
}

func  (s *commentRepository)Create(ctx context.Context,user_id int64,story_id int64,comment string)(newComment model.Comment,err error){
	timeNow := time.Now().UTC()
	results,err :=sq.Insert("comments").Columns("user_id","story_id","comment","created_at").
	Values(user_id,story_id,comment,timeNow).RunWith(s.db).ExecContext(ctx)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"user_id":user_id,
			"story_id":story_id,
			"comment":comment,
		})
		return model.Comment{},err
	}
	lastInsertId,_:= results.LastInsertId()
	newComment = model.Comment{ID:lastInsertId ,User_id:user_id,Story_id: story_id,Comment:comment,Created_at:timeNow,}
	return
	
}