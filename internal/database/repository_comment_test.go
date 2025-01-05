//go:build integration
// +build integration

package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yonyu/go-webapi/internal/domain"
)

func TestCommentDatabase(t *testing.T) {
	t.Run("test create comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		comment, err := db.PostComment(context.Background(), domain.Comment{
			Slug:   "slug",
			Author: "author",
			Body:   "body",
		})
		assert.NoError(t, err)

		newComment, err := db.GetComment(context.Background(), comment.ID)
		assert.NoError(t, err)
		assert.Equal(t, "slug", newComment.Slug)
	})

	t.Run("test delete comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		comment, err := db.PostComment(context.Background(), domain.Comment{
			Slug:   "new-slug",
			Author: "Kenneth",
			Body:   "body",
		})
		assert.NoError(t, err)

		err = db.DeleteComment(context.Background(), comment.ID)
		assert.NoError(t, err)

		_, err = db.GetComment(context.Background(), comment.ID)
		assert.Error(t, err)
	})
}
