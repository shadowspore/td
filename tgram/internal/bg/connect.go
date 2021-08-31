package bg

import (
	"context"
	"sync"
)

type Client interface {
	Run(ctx context.Context, f func(ctx context.Context) error) error
}

func Connect(ctx context.Context, client Client) (_ context.Context, closeFn func() error, err error) {
	var (
		callback = make(chan struct{})
		result   = make(chan error)
	)

	clientCtx, cancel := context.WithCancel(context.Background())

	go func() {
		result <- client.Run(clientCtx, func(ctx context.Context) error {
			close(callback)
			<-ctx.Done()
			return ctx.Err()
		})
	}()

	var (
		closeOnce sync.Once
		closeErr  error
	)
	select {
	case <-callback:
		return clientCtx, func() error {
			closeOnce.Do(func() {
				cancel()
				closeErr = <-result
			})
			return closeErr
		}, nil

	case <-ctx.Done():
		cancel()
		return nil, nil, ctx.Err()

	case err := <-result:
		cancel()
		return nil, nil, err
	}
}
