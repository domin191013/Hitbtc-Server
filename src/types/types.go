package types

type Currency struct {
	Id                 string `json:"id"`
	FullName           string `json:"fullName"`
	Crypto             bool   `json:"crypto"`
	PayinEnabled       bool   `json:"payinEnabled"`
	PayinPaymentId     bool   `json:"payinPaymentId"`
	PayinConfirmations uint   `json:"payinConfirmations"`
	PayoutEnabled      bool   `json:"payoutEnabled"`
	PayoutIsPaymentId  bool   `json:"payoutIsPaymentId"`
	TransferEnabled    bool   `json:"transferEnabled"`
}

type Symbol struct {
	Id                   string  `json:"id"`
	BaseCurrency         string  `json:"baseCurrency"`
	QuoteCurrency        string  `json:"quoteCurrency"`
	QuantityIncrement    float64 `json:"quantityIncrement,string"`
	TickSize             float64 `json:"tickSize,string"`
	TakeLiquidityRate    float64 `json:"takeLiquidityRate,string"`
	ProvideLiquidityRate float64 `json:"provideLiquidityRate,string"`
	FeeCurrency          string  `json:"feeCurrency"`
}

type Ticker struct {
	ID          string  `json:"id"`
	FullName    string  `json:"fullName"`
	Ask         float64 `json:"ask,string"`
	Bid         float64 `json:"bid,string"`
	Last        float64 `json:"last,string"`
	Open        float64 `json:"open,string"`
	Low         float64 `json:"low,string,omitempty"`
	High        float64 `json:"high,string,omitempty"`
	Symbol      string  `json:"symbol"`
	FeeCurrency string  `json:"feeCurrency"`
}

type Tickers []Ticker

type Response struct {
	Currencies []*Ticker `json:"currencies"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
