package service

import (
	"github.com/Trickster-ID/GO-CashFlow/model/entity"
	"github.com/Trickster-ID/GO-CashFlow/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type CashFlowSvc interface {
	Save(cio entity.Cashinout) (*mongo.InsertOneResult, error)
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
	return s.cfrepo.Save(cio)
}
