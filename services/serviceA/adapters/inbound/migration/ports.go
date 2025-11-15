package migration

import "ddd-hexagonal-cqrs-poc/services/common/io/input"

type Handler interface {
	RunMigration(input.Context)
}
