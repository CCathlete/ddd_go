package valueobject

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Amount    float64
	From      uuid.UUID
	To        uuid.UUID
	CreatedAt time.Time
}
