package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Balance struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User           string             `json:"user,omitempty" bson:"user,omitempty"`
	MonthlyBalance int                `json:"monthlybalance,omitempty" bson:"monthlybalance,omitempty"`
	MonthlyIn      int                `json:"monthlyin,omitempty" bson:"monthlyin,omitempty"`
	MonthlyOut     int                `json:"monthlyout,omitempty" bson:"monthlyout,omitempty"`
	AnnualBalance  int                `jsonL:"annualbalance,omitempty" bson:"annualbalance,omitempty"`
	AnnualIn       int                `jsonL:"annualin,omitempty" bson:"annualin,omitempty"`
	AnnualOut      int                `jsonL:"annualout,omitempty" bson:"annualout,omitempty"`
}
