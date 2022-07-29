package repository

import (
	"context"

	"github.com/Trickster-ID/GO-CashFlow/model/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CashFlowRepo interface {
	//SelectAll() []entity.Balance)
	Save(cio entity.Cashinout) (*mongo.InsertOneResult, error)
	SelectAll() (*mongo.Cursor, error)
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
	collection := db.con.Database("cashflow").Collection("cashinout")
	return collection.InsertOne(context.TODO(), cio)
}

func (db *cfconn) SelectAll() (*mongo.Cursor, error) {
	collection := db.con.Database("cashflow").Collection("cashinout")
	return collection.Find(context.TODO(), bson.M{})
}
