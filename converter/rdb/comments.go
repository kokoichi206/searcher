package rdb

import (
	"context"
	"fmt"

	"github.com/kokoichi206/searcher/converter/rdb/model"
)

const selectAllComments = `
SELECT
	code,
	blog_code,
	content,
	name,
	timestamp
FROM comments
WHERE blog_code = $1;
;
`

func (d *database) SelectAllComments(ctx context.Context, blogCode string) ([]*model.Comment, error) {
	rows, err := d.db.QueryContext(ctx, selectAllComments, blogCode)
	if err != nil {
		return nil, fmt.Errorf("failed to query: %w", err)
	}
	defer rows.Close()

	comments := []*model.Comment{}

	for rows.Next() {
		c := &model.Comment{}
		if err := rows.Scan(
			&c.Code,
			&c.BlogCode,
			&c.Content,
			&c.Name,
			&c.Timestamp,
		); err != nil {
			return nil, fmt.Errorf("failed to scan: %w", err)
		}

		comments = append(comments, c)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate: %w", err)
	}

	return comments, nil
}
