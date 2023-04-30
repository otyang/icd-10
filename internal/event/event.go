package event

const (
	SubjectFileUpload = "file.upload"
)

type FileUpload struct {
	Email   string
	Message string
}
