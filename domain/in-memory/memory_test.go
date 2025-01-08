package inmemory

import (
	aggreate "ddd-go/aggregate"
	"ddd-go/domain/customer"
	"testing"

	"github.com/google/uuid"
)

func TestGetCustomer(t *testing.T) {
	type testCase struct {
		test        string
		id          uuid.UUID
		expectedErr error
	}

	cst, err := aggreate.NewCustomer("KenC")
	if err != nil {
		t.Fatal(err)
	}

	id := cst.GetID()

	repo := MemoryRepo{
		customers: map[uuid.UUID]aggreate.Customer{
			id: cst,
		},
	}

	testCases := []testCase{
		{
			test:        "No customer by id",
			id:          uuid.MustParse("e3b0c442-68ce-11e9-8e3a-0242ac120002"),
			expectedErr: customer.ErrCustomerNotFound,
		},
		{
			// TODO: Continue.
			test: "Happy case - customer by id",
		},
	}
}
