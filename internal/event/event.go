package event

import "github.com/gookit/event"

const (
	TopicFileUploadComplete = "file.upload.complete"
)

// Publisher
const (
	IndexDataKey = "publisher"
)

func Publish(topic string, data any) {
	event.Fire(topic, event.M{"0": data})
}

// Subscriber
type (
	Listener     = event.Listener
	ListenerFunc = event.ListenerFunc
)

const (
	Normal = event.Normal
)

func Subscriber(name string, listener Listener, priority ...int) {
	event.On(name, listener, priority...)
}
