package aggreate

import "fmt"

var (
	ErrInvalidNamePerson = fmt.Errorf("a customer must have a valid name")
	ErrMissingValue      = fmt.Errorf("a product must have a name, a description and a price.")
)
