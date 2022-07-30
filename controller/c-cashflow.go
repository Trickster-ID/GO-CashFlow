package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Trickster-ID/GO-CashFlow/model/entity"
	"github.com/Trickster-ID/GO-CashFlow/service"
	"github.com/Trickster-ID/go-libpik"
	"github.com/gorilla/mux"
)

type CashFlowCtrl interface {
	Post(response http.ResponseWriter, request *http.Request)
	GetAll(response http.ResponseWriter, request *http.Request)
	Get(response http.ResponseWriter, request *http.Request)
	Put(response http.ResponseWriter, request *http.Request)
	Delete(response http.ResponseWriter, request *http.Request)
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
		return
	}
	json.NewEncoder(response).Encode(libpik.BSuccessResponse(res))
}

func (c *cfctrl) GetAll(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	res, err := c.cfsvc.GetAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(libpik.BSuccessResponse(res))
}

func (c *cfctrl) Get(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	id := mux.Vars(request)["id"]
	res, err := c.cfsvc.Get(id)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(libpik.BSuccessResponse(res))
}

func (c *cfctrl) Put(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	json.NewDecoder(request.Body).Decode(&cashInOut)
	id := mux.Vars(request)["id"]
	res, err := c.cfsvc.Update(id, cashInOut)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		//http.Error(response, libpik.BErrorResponse("error when update data", err.Error()),400)
		return
	}
	json.NewEncoder(response).Encode(libpik.BSuccessResponse(res))
}

func (c *cfctrl) Delete(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	id := mux.Vars(request)["id"]
	res, err := c.cfsvc.Delete(id)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(libpik.BSuccessResponse(res))
}
