package input

import (
	"context"
)

type Context interface {
	Context() context.Context
	Querier() Querier
	Commander() Commander
}

type inputContext struct {
	ctx context.Context

	querier   Querier
	commander Commander
}

func WithCommand(ctx context.Context, commander Commander) (c Context) {
	return inputContext{commander: commander}
}

func (c inputContext) Context() context.Context {
	return c.ctx
}

func (c inputContext) Querier() Querier {
	return c.querier
}

func (c inputContext) Commander() Commander {
	return c.commander
}
