package search

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/0xboji/quant-trading-skill/internal/bm25"
)

// Config defines search configuration for a domain
type Config struct {
	File       string
	SearchCols []string
	OutputCols []string
}

// DomainConfigs maps domain names to their configurations
var DomainConfigs = map[string]Config{
	"strategy": {
		File:       "strategies.csv",
		SearchCols: []string{"Strategy Name", "Category", "Keywords", "Best For", "Data Requirements"},
		OutputCols: []string{"Strategy Name", "Category", "Keywords", "Data Requirements", "Time Horizon", "Best For", "Complexity", "Key Parameters", "Performance Characteristics", "Market Conditions", "Capital Requirements", "Avoid For"},
	},
	"indicator": {
		File:       "indicators.csv",
		SearchCols: []string{"Indicator Name", "Category", "Keywords", "Best For"},
		OutputCols: []string{"Indicator Name", "Category", "Keywords", "Formula/Description", "Time-Domain", "Best For", "Parameters", "Interpretation", "Limitations", "Combine With", "Avoid For"},
	},
	"risk": {
		File:       "risk-management.csv",
		SearchCols: []string{"Risk Control", "Category", "Keywords", "Description", "Best For"},
		OutputCols: []string{"Risk Control", "Category", "Keywords", "Description", "Parameters", "Best For", "Implementation", "Advantages", "Disadvantages", "Critical For"},
	},
	"data": {
		File:       "data-sources.csv",
		SearchCols: []string{"Data Type", "Source", "Keywords", "Description", "Best For"},
		OutputCols: []string{"Data Type", "Source", "Keywords", "Description", "Format", "Frequency", "Best For", "Requirements", "Typical Cost"},
	},
	"anti-pattern": {
		File:       "anti-patterns.csv",
		SearchCols: []string{"Category", "Issue", "Keywords", "Description"},
		OutputCols: []string{"Category", "Issue", "Keywords", "Description", "Do", "Don't", "Severity", "Platform"},
	},
}

// DomainKeywords for auto-detection
var DomainKeywords = map[string][]string{
	"strategy":     {"strategy", "trading", "algorithm", "arbitrage", "ofi", "hawkes", "kalman", "momentum", "mean-reversion", "pairs", "market-making", "statistical", "execution", "vwap", "ml", "backtest"},
	"indicator":    {"indicator", "ema", "sma", "rsi", "macd", "bollinger", "atr", "stochastic", "adx", "vwap", "obv", "tci", "signal", "oscillator", "moving average"},
	"risk":         {"risk", "position", "sizing", "kelly", "stop", "loss", "drawdown", "var", "cvar", "leverage", "margin", "hedge", "portfolio", "limit", "exposure"},
	"data":         {"data", "tick", "order book", "l2", "ohlcv", "bars", "futures", "options", "on-chain", "news", "sentiment", "fundamental", "volume", "feed", "api"},
	"anti-pattern": {"mistake", "error", "avoid", "don't", "anti-pattern", "pitfall", "bias", "overfitting", "look-ahead", "survivorship", "slippage", "bug", "wrong"},
}

// Result represents a search result
type Result struct {
	Domain  string
	Query   string
	File    string
	Count   int
	Results []map[string]string
}

// DetectDomain auto-detects the most relevant domain from query
func DetectDomain(query string) string {
	queryLower := strings.ToLower(query)
	scores := make(map[string]int)

	for domain, keywords := range DomainKeywords {
		score := 0
		for _, kw := range keywords {
			if strings.Contains(queryLower, kw) {
				score++
			}
		}
		scores[domain] = score
	}

	// Find best domain
	bestDomain := "strategy"
	bestScore := 0
	for domain, score := range scores {
		if score > bestScore {
			bestScore = score
			bestDomain = domain
		}
	}

	if bestScore == 0 {
		return "strategy"
	}
	return bestDomain
}

// Search performs BM25 search on specified domain
func Search(dataDir, query, domain string, maxResults int) (*Result, error) {
	if domain == "" {
		domain = DetectDomain(query)
	}

	config, ok := DomainConfigs[domain]
	if !ok {
		return nil, fmt.Errorf("unknown domain: %s", domain)
	}

	filepath := filepath.Join(dataDir, config.File)

	// Load CSV
	data, err := loadCSV(filepath)
	if err != nil {
		return nil, err
	}

	// Build documents from search columns
	documents := make([]string, len(data))
	for i, row := range data {
		parts := make([]string, 0, len(config.SearchCols))
		for _, col := range config.SearchCols {
			if val, ok := row[col]; ok {
				parts = append(parts, val)
			}
		}
		documents[i] = strings.Join(parts, " ")
	}

	// BM25 search
	engine := bm25.New(1.5, 0.75)
	engine.Fit(documents)
	ranked := engine.Score(query)

	// Sort by score descending
	sort.Slice(ranked, func(i, j int) bool {
		return ranked[i].Score > ranked[j].Score
	})

	// Get top results with score > 0
	results := make([]map[string]string, 0, maxResults)
	for _, r := range ranked {
		if r.Score <= 0 {
			break
		}
		if len(results) >= maxResults {
			break
		}

		row := data[r.Index]
		result := make(map[string]string)
		for _, col := range config.OutputCols {
			if val, ok := row[col]; ok {
				result[col] = val
			}
		}
		results = append(results, result)
	}

	return &Result{
		Domain:  domain,
		Query:   query,
		File:    config.File,
		Count:   len(results),
		Results: results,
	}, nil
}

// loadCSV loads a CSV file and returns a slice of maps
func loadCSV(filepath string) ([]map[string]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) < 2 {
		return nil, fmt.Errorf("CSV file too short")
	}

	// First row is header
	headers := records[0]
	data := make([]map[string]string, 0, len(records)-1)

	for _, record := range records[1:] {
		row := make(map[string]string)
		for i, header := range headers {
			if i < len(record) {
				row[header] = record[i]
			}
		}
		data = append(data, row)
	}

	return data, nil
}
