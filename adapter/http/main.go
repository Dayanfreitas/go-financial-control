package http

import (
	"net/http"

	"github.com/Dayanfreitas/go-financial-control/adapter/http/actuator"
	"github.com/Dayanfreitas/go-financial-control/adapter/http/transaction"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateTransactions)

	http.HandleFunc("/health", actuator.Health)
	http.Handle("/metric", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
