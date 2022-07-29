package service

import (
	"context"
	"time"

	"github.com/Trickster-ID/GO-CashFlow/model/entity"
	"github.com/Trickster-ID/GO-CashFlow/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type CashFlowSvc interface {
	Save(cio entity.Cashinout) (*mongo.InsertOneResult, error)
	GetAll() ([]entity.Cashinout, error)
}

type cfsvc struct {
	cfrepo repository.CashFlowRepo
}

func NewCashFlowSvc(Cfrepo repository.CashFlowRepo) CashFlowSvc {
	return &cfsvc{
		cfrepo: Cfrepo,
	}
}

func (s *cfsvc) Save(cio entity.Cashinout) (*mongo.InsertOneResult, error) {
	cio.CreatedBy = "admin"
	cio.CreatedDate = time.Now()
	cio.UpdatedBy = "admin"
	cio.UpdatedDate = time.Now()
	return s.cfrepo.Save(cio)
}

func (s *cfsvc) GetAll() ([]entity.Cashinout, error) {
	var cios []entity.Cashinout
	cursor, err := s.cfrepo.SelectAll()
	if err != nil {
		return []entity.Cashinout{}, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var cio entity.Cashinout
		cursor.Decode(&cio)
		cios = append(cios, cio)
	}
	if err := cursor.Err(); err != nil {
		return []entity.Cashinout{}, err
	}
	return cios, err
}
