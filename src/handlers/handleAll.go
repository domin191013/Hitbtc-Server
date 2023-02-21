package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/domin191013/Go-Hitbtc-With-Cache/types"
)

func (h *HandleRequests) handleAllCurrency(w http.ResponseWriter, req *http.Request) {
	currencies, err := h.GetAllCurrencies()
	if err != nil {
		errorBody, _ := json.Marshal(&types.ErrorResponse{Error: err.Error()})
		writeResponse(w, http.StatusInternalServerError, errorBody)
		return
	}
	if len(currencies) == 0 {
		errorBody, _ := json.Marshal(&types.ErrorResponse{Error: "No data Found"})
		writeResponse(w, http.StatusNotFound, errorBody)
		return
	}

	var response types.Response
	response.Currencies = currencies
	currenciesJSON, err := json.Marshal(response)
	if err != nil {
		errorBody, _ := json.Marshal(&types.ErrorResponse{Error: err.Error()})
		writeResponse(w, http.StatusInternalServerError, errorBody)
		return
	}
	writeResponse(w, http.StatusOK, currenciesJSON)
}
