// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: template.sql

package query

import (
	"context"
	"database/sql"
)

const createTemplate = `-- name: CreateTemplate :execresult
INSERT INTO template (
    id, 
    type, 
    apps_name, 
    text
) VALUES (
    ?, ?, ?, ?
)
`

type CreateTemplateParams struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	AppsName string `json:"apps_name"`
	Text     string `json:"text"`
}

func (q *Queries) CreateTemplate(ctx context.Context, arg *CreateTemplateParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createTemplate,
		arg.ID,
		arg.Type,
		arg.AppsName,
		arg.Text,
	)
}

const deleteTemplate = `-- name: DeleteTemplate :exec
UPDATE template
SET is_deleted = true
WHERE
    id = ?
AND
    is_deleted = false
`

func (q *Queries) DeleteTemplate(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteTemplate, id)
	return err
}

const getListTemplate = `-- name: GetListTemplate :many
SELECT
    id, 
    type, 
    apps_name, 
    text
FROM template
WHERE is_deleted = false
LIMIT 10
`

type GetListTemplateRow struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	AppsName string `json:"apps_name"`
	Text     string `json:"text"`
}

func (q *Queries) GetListTemplate(ctx context.Context) ([]*GetListTemplateRow, error) {
	rows, err := q.db.QueryContext(ctx, getListTemplate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetListTemplateRow
	for rows.Next() {
		var i GetListTemplateRow
		if err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.AppsName,
			&i.Text,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTemplateById = `-- name: GetTemplateById :one
SELECT
    id, 
    type, 
    apps_name, 
    text
FROM template
WHERE
 type = ?
 and apps_name = ?
 and is_deleted = false
LIMIT 1
`

type GetTemplateByIdParams struct {
	Type     string `json:"type"`
	AppsName string `json:"apps_name"`
}

type GetTemplateByIdRow struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	AppsName string `json:"apps_name"`
	Text     string `json:"text"`
}

func (q *Queries) GetTemplateById(ctx context.Context, arg *GetTemplateByIdParams) (*GetTemplateByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getTemplateById, arg.Type, arg.AppsName)
	var i GetTemplateByIdRow
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.AppsName,
		&i.Text,
	)
	return &i, err
}

const updateTemplate = `-- name: UpdateTemplate :exec
UPDATE template
SET 
    type = ?,
    apps_name = ?, 
    text = ?
WHERE
    id = ?
AND
    is_deleted = false
`

type UpdateTemplateParams struct {
	Type     string `json:"type"`
	AppsName string `json:"apps_name"`
	Text     string `json:"text"`
	ID       string `json:"id"`
}

func (q *Queries) UpdateTemplate(ctx context.Context, arg *UpdateTemplateParams) error {
	_, err := q.db.ExecContext(ctx, updateTemplate,
		arg.Type,
		arg.AppsName,
		arg.Text,
		arg.ID,
	)
	return err
}
