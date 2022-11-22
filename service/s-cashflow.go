package service

import (
	"context"
	"time"

	"github.com/Trickster-ID/GO-CashFlow/model/entity"
	"github.com/Trickster-ID/GO-CashFlow/repository"
	"github.com/Trickster-ID/go-libpik"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CashFlowSvc interface {
	Save(cd entity.Cashinout) (any, error)
	GetAll() ([]entity.Cashinout, error)
	Get(id string) (entity.Cashinout, error)
	Update(id string, cd entity.Cashinout) (any, error)
	Delete(id string) (any, error)
}

type cfsvc struct {
	cfrepo repository.CashFlowRepo
}

func NewCashFlowSvc(Cfrepo repository.CashFlowRepo) CashFlowSvc {
	return &cfsvc{
		cfrepo: Cfrepo,
	}
}

var ce entity.Cashinout

func (s *cfsvc) Save(cd entity.Cashinout) (any, error) {
	cd.CreatedBy = "admin"
	cd.CreatedDate = time.Now()
	cd.UpdatedBy = "admin"
	cd.UpdatedDate = time.Now()
	res, err := s.cfrepo.Save(cd)
	if err != nil {
		return nil, err
	}
	cd.ID = res.InsertedID.(primitive.ObjectID)
	return cd, err
}

func (s *cfsvc) GetAll() ([]entity.Cashinout, error) {
	var cios []entity.Cashinout
	cursor, err := s.cfrepo.SelectAll()
	if err != nil {
		return []entity.Cashinout{}, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		cursor.Decode(&ce)
		cios = append(cios, ce)
	}
	if err := cursor.Err(); err != nil {
		return []entity.Cashinout{}, err
	}
	return cios, err
}

func (s *cfsvc) Get(id string) (entity.Cashinout, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ce, err
	}
	res := s.cfrepo.Select(oid)
	errd := res.Decode(&ce)
	if errd != nil {
		return ce, errd
	}
	return ce, nil
}

func (s *cfsvc) Update(id string, cd entity.Cashinout) (any, error) {
	oid, errp := primitive.ObjectIDFromHex(id)
	if errp != nil {
		return nil, errp
	}
	errd := s.cfrepo.Select(oid).Decode(&ce)
	if errd != nil {
		return nil, errd
	}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "Type", Value: libpik.Ifelse(cd.Type, ce.Type)},
		{Key: "lastname", Value: libpik.Ifelse(cd.Date, ce.Date)},
		{Key: "Category", Value: libpik.Ifelse(cd.Category, ce.Category)},
		{Key: "Total", Value: libpik.Ifelse(cd.Total, ce.Total)},
		{Key: "Description", Value: libpik.Ifelse(cd.Description, ce.Description)},
		{Key: "UpdatedBy", Value: "admin update"},
		{Key: "UpdatedDate", Value: time.Now()},
	}}}
	ce.Type = libpik.Ifelse(cd.Type, ce.Type).(string)
	ce.Date = libpik.Ifelse(cd.Date, ce.Date).(time.Time)
	ce.Category = libpik.Ifelse(cd.Category, ce.Category).(string)
	ce.Total = libpik.Ifelse(cd.Total, ce.Total).(int)
	ce.Description = libpik.Ifelse(cd.Description, ce.Description).(string)
	ce.UpdatedBy = "admin update"
	ce.UpdatedDate = time.Now()
	_, err := s.cfrepo.Update(oid, update)
	return ce, err
}

func (s *cfsvc) Delete(id string) (any, error) {
	oid, errp := primitive.ObjectIDFromHex(id)
	if errp != nil {
		return nil, errp
	}
	errd := s.cfrepo.Select(oid).Decode(&ce)
	if errd != nil {
		return nil, errd
	}
	err := s.cfrepo.Delete(oid)
	return ce, err
}
