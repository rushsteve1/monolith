package shared

import "context"

type Service interface {
	Addr() string
	Name() string
	UseFcgi() bool
	Serve(ctx context.Context) error
}
