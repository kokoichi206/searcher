package rdb

import (
	"context"
	"fmt"

	"github.com/kokoichi206/searcher/converter/rdb/model"
)

const selectAllMembers = `
SELECT
	code,
	name,
	english_name,
	kana,
	cate,
	img,
	link,
	pick,
	god,
	under,
	birthday,
	blood,
	constellation,
	graduation,
	group_name
FROM members;
`

func (d *database) SelectAllMembers(ctx context.Context) ([]*model.Member, error) {
	rows, err := d.db.QueryContext(ctx, selectAllMembers)
	if err != nil {
		return nil, fmt.Errorf("failed to query: %w", err)
	}
	defer rows.Close()

	members := []*model.Member{}

	for rows.Next() {
		m := &model.Member{}
		if err := rows.Scan(
			&m.Code,
			&m.Name,
			&m.EnglishName,
			&m.Kana,
			&m.Cate,
			&m.Img,
			&m.Link,
			&m.Pick,
			&m.God,
			&m.Under,
			&m.Birthday,
			&m.Blood,
			&m.Constellation,
			&m.Graduation,
			&m.Groupcode,
		); err != nil {
			return nil, fmt.Errorf("failed to scan: %w", err)
		}

		members = append(members, m)
	}

	return members, nil
}
