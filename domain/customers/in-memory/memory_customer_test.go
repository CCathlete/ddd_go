package inmemcust

import (
	aggreate "ddd-go/aggregate"
	"ddd-go/domain/customers"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestGetCustomer(t *testing.T) {
	type testCase struct {
		test        string
		id          uuid.UUID
		expectedErr error
		checkResult func(
			t *testing.T,
			res aggreate.Customer,
		)
	}

	cst, err := aggreate.NewCustomer("KenC")
	if err != nil {
		t.Fatal(err)
	}

	id := cst.GetID()

	repo := MemoryCustomerRepo{
		customers: map[uuid.UUID]aggreate.Customer{
			id: cst,
		},
	}

	testCases := []testCase{
		{
			test:        "No customer by id",
			id:          uuid.MustParse("e3b0c442-68ce-11e9-8e3a-0242ac120002"),
			expectedErr: customers.ErrCustomerNotFound,
			checkResult: func(t *testing.T,
				res aggreate.Customer) {

				expected := aggreate.Customer{}
				if !reflect.DeepEqual(res, expected) {
					t.Fatal(
						fmt.Sprintf(
							"result is different than expected customer"+
								"result: %v, expected: %v", res, cst,
						),
					)
				}

			},
		},
		{
			test:        "Happy case - customer by id",
			id:          id,
			expectedErr: nil,
			checkResult: func(t *testing.T,
				res aggreate.Customer) {

				expected := cst
				if !reflect.DeepEqual(res, expected) {
					t.Fatal(
						fmt.Sprintf(
							"result is different than expected customer"+
								"result: %v, expected: %v", res, cst,
						),
					)
				}

			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			result, err := repo.Get(tc.id)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expecter error: %v, got %v",
					tc.expectedErr, err)
			}

			tc.checkResult(t, result)
		})
	}
}
