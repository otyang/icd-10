package entity

import "net/mail"

// CreateICDRequest handles the create ICD  request body
type CreateICDRequest struct {
	// CategoryCode always at least a 3lettered and above code
	CategoryCode string `conform:"trim" json:"categoryCode" validate:"required|min_len:3"`

	// DiagnosisCode could be null but has a max of 1 digits (9)
	DiagnosisCode *string `conform:"trim" json:"diagnosisCode" validate:""`

	// AbbreviatedDescription is a brief description of the diagnosis
	AbbreviatedDescription string `conform:"trim" json:"abbreviatedDescription" validate:"required|min_len:4"`

	// FullDescription is a description of the  diagnosis in full.
	FullDescription string `conform:"trim" json:"fullDescription" validate:"required|min_len:4"`

	// CategoryTitle title of category
	CategoryTitle string `conform:"trim" json:"categoryTitle" validate:"required|min_len:3"`
}

// FullCode is a combination of categorycode and DiagnosisCode
func (req CreateICDRequest) GetFullCode() string {
	return req.CategoryCode + *req.DiagnosisCode
}

type CreateICDResponse struct {
	CategoryCode           string  `json:"categoryCode" `
	DiagnosisCode          *string `json:"diagnosisCode"`
	AbbreviatedDescription string  `json:"abbreviatedDescription"`
	FullDescription        string  `json:"fullDescription"`
	CategoryTitle          string  `json:"categoryTitle"`
}

// EditICDRequest handles the edit ICD  request body
type EditICDRequest struct {
	CategoryCode           string  `conform:"trim" json:"categoryCode" validate:"required|min_len:3"`
	DiagnosisCode          *string `conform:"trim" json:"diagnosisCode" validate:""`
	AbbreviatedDescription string  `conform:"trim" json:"abbreviatedDescription" validate:"required|min_len:4"`
	FullDescription        string  `conform:"trim" json:"fullDescription" validate:"required|min_len:4"`
	CategoryTitle          string  `conform:"trim" json:"categoryTitle" validate:"required|min_len:3"`
}

// FullCode is a combination of categorycode and DiagnosisCode
func (req EditICDRequest) GetFullCode() string {
	if req.DiagnosisCode != nil {
		return req.CategoryCode + *req.DiagnosisCode
	}
	return req.CategoryCode
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
