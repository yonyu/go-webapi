package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/satori/go.uuid"
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

func (d *Database) PostComment(ctx context.Context, comment domain.Comment) (domain.Comment, error) {
	comment.ID = uuid.NewV4().String()
	commentRow := CommentRow{
		ID:     comment.ID,
		Slug:   sql.NullString{String: comment.Slug, Valid: true},
		Author: sql.NullString{String: comment.Author, Valid: true},
		Body:   sql.NullString{String: comment.Body, Valid: true},
	}
	rows, err := d.Client.NamedQueryContext(
		ctx,
		`INSERT INTO comments (id, slug, author, body)
		VALUES
		(:id, :slug, :author, :body)`,
		commentRow,
	)

	if err != nil {
		return domain.Comment{}, fmt.Errorf("failied to insert comment: %w", err)
	}

	if err := rows.Close(); err != nil {
		return domain.Comment{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return comment, nil
}

func (d *Database) DeleteComment(ctx context.Context, id string) error {
	_, err := d.Client.ExecContext(
		ctx,
		`DELETE FROM comments WHERE id = $1`,
		id,
	)
	if err != nil {
		return fmt.Errorf("failed to delete comment from database: %w", err)
	}
	return nil
}

func (d *Database) UpdateComment(ctx context.Context, id string, comment domain.Comment) (domain.Comment, error) {
	commentRow := CommentRow{
		ID:     id,
		Slug:   sql.NullString{String: comment.Slug, Valid: true},
		Author: sql.NullString{String: comment.Author, Valid: true},
		Body:   sql.NullString{String: comment.Body, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(
		ctx,
		`UPDATE comments SET 
		slug = :slug,
		author = :author,
		body = :body
		WHERE id = :id`,
		commentRow,
	)
	if err != nil {
		return domain.Comment{}, fmt.Errorf("failed to update comment: %w", err)
	}
	if err := rows.Close(); err != nil {
		return domain.Comment{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return convertCommentRowToComment(commentRow), nil
}
