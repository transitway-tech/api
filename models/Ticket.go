package models

type Ticket struct {
  Lines []string `bson:"line" json:"line"`
  
}