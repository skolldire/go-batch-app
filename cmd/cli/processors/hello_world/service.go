package hello_world

import (
	"context"
	"fmt"
)

type service struct {
}

var _ Service = (*service)(nil)

func NewService() *service {
	return &service{}
}

func (s service) Run(ctx context.Context) {
	fmt.Println("hello world")
}
