package collection

import "context"

type BazaarCursor interface {
	All(ctx context.Context, v interface{}) error
	Next(ctx context.Context) bool
	Decode(v interface{}) error
	Close() error
}
