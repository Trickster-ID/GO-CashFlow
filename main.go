package main

import (
	"fmt"
	"net/http"

	"github.com/Trickster-ID/GO-CashFlow/config"
	"github.com/Trickster-ID/GO-CashFlow/controller"
	"github.com/Trickster-ID/GO-CashFlow/repository"
	"github.com/Trickster-ID/GO-CashFlow/service"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	db     *mongo.Client           = config.ResolveClientDB()
	cfrepo repository.CashFlowRepo = repository.NewCashFlowRepo(db)
	cfsvc  service.CashFlowSvc     = service.NewCashFlowSvc(cfrepo)
	cfctrl controller.CashFlowCtrl = controller.NewCashFlowCtrl(cfsvc)
)

func main() {
	defer config.CloseClientDB(db)

	r := mux.NewRouter()

	r.HandleFunc("/", cfctrl.POST).Methods("POST")
	fmt.Println("server is listening...")
	http.ListenAndServe(":8888", r)
}
