// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package asset

import (
	"log/slog"
	"net/http"
	"time"
)

// TiingoMeta is the response from the meta endpoint.
// https://www.tiingo.com/documentation/end-of-day
type TiingoMeta struct {
	// Ticker related to the asset.
	Ticker string `json:"ticker"`

	// Name is the full name of the asset.
	Name string `json:"name"`

	// ExchangeCode is the exchange where the asset is listed on.
	ExchangeCode string `json:"exchangeCode"`

	// Description is the description of the asset.
	Description string `json:"description"`

	// StartDate is the earliest date for the asset data.
	StartDate time.Time `json:"startDate"`

	// EndDate is the latest date for the asset data.
	EndDate time.Time `json:"endDate"`
}

// TiingoEndOfDay is the repose from the end-of-day endpoint.
// https://www.tiingo.com/documentation/end-of-day
type TiingoEndOfDay struct {
	// Date is the date this data pertains to.
	Date time.Time `json:"date"`

	// Open is the opening price.
	Open float64 `json:"open"`

	// High is the highest price.
	High float64 `json:"high"`

	// Low is the lowest price.
	Low float64 `json:"low"`

	// Close is the closing price.
	Close float64 `json:"close"`

	// Volume is the total volume.
	Volume float64 `json:"volume"`

	// AdjOpen is the adjusted opening price.
	AdjOpen float64 `json:"adjOpen"`

	// AdjHigh is the adjusted highest price.
	AdjHigh float64 `json:"adjHigh"`

	// AdjLow is the adjusted lowest price.
	AdjLow float64 `json:"adjLow"`

	// AdjClose is the adjusted closing price.
	AdjClose float64 `json:"adjClose"`

	// AdjVolume is the adjusted total volume.
	AdjVolume float64 `json:"adjVolume"`

	// Dividend is the dividend paid out.
	Dividend float64 `json:"divCash"`

	// Split to adjust values after a split.
	Split float64 `json:"splitFactor"`
}

// ToSnapshot converts the Tiingo end-of-day to a snapshot.
func (e *TiingoEndOfDay) ToSnapshot() *Snapshot { _ = "STUB: not implemented"; return nil }

// TiingoRepository provides access to financial market data, retrieving
// asset snapshots, by interacting with the Tiingo Stock & Financial
// Markets API. To use this repository, you'll need a valid API key
// from https://www.tiingo.com.
type TiingoRepository struct {
	Repository

	// apiKey is the Tiingo API key.
	apiKey string

	// Client is the HTTP client.
	client *http.Client

	// BaseURL is the Tiingo API URL.
	BaseURL string

	// Logger is the slog logger instance.
	Logger *slog.Logger
}

// NewTiingoRepository initializes a file system repository with
// the given API key.
func NewTiingoRepository(apiKey string) *TiingoRepository { _ = "STUB: not implemented"; return nil }

// Assets returns the names of all assets in the repository.
func (*TiingoRepository) Assets() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// Get attempts to return a channel of snapshots for the asset with the given name.
func (r *TiingoRepository) Get(name string) (<-chan *Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetSince attempts to return a channel of snapshots for the asset with the given name since the given date.
func (r *TiingoRepository) GetSince(name string, date time.Time) (<-chan *Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LastDate returns the date of the last snapshot for the asset with the given name.
func (r *TiingoRepository) LastDate(name string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// Append adds the given snapshows to the asset with the given name.
func (*TiingoRepository) Append(_ string, _ <-chan *Snapshot) error {
	_ = "STUB: not implemented"
	return nil
}
