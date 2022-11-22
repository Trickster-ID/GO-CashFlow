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
	db                 *mongo.Client           = config.ResolveClientDB()
	cashFlowRepository repository.CashFlowRepo = repository.NewCashFlowRepo(db)
	cashFlowService    service.CashFlowSvc     = service.NewCashFlowSvc(cashFlowRepository)
	cashFlowController controller.CashFlowCtrl = controller.NewCashFlowCtrl(cashFlowService)
)

func main() {
	defer config.CloseClientDB(db)

	r := mux.NewRouter()

	r.HandleFunc("/", cashFlowController.Post).Methods("POST")
	r.HandleFunc("/", cashFlowController.GetAll).Methods("GET")
	r.HandleFunc("/{id}", cashFlowController.Get).Methods("GET")
	r.HandleFunc("/{id}", cashFlowController.Delete).Methods("DELETE")
	r.HandleFunc("/{id}", cashFlowController.Put).Methods("PUT")
	fmt.Println("server is listening...")
	http.ListenAndServe(":8888", r)
}
