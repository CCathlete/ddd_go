package customers

import (
	aggreate "ddd-go/aggregate"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepo struct {
	db        *mongo.Database
	customers *mongo.Collection
}

// An internal used for storing the customer aggregate inside the mongo db.
type mongoCustomer struct {
	ID   uuid.UUID `bson:"id"`
	Name string    `bson:"name"`
}

func ConvertToMongoCustomer(cst aggreate.Customer) (mc mongoCustomer) {
	mc = mongoCustomer{
		ID:   cst.GetID(),
		Name: cst.GetName(),
	}

	return
}
