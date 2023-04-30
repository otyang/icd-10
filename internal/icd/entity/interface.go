package entity

import (
	"context"
	"database/sql"
	"errors"
)

type IICDRepository interface {
	CreateRecord(ctx context.Context, in CreateICDRequest) (*ICD, error)
	UpdateRecord(ctx context.Context, fullCode string, in ICD) (*ICD, error)
	GetRecord(ctx context.Context, id string) (*ICD, error)
	DeleteRecord(ctx context.Context, id string) error
	ListRecords(ctx context.Context, limit int, cursor string) (nextCursor *string, results []ICD, e error)
	HasPrevPage(ctx context.Context, onFirstPage bool, cursor string) (cursorPrevPage *string, err error)
}

var ErrICDRecordNotFound = errors.New("no record found for this code")

func IsErrNotFound(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
