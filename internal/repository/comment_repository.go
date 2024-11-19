package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/kodinggo/gb-2-api-story-service/internal/model"
	"github.com/sirupsen/logrus"
)

type commentRepository struct {
	db *sql.DB
}

func InitCommentRepository(db *sql.DB) model.CommentRepository {
	return &commentRepository{db: db}
}

func (s *commentRepository) Create(ctx context.Context, data *model.Comment) (newComment model.Comment, err error) {
	timeNow := time.Now().UTC()
	results, err := sq.Insert("comments").Columns("user_id", "story_id", "comment", "created_at").
		Values(data.UserID, data.StoryID, data.Comment, timeNow).RunWith(s.db).ExecContext(ctx)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"user_id":    data.UserID,
			"story_id":   data.StoryID,
			"comment":    data.Comment,
			"created_at": timeNow,
		})
		return model.Comment{}, err
	}
	lastInsertId, _ := results.LastInsertId()
	newComment = model.Comment{ID: lastInsertId, UserID: data.UserID, StoryID: data.StoryID, Comment: data.Comment, Created_at: timeNow}
	return

}
func (s *commentRepository) FindById(ctx context.Context, id int64) (*model.Comment, error) {
	Query, args, err := sq.Select("*").From("comments").Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		return &model.Comment{}, err
	}
	row,err := s.db.QueryContext(ctx,Query, args...)
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()
	comment := &model.Comment{}

	if row.Next(){

		err = row.Scan(&comment.ID,&comment.Comment,&comment.UserID,&comment.StoryID,&comment.Created_at,&comment.Updated_at)
		if err !=nil{
			if err == sql.ErrNoRows{
				return nil,nil
			}
			return nil,err
		}
		
	}
	return comment, nil
}
func (s *commentRepository) Update(ctx context.Context, id int64, data *model.Comment) (*model.Comment, error) {
	query, args, err := sq.Update("comments").Set("comment", data.Comment).Set("user_id",data.UserID).Set("story_id",data.StoryID).Set("updated_at", time.Now().UTC()).Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		return nil,err
	}
	_, err = s.db.Exec(query, args...)
	if err != nil {
		return nil,err
	}
	results := model.Comment{ ID: data.ID,Comment: data.Comment,StoryID: data.StoryID,UserID: data.UserID}
	return &results, nil
}
