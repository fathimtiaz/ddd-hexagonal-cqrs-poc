package query

import (
	"errors"
	"poc/services/serviceA/adapters/inbound/query/intraservice"
	"poc/services/serviceA/adapters/outbound/repo"
)

func NewHandler(adapter string, reader repo.Reader) (Handler, error) {
	switch adapter {
	case "intraservice":
		return intraservice.New(reader), nil
	default:
		return nil, errors.New("ASDS")
	}
}
