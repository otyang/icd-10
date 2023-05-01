package event

const (
	TopicFileUploadComplete = "file.upload.complete"
)

type FileUploadComplete struct {
	Email   string
	Message string
}
