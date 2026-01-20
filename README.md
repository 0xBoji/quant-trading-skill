# Quant Trading Skill

> High-performance quantitative trading knowledge base with BM25 search - written in Go

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)](https://go.dev/dl/)
[![Knowledge Base](https://img.shields.io/badge/entries-122-green.svg)](./data/)

A blazing-fast CLI tool providing instant access to curated quantitative trading knowledge: strategies, indicators, risk management, data sources, and common pitfalls.

## 🚀 Quick Start

### One-Line Installation

```bash
curl -fsSL https://raw.githubusercontent.com/0xboji/quant-trading-skill/main/install.sh | bash
```

### Manual Build

```bash
git clone https://github.com/0xboji/quant-trading-skill.git
cd quant-trading-skill
go build -o qts ./cmd/qts
./qts "order flow crypto"
```

### Basic Usage

```bash
# Auto-detect domain
qts "order flow imbalance crypto"

# Search specific domain
qts "stop loss kelly" -d risk

# Get more results
qts "rsi bollinger moving average" -d indicator -n 5

# Available domains: strategy, indicator, risk, data, anti-pattern
```

## 📊 Knowledge Base

| Domain | Entries | Description |
|--------|---------|-------------|
| **Strategies** | 24 | HFT to long-term: OFI, Hawkes, Kalman, Pairs, Market Making, ML/RL |
| **Indicators** | 23 | Technical indicators: EMA, RSI, MACD, ATR, Volume Profile, TCI |
| **Risk Management** | 24 | Position sizing, stops, VaR, Kelly, hedging, circuit breakers |
| **Data Sources** | 24 | Tick data, order books, options, on-chain, news, costs |
| **Anti-Patterns** | 27 | Common mistakes: overfitting, look-ahead bias, slippage |
| **TOTAL** | **122** | Curated quant trading knowledge |

## 🎯 Features

- **⚡ Blazing Fast**: Go implementation, <1ms search on 122 entries
- **🔍 BM25 Search**: Industry-standard TF-IDF ranking
- **🤖 Auto-Detection**: Intelligently routes queries to best domain
- **🎨 Colored Output**: Beautiful terminal formatting
- **📦 Zero Config**: Single binary, works out of the box
- **🚀 Cross-Platform**: Linux, macOS, Windows (amd64, arm64)

## 📖 Examples

### Search Strategies

```bash
$ qts "mean reversion crypto" -d strategy

====================================================================================================
QUERY: mean reversion crypto
====================================================================================================

Domain: strategy
Found 2 results:

1. Statistical Arbitrage (Pairs)
   Category: Mean Reversion
   Time Horizon: Medium-term (hours-days)
   Complexity: Medium
   Best For: Correlated assets, market-neutral strategies, hedged exposure

2. Mean Reversion (RSI/Bollinger)
   Category: Oscillator
   Time Horizon: Short-medium (minutes-hours-days)
   Complexity: Low
   Best For: Range-bound markets, high-volatility instruments, short-term reversals
```

### Search Indicators

```bash
$ qts "exponential moving average" -d indicator -n 2

====================================================================================================
QUERY: exponential moving average
====================================================================================================

Domain: indicator
Found 2 results:

1. EMA (Exponential Moving Average)
   Category: Trend
   Formula/Description: EMA_t = α·Price_t + (1-α)·EMA_{t-1}, α = 2/(N+1)
   Parameters: Period N (12, 26, 50, 200), α (smoothing factor)
   Best For: Trend detection, crossovers, smoothing price series

2. MACD (Moving Average Convergence Divergence)
   Category: Momentum Trend
   Formula/Description: MACD = EMA(12) - EMA(26), Signal = EMA(9) of MACD
   Parameters: Fast (12), Slow (26), Signal (9)
```

### Check for Mistakes

```bash
$ qts "backtesting overfitting" -d anti-pattern

====================================================================================================
QUERY: backtesting overfitting
====================================================================================================

Domain: anti-pattern
Found 3 results:

1. ⚠️  Overfitting on In-Sample Data
   Category: Strategy Design
   Severity: CRITICAL
   Don't: Optimize 10+ parameters, test only in-sample, maximize historical Sharpe
   Do: Use walk-forward analysis, out-of-sample testing, limit parameters (<5)

2. ⚠️  Look-Ahead Bias
   Category: Strategy Design
   Severity: CRITICAL
   Don't: Use 'close' prices before they're known, reindex data without timestamps
   Do: Ensure all signals use only data available at signal time, timestamp everything
```

## 🏗️ Architecture

```
quant-trading-skill/
├── cmd/
│   └── qts/                # CLI application
│       └── main.go         # Cobra CLI with colored output
├── internal/
│   ├── bm25/              # BM25 ranking algorithm
│   │   └── bm25.go        # Core search engine
│   └── search/            # Search logic
│       └── search.go      # Domain detection, CSV loading
├── data/                  # Knowledge base (122 entries)
│   ├── strategies.csv
│   ├── indicators.csv
│   ├── risk-management.csv
│   ├── data-sources.csv
│   └── anti-patterns.csv
├── go.mod                 # Go dependencies
└── README.md              # This file
```

## 🔧 Build from Source

### Prerequisites

- Go 1.21 or later
- Git

### Build

```bash
git clone https://github.com/0xboji/quant-trading-skill.git
cd quant-trading-skill
go mod download
go build -o qts ./cmd/qts
```

### Run Tests

```bash
go test -v ./...
```

### Cross-Compile

```bash
# Linux amd64
GOOS=linux GOARCH=amd64 go build -o qts-linux-amd64 ./cmd/qts

# macOS arm64 (M1/M2)
GOOS=darwin GOARCH=arm64 go build -o qts-darwin-arm64 ./cmd/qts

# Windows amd64
GOOS=windows GOARCH=amd64 go build -o qts-windows-amd64.exe ./cmd/qts
```

## 🚀 Advanced Usage

### Environment Variables

```bash
# Custom data directory
export QTS_DATA_DIR="/path/to/data"
qts "query"

# Use --data-dir flag
qts "query" --data-dir /custom/path/data
```

### Integration with Go Projects

```go
import (
    "github.com/0xboji/quant-trading-skill/internal/search"
)

func main() {
    result, err := search.Search("./data", "order flow crypto", "", 3)
    if err != nil {
        log.Fatal(err)
    }
    
    for _, r := range result.Results {
        fmt.Printf("%s: %s\n", r["Strategy Name"], r["Best For"])
    }
}
```

## 🎓 Use Cases

- **Strategy Research**: Find strategies matching your market/capital/experience
- **Technical Analysis**: Discover compatible indicators for signal generation
- **Risk Planning**: Identify required risk controls before deployment
- **Data Validation**: Verify you can access necessary data sources
- **Error Prevention**: Check for common pitfalls specific to your strategy type
- **Education**: Learn microstructure, market making, execution algorithms

## 📊 Performance

- **Search Speed**: <1ms (Go native, BM25 algorithm)
- **Memory**: ~5MB RAM (includes loaded CSVs)
- **Binary Size**: ~8MB (statically compiled)
- **Startup Time**: <10ms (instant CLI)

## 🤝 Contributing

Contributions welcome! To add knowledge:

1. **Add entries**: Edit CSV files in `data/`
2. **Update keywords**: Ensure searchability
3. **Test**: Run `qts "your query"`
4. **Submit PR**: With description of additions

## 📜 License

MIT License - see LICENSE file

Knowledge base: Curated from public domain trading knowledge

## 🙏 Credits

Inspired by:
- Market microstructure literature (Hautsch, Hasbrouck, Cartea)
- Production HFT systems (GO-Microstructures project)
- Risk management frameworks
- Practical trading experience

Built with:
- [cobra](https://github.com/spf13/cobra) - CLI framework
- [color](https://github.com/fatih/color) - Terminal colors

---

**Author**: 0xboji  
**Repository**: [github.com/0xboji/quant-trading-skill](https://github.com/0xboji/quant-trading-skill)  
**Version**: 1.0.0  
**Last Updated**: 2026-01-20

⭐ Star this repo if you find it useful!
