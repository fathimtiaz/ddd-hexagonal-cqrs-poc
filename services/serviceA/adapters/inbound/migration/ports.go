package migration

import "poc/services/common/io/input"

type Handler interface {
	RunMigration(input.Context)
}
