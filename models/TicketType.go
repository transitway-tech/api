package models

import (
	"api/db"

	"go.mongodb.org/mongo-driver/bson"
)

type TicketType struct {
  Name string `bson:"name" json:"name"`
  Fare float64 `bson:"fare" json:"fare"`
  City string `bson:"city" json:"city"`
  NoLines int `bson:"noLines" json:"noLines"`
  Expiry string `bson:"expiry" json:"expiry"`
}

func GetTicketTypes() ([]TicketType, error) {
    cursor, err := db.Tickets.Find(ctx, bson.M{})

    if err != nil {
      return []TicketType {}, err
    }

    ticketTypes := []TicketType {}

    err = cursor.All(ctx, &ticketTypes)
    if len(ticketTypes) == 0 {
      ticketTypes = []TicketType {}
    }

    return ticketTypes, err
}