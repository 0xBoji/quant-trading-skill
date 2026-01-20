# QuantPro

> Professional quantitative trading toolkit with intelligent knowledge base

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)](https://go.dev/dl/)
[![Knowledge Base](https://img.shields.io/badge/entries-122-green.svg)](./data/)

A blazing-fast CLI tool for quantitative trading research. Search strategies, indicators, risk management techniques, and avoid common pitfalls with intelligent BM25 search.

## 🚀 Quick Start

### One-Line Installation

```bash
curl -fsSL https://raw.githubusercontent.com/0xboji/quant-trading-skill/main/install.sh | bash
```

### Initialize in Your Project

```bash
cd /path/to/your/quant-project
quantpro init --ai antigravity
```

This creates:
```
.agent/
└── workflows/
    └── use-quant-skill.md
.shared/
└── quant-trading-pro/
    ├── data/           # 5 CSV files (122 entries)
    └── SKILL.md       # Documentation
```

### Search Knowledge Base

```bash
# Auto-detect domain
quantpro search "order flow crypto"

# Search specific domain
quantpro search "stop loss kelly" -d risk

# Get more results
quantpro search "rsi bollinger" -d indicator -n 5
```

## 📊 Knowledge Base

| Domain | Entries | Description |
|--------|---------|-------------|
| **strategies** | 24 | HFT to long-term: OFI, Hawkes, Kalman, Pairs, Market Making, ML/RL |
| **indicators** | 23 | Technical indicators: EMA, RSI, MACD, ATR, Volume Profile, TCI |
| **risk** | 24 | Position sizing, stops, VaR, Kelly, hedging, circuit breakers |
| **data** | 24 | Tick data, order books, options, on-chain, news, costs |
| **anti-pattern** | 27 | Common mistakes: overfitting, look-ahead bias, slippage |
| **TOTAL** | **122** | Curated quant trading knowledge |

## 🎯 Features

- **⚡ Blazing Fast**: <1ms search on 122 entries (Go native)
- **🔍 BM25 Search**: Industry-standard TF-IDF ranking
- **🤖 Auto-Detection**: Intelligent domain routing
- **🎨 Beautiful CLI**: Colored output, formatted results
- **📦 Easy Setup**: `init` command creates project structure
- **🚀 Cross-Platform**: Linux, macOS, Windows (amd64, arm64)

## 📖 Usage

### Initialize Project

```bash
# Create .agent and .shared directories with skill
quantpro init --ai antigravity

# Custom target directory
quantpro init --ai custom-agent --dir /path/to/project
```

### Search Commands

```bash
# Basic search (auto-detect domain)
quantpro search "mean reversion crypto"

# Domain-specific search
quantpro search "exponential moving average" -d indicator

# Get more results
quantpro search "position sizing" -d risk -n 5

# Custom data directory
quantpro search "query" --data-dir /path/to/data
```

### Available Domains

- `strategy` - Trading strategies (OFI, Hawkes, Pairs, Market Making, etc.)
- `indicator` - Technical indicators (EMA, RSI, MACD, ATR, etc.)
- `risk` - Risk management (Kelly, VaR, Position Sizing, Stops, etc.)
- `data` - Data sources (Tick data, Order books, Options, etc.)
- `anti-pattern` - Common mistakes to avoid

## 💡 Examples

### Search Strategies

```bash
$ quantpro search "high frequency order flow" -d strategy

====================================================================================================
QUERY: high frequency order flow
====================================================================================================

Domain: strategy
Found 3 results:

1. Order Flow Imbalance (OFI)
   Category: Microstructure
   Time Horizon: Intraday (1ms-60s)
   Complexity: Medium
   Best For: High-frequency futures, liquid crypto pairs, institutional execution

2. Dual-Scale Hawkes Process
   Category: Point Process
   Time Horizon: Ultra-short (100ms-10s)
   Complexity: High
   Best For: Detecting sweeps, momentum ignition, liquidity crises
```

### Check for Mistakes

```bash
$ quantpro search "backtesting overfitting" -d anti-pattern

====================================================================================================
QUERY: backtesting overfitting
====================================================================================================

Domain: anti-pattern
Found 2 results:

1. Overfitting on In-Sample Data
   Category: Strategy Design
   Severity: CRITICAL
   Don't: Optimize 10+ parameters, test only in-sample, maximize historical Sharpe
   Do: Use walk-forward analysis, out-of-sample testing, limit parameters (<5)
```

## 🏗️ Project Structure

```
quant-trading-skill/
├── cmd/
│   └── quantpro/           # CLI application
│       └── main.go         # Cobra CLI + init command
├── internal/
│   ├── bm25/              # BM25 algorithm
│   │   └── bm25.go
│   └── search/            # Search logic
│       └── search.go
├── data/                  # Knowledge base (122 entries)
│   ├── strategies.csv
│   ├── indicators.csv
│   ├── risk-management.csv
│   ├── data-sources.csv
│   └── anti-patterns.csv
├── .github/workflows/     # CI/CD
│   └── build.yml          # Automated builds
├── go.mod
├── install.sh             # One-line installer
└── README.md
```

## 🔧 Build from Source

### Prerequisites

- Go 1.21+
- Git

### Build

```bash
git clone https://github.com/0xboji/quant-trading-skill.git
cd quant-trading-skill
go build -o quantpro ./cmd/quantpro
./quantpro --version
```

### Cross-Compile

```bash
# Linux amd64
GOOS=linux GOARCH=amd64 go build -o quantpro-linux-amd64 ./cmd/quantpro

# macOS arm64 (M1/M2)
GOOS=darwin GOARCH=arm64 go build -o quantpro-darwin-arm64 ./cmd/quantpro

# Windows
GOOS=windows GOARCH=amd64 go build -o quantpro-windows-amd64.exe ./cmd/quantpro
```

## 🚀 Integration

### As Go Package

```go
import "github.com/0xboji/quant-trading-skill/internal/search"

result, err := search.Search("./data", "order flow crypto", "", 3)
if err != nil {
    log.Fatal(err)
}

for _, r := range result.Results {
    fmt.Printf("%s: %s\n", r["Strategy Name"], r["Best For"])
}
```

### In Your Project

After running `quantpro init --ai antigravity`:

```bash
# Data available at
.shared/quant-trading-pro/data/*.csv

# Workflow documentation at
.agent/workflows/use-quant-skill.md

# Full documentation at
.shared/quant-trading-pro/SKILL.md
```

## 📊 Performance

- **Search Speed**: <1ms (Go native, BM25)
- **Memory**: ~5MB RAM
- **Binary Size**: ~8MB (static compilation)
- **Startup**: <10ms

## 🎓 Use Cases

- **Strategy Research**: Find strategies for your market/capital
- **Technical Analysis**: Discover compatible indicators
- **Risk Planning**: Identify required controls
- **Data Validation**: Verify data source access
- **Error Prevention**: Check for common pitfalls
- **Education**: Learn microstructure and HFT

## 🤝 Contributing

Contributions welcome!

1. Add entries to CSV files in `data/`
2. Update keywords for searchability
3. Test: `quantpro search "your query"`
4. Submit PR

## 📜 License

MIT License - See LICENSE file

## 🙏 Credits

Knowledge base inspired by:
- Market microstructure literature
- Production HFT systems
- Risk management frameworks  
- Practical trading experience

Built with:
- [cobra](https://github.com/spf13/cobra) - CLI framework
- [color](https://github.com/fatih/color) - Terminal colors

---

**Author**: 0xboji  
**Repository**: [github.com/0xboji/quant-trading-skill](https://github.com/0xboji/quant-trading-skill)  
**Version**: 1.0.0

⭐ Star this repo if you find it useful!
