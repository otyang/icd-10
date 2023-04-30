package bun

import (
	"context"
	"time"

	"github.com/otyang/icd-10/internal/icd/entity"
	"github.com/otyang/icd-10/pkg/datastore"
)

var _ entity.IICDRepository = (*ICDRepository)(nil)

type ICDRepository struct {
	db datastore.OrmDB
}

func NewICDRepository(db datastore.OrmDB) *ICDRepository {
	return &ICDRepository{db: db}
}

func (r *ICDRepository) CreateRecord(ctx context.Context, in entity.CreateICDRequest) (*entity.ICD, error) {
	c := entity.ICD{
		CategoryCode:           in.CategoryCode,
		DiagnosisCode:          in.DiagnosisCode,
		FullCode:               in.GetFullCode(),
		AbbreviatedDescription: in.AbbreviatedDescription,
		FullDescription:        in.FullDescription,
		CategoryTitle:          in.CategoryTitle,
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
	}

	_, err := r.db.
		NewInsert().
		Model(&c).
		Exec(ctx)
	return &c, err
}

func (r *ICDRepository) UpdateRecord(ctx context.Context, fullCode string, in entity.ICD) (*entity.ICD, error) {
	in.UpdatedAt = time.Now()
	err := r.db.
		NewUpdate().
		Model(&in).
		WherePK().
		// Where("full_code = ?", fullCode).
		Scan(ctx)
	return &in, err
}

func (r *ICDRepository) GetRecord(ctx context.Context, id string) (*entity.ICD, error) {
	mod := &entity.ICD{
		FullCode: id,
	}

	err := r.db.
		NewSelect().
		Model(mod).
		WherePK().
		Scan(ctx)
	return mod, err
}

func (r *ICDRepository) DeleteRecord(ctx context.Context, id string) error {
	_, err := r.db.
		NewDelete().
		Model(&entity.ICD{
			FullCode: id,
		}).
		WherePK().Exec(ctx)
	return err
}

func (r *ICDRepository) ListRecords(ctx context.Context, limit int, cursor string) (nextCursor *string, results []entity.ICD, e error) {
	records := []entity.ICD{}
	ourLimit := limit + 1

	count, err := r.db.NewSelect().
		Model(&records).
		Limit(ourLimit).
		Where("full_code >= ?", cursor).
		Order("full_code ASC").
		ScanAndCount(ctx)
	if err != nil {
		return nil, nil, err
	}

	// Next Page
	if count < ourLimit {
		return nil, records, nil
	}

	return &records[ourLimit-1].FullCode, records[:limit], nil
}

func (r *ICDRepository) HasPrevPage(ctx context.Context, onFirstPage bool, cursor string) (cursorPrevPage *string, err error) {
	var prevRecords entity.ICD

	if !onFirstPage {
		count, err := r.db.NewSelect().
			Model(&prevRecords).
			Limit(1).
			Where("full_code < ?", cursor).
			Order("full_code DESC").
			ScanAndCount(ctx)
		if err != nil {
			return nil, err
		}

		if count < 1 {
			return nil, nil
		}

		return &prevRecords.FullCode, nil
	}

	return nil, nil
}
