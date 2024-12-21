package db

import (
	"context"
	"database/sql"

	"github.com/yonyu/go-webapi/internal/comment"
)

type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Author sql.NullString
	Body   sql.NullString
}

func (d *Database) GetComment(ctx context.Context, uuid string) (comment.Comment, error) {
	return comment.Comment{}, nil
}
