package message

import "context"

// Publisher Message queue publisher interface
// Use for Publish topic with byte payload
type Publisher interface {
	Publish(ctx context.Context, topic string, b []byte) (int64, error)
}
