package hello_world

import "context"

type Service interface {
	Run(ctx context.Context)
}
