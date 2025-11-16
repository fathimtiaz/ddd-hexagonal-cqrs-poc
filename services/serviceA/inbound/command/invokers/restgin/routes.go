package restgin

import (
	"ddd-hexagonal-cqrs-poc/services/common/registry"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	invoker := NewInvokerWithHandler(registry.GetCommandHandlerA())

	r.PATCH("/entity-a/move", invoker.entityAMove)
}
