package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luvsangombos/banking/dto"
	"github.com/luvsangombos/banking/service"
)

type AccountHandler struct {
	service service.AccountService
}

func (a AccountHandler) new(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	cusomerId := vars["customer_id"]
	var request dto.NewAccountResquest

	err := json.NewDecoder(req.Body).Decode(&request)
	request.CustomerId = cusomerId

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		account, appError := a.service.NewAccount(request)

		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}

func (a AccountHandler) makeTransaction(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]

	var request dto.TransactionRequest

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.AccountId = accountId
		request.CustomerId = customerId

		account, appError := a.service.MakeTransaction(request)

		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, account)
		}
	}

}
