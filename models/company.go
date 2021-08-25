package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Company struct {
	ID            bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Company_Name  string        `json:"company_name" bson:"company_name" nosql:"company_name" validate:"required"`
	Company_ID    string        `json:"company_id" bson:"company_id" nosql:"company_id" `
	Company_Code  string        `json:"company_code" bson:"company_code" nosql:"company_code" validate:"required"`
	Location_Name string        `json:"location_name" bson:"location_name" nosql:"location_name" `
	CreatedAt     time.Time     `json:"-" bson:"created_at"`
	UpdatedAt     time.Time     `json:"-" bson:"updated_at"`
}
