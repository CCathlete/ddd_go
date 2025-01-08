package entity

import "github.com/google/uuid"

// Represents a person in all domains.
type Person struct {
	ID   uuid.UUID
	Name string
	Age  int
}
