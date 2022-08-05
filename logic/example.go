package logic

import (
	"context"
)

var ExampleLogicInstance ExampleLogic

type ExampleLogic struct {
}

func (ExampleLogic) Ping(ctx context.Context, name string) (resp string, err error) {
	return "hello world!" + name, nil
}
