package datastore

import (
	"context"
	"database/sql"
	"errors"

	"github.com/otyang/icd-10/pkg/logger"
	"github.com/uptrace/bun"
)

// IsErrNotFound if the error returned is as named
func IsErrNotFound(err error) bool {
	switch {
	case
		errors.Is(err, sql.ErrNoRows):
		return true
	}
	return false
}

type IDBHelpers interface {
	Close(db OrmDB)
	AutoMigrateTables(ctx context.Context, db OrmDbTx, models ...any) error
	DeleteByPK(ctx context.Context, db OrmDbTx, modelsPtr any) error
	DeleteByCol(ctx context.Context, db OrmDbTx, modelsPtr, columnName string, columnValue any) error
	UpdateByPK(ctx context.Context, db OrmDbTx, modelsPtr any) error
	UpsertByPK(ctx context.Context, db OrmDbTx, modelsPtr any) error
	InsertByPK(ctx context.Context, db OrmDbTx, modelsPtr any, ignoreDuplicates bool) error
	FindByPK(ctx context.Context, db OrmDbTx, modelsPtr any) error
	FindByCol(ctx context.Context, db OrmDbTx, modelsPtr any, columnName string, columnValue any) error
}

var _ IDBHelpers = (*ORMHelpers)(nil)

type ORMHelpers struct {
	log logger.Interface
}

func NewDBHelpers(log logger.Interface) *ORMHelpers {
	return &ORMHelpers{log}
}

// Close: It closes the db connection
func (o *ORMHelpers) Close(db OrmDB) {
	if err := db.Close(); err != nil {
		o.log.Fatal("Error closing database: " + err.Error())
	}
}

func (o *ORMHelpers) DeleteByPK(ctx context.Context, db OrmDbTx, modelsPtr any) error {
	_, err := db.NewDelete().Model(modelsPtr).WherePK().Exec(ctx)
	return err
}

func (o *ORMHelpers) DeleteByCol(
	ctx context.Context, db OrmDbTx, modelsPtr, columnName string, columnValue any,
) error {
	_, err := db.NewDelete().Model(modelsPtr).Where("? = ?", bun.Ident(columnName), columnValue).Exec(ctx)
	return err
}

func (o *ORMHelpers) UpdateByPK(ctx context.Context, db OrmDbTx, modelsPtr any) error {
	_, err := db.NewUpdate().Model(modelsPtr).WherePK().Exec(ctx)
	return err
}

func (o *ORMHelpers) UpsertByPK(ctx context.Context, db OrmDbTx, modelsPtr any) error {
	_, err := db.NewInsert().Model(modelsPtr).On("CONFLICT DO UPDATE").Exec(ctx)
	return err
}

func (o *ORMHelpers) InsertByPK(ctx context.Context, db OrmDbTx, modelsPtr any, ignoreDuplicates bool) error {
	if ignoreDuplicates {
		_, err := db.NewInsert().Model(modelsPtr).Ignore().Exec(ctx)
		return err
	}
	_, err := db.NewInsert().Model(modelsPtr).Exec(ctx)
	return err
}

func (o *ORMHelpers) FindByPK(ctx context.Context, db OrmDbTx, modelsPtr any) error {
	err := db.NewSelect().Model(modelsPtr).WherePK().Scan(ctx)
	return err
}

func (o *ORMHelpers) FindByCol(ctx context.Context, db OrmDbTx, modelsPtr any, columnName string, columnValue any) error {
	err := db.NewSelect().Model(modelsPtr).Where("? = ?", bun.Ident(columnName), columnValue).Scan(ctx)
	return err
}

// AutoMigrateTables - automigrating tables from struct WHEN they dont exists.
// Usage: AutoMigrateTables(ctx, db, (*Model1)(nil), (*Model2)(nil))
func (o *ORMHelpers) AutoMigrateTables(ctx context.Context, db OrmDbTx, models ...any) error {
	for _, model := range models {
		if _, err := db.NewCreateTable().Model(model).IfNotExists().Exec(ctx); err != nil {
			errMsg := "failed creating schema resources: " + err.Error()
			o.log.Error(errMsg)
			return errors.New(errMsg)
		}
	}
	return nil
}
