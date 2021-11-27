package transaction

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Dayanfreitas/go-financial-control/model/transaction"
	"github.com/Dayanfreitas/go-financial-control/util"
)

type Transactions []transaction.Transaction

func GetTransactions(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	rw.Header().Set("Content-type", "application/json")

	transactions := Transactions{
		transaction.Transaction{
			Title:     "Salário",
			Amount:    1000,
			Type:      0,
			CreatedAt: util.StringToTime("2021-11-28T11:00:00"),
		},
	}

	json.NewEncoder(rw).Encode(transactions)
}

func CreateTransactions(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	transaction := transaction.Transaction{}
	var body, _ = ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &transaction)

	log.Print(transaction)
}
