package entity

import "github.com/google/uuid"

// Represents a person in all domains.
type Item struct {
	ID          uuid.UUID
	Name        string
	Description string
}
