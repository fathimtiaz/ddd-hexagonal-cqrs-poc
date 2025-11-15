package input

type Querier interface {
	AddFilter(field, operator string, value any)
	// AddOrder(field, direction string)
}

type BaseQuery struct {
	Filters []Filter
	Orders  []Order

	Pagination
}

func (bq BaseQuery) AddFilter(field, operator string, value any) {

}

type Filter struct {
	Field    string
	Operator string
	Value    any
}

type Order struct {
	Filed     string
	Direction string
}

type Pagination struct {
	Page  int
	Limit int
}
