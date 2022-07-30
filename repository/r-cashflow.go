package repository

import (
	"context"

	"github.com/Trickster-ID/GO-CashFlow/model/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CashFlowRepo interface {
	// Save(cio entity.Cashinout) (*mongo.InsertOneResult, error)
	SelectAll() (*mongo.Cursor, error)
	Select(id primitive.ObjectID) *mongo.SingleResult
	Save(cio entity.Cashinout) (*mongo.InsertOneResult, error)
	Update(id primitive.ObjectID, update primitive.D) (*mongo.UpdateResult, error)
	Delete(id primitive.ObjectID) error
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

func (db *cfconn) Select(id primitive.ObjectID) *mongo.SingleResult {
	collection := db.con.Database("cashflow").Collection("cashinout")
	return collection.FindOne(context.TODO(), entity.Cashinout{ID: id})
}

func (db *cfconn) Update(id primitive.ObjectID, update primitive.D) (*mongo.UpdateResult, error) {
	collection := db.con.Database("cashflow").Collection("cashinout")
	return collection.UpdateByID(context.TODO(), id, update)
}

func (db *cfconn) Delete(id primitive.ObjectID) error {
	collection := db.con.Database("cashflow").Collection("cashinout")
	_, err := collection.DeleteOne(context.TODO(), entity.Cashinout{ID: id})
	return err
}
