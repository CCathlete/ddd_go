package customers

import "fmt"

var (
	ErrCustomerNotFound = fmt.Errorf(
		"the customerwas not found in the repository",
	)
	ErrFailedToAddCustomer = fmt.Errorf(
		"failed to add the customer to the repository",
	)
	ErrUpdateCustomer = fmt.Errorf(
		"failed to update the customer's details",
	)
)
