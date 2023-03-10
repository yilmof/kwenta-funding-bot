package dydx

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type MarketsResponse struct {
	Markets map[string]Market `json:"markets"`
}

type Market struct {
	Market                           string    `json:"market"`
	BaseAsset                        string    `json:"baseAsset"`
	QuoteAsset                       string    `json:"quoteAsset"`
	StepSize                         string    `json:"stepSize"`
	TickSize                         string    `json:"tickSize"`
	IndexPrice                       string    `json:"indexPrice"`
	OraclePrice                      string    `json:"oraclePrice"`
	PriceChange24H                   string    `json:"priceChange24H"`
	NextFundingRate                  string    `json:"nextFundingRate"`
	MinOrderSize                     string    `json:"minOrderSize"`
	Type                             string    `json:"type"`
	InitialMarginFraction            string    `json:"initialMarginFraction"`
	MaintenanceMarginFraction        string    `json:"maintenanceMarginFraction"`
	BaselinePositionSize             string    `json:"baselinePositionSize"`
	IncrementalPositionSize          string    `json:"incrementalPositionSize"`
	IncrementalInitialMarginFraction string    `json:"incrementalInitialMarginFraction"`
	Volume24H                        string    `json:"volume24H"`
	Trades24H                        string    `json:"trades24H"`
	OpenInterest                     string    `json:"openInterest"`
	MaxPositionSize                  string    `json:"maxPositionSize"`
	AssetResolution                  string    `json:"assetResolution"`
	SyntheticAssetID                 string    `json:"syntheticAssetId"`
	Status                           string    `json:"status"`
	NextFundingAt                    time.Time `json:"nextFundingAt"`
}

func GetFunding(symbol string) (float64, error) {
	url := fmt.Sprintf("https://api.dydx.exchange/v3/markets?market=%s-USD", strings.ToUpper(symbol))
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	m := &MarketsResponse{}
	if err := json.Unmarshal(body, m); err != nil {
		return 0, err
	}

	rate, _ := strconv.ParseFloat(m.Markets[fmt.Sprintf("%s-USD", strings.ToUpper(symbol))].NextFundingRate, 32)

	return rate * 100, nil
}
