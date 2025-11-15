package initiation

import "poc/services/common/io/input"

type Handler interface {
	InitEntityB(input.Context)
}
