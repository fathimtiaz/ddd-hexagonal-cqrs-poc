package query

import (
	"ddd-hexagonal-cqrs-poc/services/serviceA/adapters/inbound/query/intraservice"
	"ddd-hexagonal-cqrs-poc/services/serviceA/adapters/outbound/repo"
	"errors"
)

func NewHandler(adapter string, reader repo.Reader) (Handler, error) {
	switch adapter {
	case "intraservice":
		return intraservice.New(reader), nil
	default:
		return nil, errors.New("ASDS")
	}
}
