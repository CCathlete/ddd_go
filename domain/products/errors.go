package products

import "fmt"

var (
	ErrProductNotFound = fmt.Errorf(
		"the product was not found in the repository",
	)
	ErrFailedToAddProduct = fmt.Errorf(
		"failed to add the product to the repository",
	)
	ErrUpdateProduct = fmt.Errorf(
		"failed to update the product's details",
	)
	ErrDeleteProduct = fmt.Errorf(
		"failed to delete the product from the repo",
	)
)
