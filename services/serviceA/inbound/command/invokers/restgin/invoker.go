package restgin

import (
	"context"

	inboundCommand "ddd-hexagonal-cqrs-poc/services/serviceA/inbound/command"
	inputCommand "ddd-hexagonal-cqrs-poc/services/serviceA/io/input/command"

	"github.com/gin-gonic/gin"
)

type Invoker struct {
	handler inboundCommand.Handler
}

// NewInvokerWithHandler constructs an inbound invoker with injected handler.
func NewInvokerWithHandler(handler inboundCommand.Handler) *Invoker {
	return &Invoker{handler}
}

func (i *Invoker) entityAMove(c *gin.Context) {
	if err := i.handler.EntityAMove(c.Request.Context(), inputCommand.EntityAMoveCommand{}); err != nil {

	}
}

func (i *Invoker) EntityACreate(ctx context.Context, cmd inputCommand.EntityACreateCommand) error {
	return i.handler.EntityACreate(ctx, cmd)
}
