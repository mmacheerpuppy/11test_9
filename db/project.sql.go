// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: project.sql

package db

import (
	"context"
)

const ProjectGet = `-- name: ProjectGet :one
select id, default_build_id
from unweave.project
where id = $1
`

func (q *Queries) ProjectGet(ctx context.Context, id string) (UnweaveProject, error) {
	row := q.db.QueryRowContext(ctx, ProjectGet, id)
	var i UnweaveProject
	err := row.Scan(&i.ID, &i.DefaultBuildID)
	return i, err
}
