package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/domin191013/Go-Hitbtc-With-Cache/config"
	"github.com/domin191013/Go-Hitbtc-With-Cache/hitbtc"
	"github.com/domin191013/Go-Hitbtc-With-Cache/types"
)

type HandleRequests struct {
	HitWrapper *hitbtc.HitBtcWrapper
}

func writeResponse(w http.ResponseWriter, code int, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (h *HandleRequests) subscribeMarketFeeds() error {
	err := h.HitWrapper.FeedConnect()
	if err != nil {
		return err
	}
	return nil
}

func (h *HandleRequests) GetAllCurrencies() ([]*types.Ticker, error) {
	data, err := h.HitWrapper.GetCurrenciesFromCache()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (h *HandleRequests) handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/currency/all", h.handleAllCurrency).Methods("GET")
	myRouter.HandleFunc("/currency/{symbol}", h.handleCurrencyBySymbol).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+config.GetPort(), myRouter))
}

func Start() {
	h := &HandleRequests{
		HitWrapper: hitbtc.NewHitBtcV2Wrapper(),
	}
	err := h.HitWrapper.CacheAllSymbols()
	if err != nil {
		fmt.Println(err)
	}
	err = h.HitWrapper.CacheFullName()
	if err != nil {
		fmt.Println(err)
	}
	err = h.subscribeMarketFeeds()
	if err != nil {
		fmt.Println(err)
	}

	h.handleRequests()
}
