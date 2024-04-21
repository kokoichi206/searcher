package rdb

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kokoichi206/searcher/converter/rdb/model"
)

const selectAllBlogs = `
SELECT
	code,
	title,
	content,
	timestamp,
	start_time,
	end_time,
	cate,
	link,
	member_code,
	latest_comment_timestamp
FROM blogs
WHERE member_code = $1;
;
`

func (d *database) SelectAllBlogs(ctx context.Context, memberCode string) ([]*model.Blog, error) {
	rows, err := d.db.QueryContext(ctx, selectAllBlogs, memberCode)
	if err != nil {
		return nil, fmt.Errorf("failed to query: %w", err)
	}
	defer rows.Close()

	blogs := []*model.Blog{}

	for rows.Next() {
		m := &model.Blog{}
		var lct sql.NullString
		if err := rows.Scan(
			&m.Code,
			&m.Title,
			&m.Content,
			&m.Timestamp,
			&m.StartTime,
			&m.EndTime,
			&m.Cate,
			&m.Link,
			&m.MemberCode,
			&lct,
		); err != nil {
			return nil, fmt.Errorf("failed to scan: %w", err)
		}

		m.LatestCommentTimestamp = lct.String

		blogs = append(blogs, m)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate: %w", err)
	}

	return blogs, nil
}
