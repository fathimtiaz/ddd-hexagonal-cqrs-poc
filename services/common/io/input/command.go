package input

type Commander interface {
	ValidateQuery() error
	ValidatePayload() error
}
