package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Trickster-ID/GO-CashFlow/model/entity"
	"github.com/Trickster-ID/GO-CashFlow/service"
	"github.com/Trickster-ID/go-libpik"
)

type CashFlowCtrl interface {
	Post(response http.ResponseWriter, request *http.Request)
	GetAll(response http.ResponseWriter, request *http.Request)
}

type cfctrl struct {
	cfsvc service.CashFlowSvc
}

func NewCashFlowCtrl(Cfsvc service.CashFlowSvc) CashFlowCtrl {
	return &cfctrl{
		cfsvc: Cfsvc,
	}
}

var cashInOut entity.Cashinout

func (c *cfctrl) Post(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	json.NewDecoder(request.Body).Decode(&cashInOut)
	fmt.Print("request body: ")
	fmt.Println(cashInOut)
	res, err := c.cfsvc.Save(cashInOut)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
	}
	json.NewEncoder(response).Encode(libpik.BSuccessResponse(res))
}

func (c *cfctrl) GetAll(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	res, err := c.cfsvc.GetAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
	}
	json.NewEncoder(response).Encode(libpik.BSuccessResponse(res))
}
