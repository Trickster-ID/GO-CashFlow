package repository

import (
	"context"

	"github.com/Trickster-ID/GO-CashFlow/model/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type CashFlowRepo interface {
	//SelectAll() []entity.Balance)
	Save(cio entity.Cashinout) (*mongo.InsertOneResult, error)
}

type cfconn struct {
	con *mongo.Client
}

func NewCashFlowRepo(DB *mongo.Client) CashFlowRepo {
	return &cfconn{
		con: DB,
	}
}

var client *mongo.Client

func (db *cfconn) Save(cio entity.Cashinout) (*mongo.InsertOneResult, error) {
	//collection := config.ResolveClientDB().Database("CashFlow").Collection("CashInOut")
	collection := db.con.Database("CashFlow").Collection("CashInOut")
	result, err := collection.InsertOne(context.TODO(), cio)
	return result, err
}
