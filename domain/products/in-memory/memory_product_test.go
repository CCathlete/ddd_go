package inmemprod

import (
	aggreate "ddd-go/aggregate"
	"ddd-go/domain/products"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestGetProduct(t *testing.T) {
	type testCase struct {
		test        string
		id          uuid.UUID
		expectedErr error
		checkResult func(
			t *testing.T,
			res aggreate.Product,
		)
	}

	cst, err := aggreate.NewProduct(
		"Beer",
		"Liquid bread.",
		5.5,
	)
	if err != nil {
		t.Fatal(err)
	}

	id := cst.GetID()

	repo := MemoryProductRepo{
		products: map[uuid.UUID]aggreate.Product{
			id: cst,
		},
	}

	testCases := []testCase{
		{
			test:        "No product by id",
			id:          uuid.MustParse("e3b0c442-68ce-11e9-8e3a-0242ac120002"),
			expectedErr: products.ErrProductNotFound,
			checkResult: func(t *testing.T,
				res aggreate.Product) {

				expected := aggreate.Product{}
				if !reflect.DeepEqual(res, expected) {
					t.Fatal(
						fmt.Sprintf(
							"result is different than expected product"+
								"result: %v, expected: %v", res, cst,
						),
					)
				}

			},
		},
		{
			test:        "Happy case - product by id",
			id:          id,
			expectedErr: nil,
			checkResult: func(t *testing.T,
				res aggreate.Product) {

				expected := cst
				if !reflect.DeepEqual(res, expected) {
					t.Fatal(
						fmt.Sprintf(
							"result is different than expected product"+
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
