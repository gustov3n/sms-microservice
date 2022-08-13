// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: auth.sql

package query

import (
	"context"
	"database/sql"
)

const createOtp = `-- name: CreateOtp :execresult
INSERT into
    otp (phone_number, otp, expired_date)
VALUES
    (?, ?, ?)
`

type CreateOtpParams struct {
	PhoneNumber string `json:"phone_number"`
	Otp         string `json:"otp"`
	ExpiredDate string `json:"expired_date"`
}

func (q *Queries) CreateOtp(ctx context.Context, arg *CreateOtpParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createOtp, arg.PhoneNumber, arg.Otp, arg.ExpiredDate)
}

const getPhoneNumber = `-- name: GetPhoneNumber :one
SELECT
    id, phone_number, otp, expired_date
FROM
    otp
WHERE
    phone_number = ?
`

func (q *Queries) GetPhoneNumber(ctx context.Context, phoneNumber string) (*Otp, error) {
	row := q.db.QueryRowContext(ctx, getPhoneNumber, phoneNumber)
	var i Otp
	err := row.Scan(
		&i.ID,
		&i.PhoneNumber,
		&i.Otp,
		&i.ExpiredDate,
	)
	return &i, err
}

const updateOtp = `-- name: UpdateOtp :execresult
UPDATE
    otp
SET
    otp = ?,
    expired_date = ?
WHERE
    phone_number = ?
`

type UpdateOtpParams struct {
	Otp         string `json:"otp"`
	ExpiredDate string `json:"expired_date"`
	PhoneNumber string `json:"phone_number"`
}

func (q *Queries) UpdateOtp(ctx context.Context, arg *UpdateOtpParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateOtp, arg.Otp, arg.ExpiredDate, arg.PhoneNumber)
}