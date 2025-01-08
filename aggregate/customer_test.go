package aggreate

import (
	"errors"
	"testing"
)

func TestCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Empty name validation",
			name:        "",
			expectedErr: ErrInvalidNamePerson,
		},
		{
			test:        "Valid name",
			name:        "KenC",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := NewCustomer(tc.name)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf(
					"expected error: %v, got %v", tc.expectedErr, err,
				)
			}
		})
	}
}
