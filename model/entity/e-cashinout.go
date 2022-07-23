package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cashinout struct {
	ID          primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Type        string              `json:"type,omitempty" bson:"type,omitempty"`
	Date        primitive.DateTime  `json:"date,omitempty" bson:"date,omitempty"`
	Category    string              `json:"category,omitempty" bson:"category,omitempty"`
	Total       int                 `json:"total,omitempty" bson:"total,omitempty"`
	Description string              `json:"description,omitempty" bson:"description,omitempty"`
	CreatedBy   string              `json:"createdby,omitempty" bson:"createdby,omitempty"`
	CreatedDate primitive.Timestamp `json:"createddate,omitempty" bson:"createddate,omitempty"`
	UpdatedBy   string              `json:"updatedby,omitempty" bson:"updatedby,omitempty"`
	UpdatedDate primitive.Timestamp `json:"updateddate,omitempty" bson:"updateddate,omitempty"`
}
