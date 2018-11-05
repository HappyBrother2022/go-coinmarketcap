// Package types of both request and response for API
package types

// Status structure
type Status struct {
	Timestamp    string  `json:"timestamp"`
	ErrorCode    int     `json:"error_code"`
	ErrorMessage *string `json:"error_message"`
	Elapsed      int     `json:"elapsed"`
	CreditCount  int     `json:"credit_count"`
}

// Response structure
type Response struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data"`
}

// CryptoInfoMap structure
type CryptoInfoMap struct {
	CryptoInfo map[string]*CryptoInfo `json:"data"`
}

// CryptoInfo structure
type CryptoInfo struct {
	ID       float64                `json:"id"`
	Name     string                 `json:"name"`
	Symbol   string                 `json:"symbol"`
	Category string                 `json:"category"`
	Slug     string                 `json:"slug"`
	Logo     string                 `json:"logo"`
	Tags     []string               `json:"tags"`
	Urls     map[string]interface{} `json:"urls"`
}

// CryptoMapList structure
type CryptoMapList struct {
	CryptoMap []*CryptoMap `json:"data"`
}

// CryptoMap structure
type CryptoMap struct {
	ID                  float64 `json:"id"`
	Name                string  `json:"name"`
	Symbol              string  `json:"symbol"`
	Slug                string  `json:"slug"`
	IsActive            int     `json:"is_active"`
	FirstHistoricalData string  `json:"first_historical_data"`
	LastHistoricalData  string  `json:"last_historical_data"`
}

// CryptoMarketList structure
type CryptoMarketList struct {
	CryptoMarket []*CryptoMarket `json:"data"`
}

// CryptoMarketMap structure
type CryptoMarketMap struct {
	CryptoMarket map[string]*CryptoMarket `json:"data"`
}

// CryptoMarket structure
type CryptoMarket struct {
	ID                float64           `json:"id"`
	Name              string            `json:"name"`
	Symbol            string            `json:"symbol"`
	Slug              string            `json:"slug"`
	CirculatingSupply float64           `json:"circulating_supply"`
	TotalSupply       float64           `json:"total_supply"`
	MaxSupply         float64           `json:"max_supply"`
	DateAdded         string            `json:"date_added"`
	NumMarketPairs    int               `json:"num_market_pairs"`
	CMCRank           int               `json:"cmc_rank"`
	LastUpdated       string            `json:"last_updated"`
	Quote             map[string]*Quote `json:"quote"`
}

// ExchangeInfoMap structure
type ExchangeInfoMap struct {
	ExchangeInfo map[string]*ExchangeInfo `json:"data"`
}

// ExchangeInfo options
type ExchangeInfo struct {
	ID   float64                `json:"id"`
	Name string                 `json:"name"`
	Slug string                 `json:"slug"`
	Logo string                 `json:"logo"`
	Urls map[string]interface{} `json:"urls"`
}

// ExchangeMarketPairs structure
type ExchangeMarketPairs struct {
	ID             float64        `json:"id"`
	Name           string         `json:"name"`
	Slug           string         `json:"slug"`
	NumMarketPairs int            `json:"num_market_pairs"`
	MarketPairs    []*MarketPairs `json:"market_pairs"`
}

// MarketPairs structure
type MarketPairs struct {
	MarketPair      string            `json:"market_pair"`
	MarketPairBase  *Currency         `json:"market_pair_base"`
	MarketPairQuote *Currency         `json:"market_pair_quote"`
	Quote           map[string]*Quote `json:"quote"`
}

// ExchangeMarketQuotes structure
type ExchangeMarketQuotes struct {
	MarketQuote map[string]*MarketQuote `json:"data"`
}

// ExchangeMarketList structure
type ExchangeMarketList struct {
	MarketQuote []*MarketQuote `json:"data"`
}

// MarketQuote structure
type MarketQuote struct {
	ID             float64           `json:"id"`
	Name           string            `json:"name"`
	Slug           string            `json:"slug"`
	NumMarketPairs int               `json:"num_market_pairs"`
	LastUpdated    string            `json:"last_updated"`
	Quote          map[string]*Quote `json:"quote"`
}

// Currency structure
type Currency struct {
	ID     int    `json:"currency_id"`
	Symbol string `json:"currency_symbol"`
	Type   string `json:"currency_type"`
}

// Quote structure
type Quote struct {
	Price            float64 `json:"price"`
	Volume24H        float64 `json:"volume_24h"`
	Volume7D         float64 `json:"volume_7d"`
	Volume30D        float64 `json:"volume_30d"`
	Volume24Hbase    float64 `json:"volume_24h_base"`
	Volume24Hquote   float64 `json:"volume_24h_quote"`
	PercentChange1H  float64 `json:"percent_change_1h"`
	PercentChange24H float64 `json:"percent_change_24h"`
	PercentChange7D  float64 `json:"percent_change_7d"`
	PercentChange30D float64 `json:"percent_change_30d"`
	MarketCap        float64 `json:"market_cap"`
	LastUpdated      string  `json:"last_updated"`
}

// Ticker struct
type Ticker struct {
	ID                int                     `json:"id"`
	Name              string                  `json:"name"`
	Symbol            string                  `json:"symbol"`
	Slug              string                  `json:"website_slug"`
	Rank              int                     `json:"rank"`
	CirculatingSupply float64                 `json:"circulating_supply"`
	TotalSupply       float64                 `json:"total_supply"`
	MaxSupply         float64                 `json:"max_supply"`
	Quotes            map[string]*TickerQuote `json:"quotes"`
	LastUpdated       int                     `json:"last_updated"`
}

// TickerQuote struct
type TickerQuote struct {
	Price            float64 `json:"price"`
	Volume24H        float64 `json:"volume_24h"`
	MarketCap        float64 `json:"market_cap"`
	PercentChange1H  float64 `json:"percent_change_1h"`
	PercentChange24H float64 `json:"percent_change_24h"`
	PercentChange7D  float64 `json:"percent_change_7d"`
}

// GlobalMarket struct
type GlobalMarket struct {
	ActiveCurrencies             int                           `json:"active_cryptocurrencies"`
	ActiveMarkets                int                           `json:"active_markets"`
	BitcoinPercentageOfMarketCap float64                       `json:"bitcoin_percentage_of_market_cap"`
	LastUpdated                  int                           `json:"last_updated"`
	Quotes                       map[string]*GlobalMarketQuote `json:"quotes"`
}

// GlobalMarketQuote struct
type GlobalMarketQuote struct {
	TotalMarketCap float64 `json:"total_market_cap"`
	TotalVolume24H float64 `json:"total_volume_24h"`
}

// TickerGraph struct
type TickerGraph struct {
	MarketCapByAvailableSupply [][]float64 `json:"market_cap_by_available_supply"`
	PriceBTC                   [][]float64 `json:"price_btc"`
	PriceUSD                   [][]float64 `json:"price_usd"`
	VolumeUSD                  [][]float64 `json:"volume_usd"`
}

// Market struct
type Market struct {
	Rank          int
	Exchange      string
	Pair          string
	VolumeUSD     float64
	Price         float64
	VolumePercent float64
	Updated       string
}

// MarketGraph struct
type MarketGraph struct {
	MarketCapByAvailableSupply [][]float64 `json:"market_cap_by_available_supply"`
	VolumeUSD                  [][]float64 `json:"volume_usd"`
}
