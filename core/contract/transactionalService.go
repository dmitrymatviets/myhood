package contract

import "context"

type ITransactional interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) (err error)
}
