package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type ICD struct {
	bun.BaseModel          `bun:"table:icd10_codes"`
	CategoryCode           string    `json:"categoryCode"`
	DiagnosisCode          *string   `json:"diagnosisCode"`
	FullCode               string    `json:"fullCode" bun:",pk"`
	AbbreviatedDescription string    `json:"abbreviatedDeScription"`
	FullDescription        string    `json:"fullDecription"`
	CategoryTitle          string    `json:"categoryTitle"`
	CreatedAt              time.Time `json:"createdAt"`
	UpdatedAt              time.Time `json:"updatedAt"`
}
