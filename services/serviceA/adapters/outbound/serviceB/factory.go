package serviceb

import "ddd-hexagonal-cqrs-poc/services/serviceA/adapters/outbound/serviceB/intraservice"

func NewServiceBCommands(adapter string) CommandCallers {
	switch adapter {
	case "intraservice":
		return intraservice.NewAdapter()
	// case "otherAdapter":
	// 	return &OtherAdapterServiceBCommander{}
	default:
		return nil
	}
}
