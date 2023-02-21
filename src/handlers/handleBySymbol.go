package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/domin191013/Go-Hitbtc-With-Cache/types"
)

func (h *HandleRequests) handleCurrencyBySymbol(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key := vars["symbol"]
	var currenciesJSON []byte
	if h.HitWrapper.Contains(h.HitWrapper.AllSymbols, key) {
		currency, err := h.HitWrapper.GetMarketSummary(key)
		if err != nil {
			errorBody, _ := json.Marshal(&types.ErrorResponse{Error: err.Error()})
			writeResponse(w, http.StatusInternalServerError, errorBody)
			return
		}
		if currency == nil {
			errorBody, _ := json.Marshal(&types.ErrorResponse{Error: "No data Found"})
			writeResponse(w, http.StatusNotFound, errorBody)
			return
		}
		currenciesJSON, err = json.Marshal(currency)
		if err != nil {
			errorBody, _ := json.Marshal(&types.ErrorResponse{Error: err.Error()})
			writeResponse(w, http.StatusInternalServerError, errorBody)
			return
		}
	} else {
		errorBody, _ := json.Marshal(&types.ErrorResponse{Error: "Not a valid Symbol"})
		writeResponse(w, http.StatusNotFound, errorBody)
		return
	}

	writeResponse(w, http.StatusOK, currenciesJSON)
}
