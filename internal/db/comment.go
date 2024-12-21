package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/yonyu/go-webapi/internal/domain"
)

// CommentRow - We have this type to handle null value and achieve better separation
type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Author sql.NullString
	Body   sql.NullString
}

func convertCommentRowToComment(c CommentRow) domain.Comment {
	return domain.Comment{
		ID:     c.ID,
		Slug:   c.Slug.String,
		Body:   c.Body.String,
		Author: c.Author.String,
	}
}

func (d *Database) GetComment(ctx context.Context, uuid string) (domain.Comment, error) {
	var commentRow CommentRow

	row := d.Client.QueryRowContext(
		ctx,
		`SELECT id, slug, body, author 
		FROM comments 
		WHERE id = $1`,
		uuid,
	)

	err := row.Scan(&commentRow.ID, &commentRow.Slug, &commentRow.Body, &commentRow.Author)
	if err != nil {
		return domain.Comment{}, fmt.Errorf("error fetching the domain by uuid: %w", err)
	}

	return convertCommentRowToComment(commentRow), nil
}
