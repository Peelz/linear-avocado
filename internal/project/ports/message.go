package ports

import "context"

type MessagePublisher interface {
	Publish(ctx context.Context, topic string, b []byte) error
}
