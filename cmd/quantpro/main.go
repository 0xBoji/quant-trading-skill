package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/0xboji/quant-trading-skill/internal/search"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	domain     string
	maxResults int
	dataDir    string
	aiName     string
	targetDir  string
)

var rootCmd = &cobra.Command{
	Use:   "quantpro",
	Short: "QuantPro - Professional quantitative trading toolkit",
	Long: `QuantPro - Professional quantitative trading toolkit
	
Search quantitative trading knowledge base: strategies, indicators, 
risk management, data sources, and common pitfalls.`,
	Version: "1.0.0",
}

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search the quantitative trading knowledge base",
	Long: `Search across trading strategies, indicators, risk management,
data sources, and common pitfalls.

Examples:
  quantpro search "order flow crypto"
  quantpro search "stop loss kelly" -d risk
  quantpro search "rsi bollinger" -d indicator -n 5`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := strings.Join(args, " ")

		// Find data directory
		if dataDir == "" {
			dataDir = findDataDir()
		}

		result, err := search.Search(dataDir, query, domain, maxResults)
		if err != nil {
			color.Red("Error: %v", err)
			os.Exit(1)
		}

		printResults(result)
	},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize QuantPro skill in your project",
	Long: `Initialize QuantPro skill by creating .agent and .shared directories
with the quantitative trading knowledge base.

Examples:
  quantpro init --ai antigravity
  quantpro init --ai custom-agent --dir /path/to/project`,
	Run: func(cmd *cobra.Command, args []string) {
		if aiName == "" {
			color.Red("Error: --ai flag is required")
			fmt.Println("\nUsage: quantpro init --ai antigravity")
			os.Exit(1)
		}

		if targetDir == "" {
			targetDir = "."
		}

		if err := initializeSkill(targetDir, aiName); err != nil {
			color.Red("Error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	// Search command flags
	searchCmd.Flags().StringVarP(&domain, "domain", "d", "", "Domain to search (strategy, indicator, risk, data, anti-pattern)")
	searchCmd.Flags().IntVarP(&maxResults, "max-results", "n", 3, "Maximum number of results")
	searchCmd.Flags().StringVar(&dataDir, "data-dir", "", "Path to data directory")

	// Init command flags
	initCmd.Flags().StringVar(&aiName, "ai", "", "AI agent name (required)")
	initCmd.Flags().StringVar(&targetDir, "dir", ".", "Target directory (default: current)")
	initCmd.MarkFlagRequired("ai")

	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(initCmd)
}

func findDataDir() string {
	// Try current directory first
	if _, err := os.Stat("data"); err == nil {
		return "data"
	}

	// Try .shared/quant-trading-pro/data
	if _, err := os.Stat(".shared/quant-trading-pro/data"); err == nil {
		return ".shared/quant-trading-pro/data"
	}

	// Try executable directory
	ex, err := os.Executable()
	if err == nil {
		exPath := filepath.Dir(ex)
		dataPath := filepath.Join(exPath, "data")
		if _, err := os.Stat(dataPath); err == nil {
			return dataPath
		}
	}

	// Try home directory installation
	home, err := os.UserHomeDir()
	if err == nil {
		dataPath := filepath.Join(home, ".quant-trading-skill", "data")
		if _, err := os.Stat(dataPath); err == nil {
			return dataPath
		}
	}

	return "data"
}

func initializeSkill(targetDir, aiName string) error {
	cyan := color.New(color.FgCyan, color.Bold)
	green := color.New(color.FgGreen)

	cyan.Printf("\n%s\n", strings.Repeat("=", 100))
	cyan.Printf("QuantPro Initialization - AI Agent: %s\n", aiName)
	cyan.Printf("%s\n\n", strings.Repeat("=", 100))

	// Create .agent directory
	agentDir := filepath.Join(targetDir, ".agent", "workflows")
	if err := os.MkdirAll(agentDir, 0755); err != nil {
		return fmt.Errorf("failed to create .agent directory: %w", err)
	}
	green.Printf("✅ Created: %s\n", agentDir)

	// Create .shared directory
	sharedDir := filepath.Join(targetDir, ".shared", "quant-trading-pro")
	if err := os.MkdirAll(sharedDir, 0755); err != nil {
		return fmt.Errorf("failed to create .shared directory: %w", err)
	}
	green.Printf("✅ Created: %s\n", sharedDir)

	// Create data directory
	dataDestDir := filepath.Join(sharedDir, "data")
	if err := os.MkdirAll(dataDestDir, 0755); err != nil {
		return fmt.Errorf("failed to create data directory: %w", err)
	}

	// Find source data directory
	sourceDataDir := findDataDir()

	// Copy CSV files
	csvFiles := []string{"strategies.csv", "indicators.csv", "risk-management.csv", "data-sources.csv", "anti-patterns.csv"}
	for _, csv := range csvFiles {
		src := filepath.Join(sourceDataDir, csv)
		dst := filepath.Join(dataDestDir, csv)

		if err := copyFile(src, dst); err != nil {
			return fmt.Errorf("failed to copy %s: %w", csv, err)
		}
	}
	green.Printf("✅ Copied: 5 CSV files (122 knowledge entries)\n")

	// Create workflow file
	workflowPath := filepath.Join(agentDir, "use-quant-skill.md")
	workflowContent := generateWorkflowContent(aiName)
	if err := os.WriteFile(workflowPath, []byte(workflowContent), 0644); err != nil {
		return fmt.Errorf("failed to create workflow: %w", err)
	}
	green.Printf("✅ Created: %s\n", workflowPath)

	// Create SKILL.md
	skillPath := filepath.Join(sharedDir, "SKILL.md")
	skillContent := generateSkillDoc()
	if err := os.WriteFile(skillPath, []byte(skillContent), 0644); err != nil {
		return fmt.Errorf("failed to create SKILL.md: %w", err)
	}
	green.Printf("✅ Created: %s\n", skillPath)

	// Summary
	fmt.Println()
	cyan.Printf("%s\n", strings.Repeat("=", 100))
	cyan.Println("✅ Initialization Complete!")
	cyan.Printf("%s\n\n", strings.Repeat("=", 100))

	fmt.Println("📁 Structure created:")
	fmt.Println("   .agent/")
	fmt.Println("   └── workflows/")
	fmt.Println("       └── use-quant-skill.md")
	fmt.Println("   .shared/")
	fmt.Println("   └── quant-trading-pro/")
	fmt.Println("       ├── data/           (5 CSV files)")
	fmt.Println("       └── SKILL.md")
	fmt.Println()
	fmt.Println("🚀 Quick Start:")
	fmt.Printf("   quantpro search \"order flow crypto\"\n")
	fmt.Printf("   quantpro search \"stop loss\" -d risk\n")
	fmt.Println()
	fmt.Println("📖 See .shared/quant-trading-pro/SKILL.md for full documentation")
	fmt.Println()

	return nil
}

func copyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, data, 0644)
}

func generateWorkflowContent(aiName string) string {
	return fmt.Sprintf(`---
description: How to use the QuantPro skill for quantitative trading research
---

# Using QuantPro Skill

AI Agent: %s

## Quick Start

### Search Knowledge Base

`+"```bash"+`
# Auto-detect domain
quantpro search "order flow crypto"

# Search specific domain
quantpro search "stop loss kelly" -d risk

# Get more results
quantpro search "rsi bollinger" -d indicator -n 5
`+"```"+`

### Available Domains

- **strategy** (24 entries): Trading strategies from HFT to long-term
- **indicator** (23 entries): Technical indicators with formulas
- **risk** (24 entries): Position sizing, stops, portfolio protection
- **data** (24 entries): Market data sources and costs  
- **anti-pattern** (27 entries): Common mistakes to avoid

## Examples

### Research New Strategy

`+"```bash"+`
# 1. Find strategies
quantpro search "crypto futures intraday" -d strategy

# 2. Get indicators
quantpro search "ema tci microstructure" -d indicator

# 3. Plan risk controls
quantpro search "atr position sizing" -d risk

# 4. Check data needs
quantpro search "tick data order book" -d data

# 5. Avoid mistakes
quantpro search "hft mistakes" -d anti-pattern
`+"```"+`

### Validate Strategy

`+"```bash"+`
# Check for common pitfalls
quantpro search "backtesting overfitting" -d anti-pattern
`+"```"+`

## Documentation

See .shared/quant-trading-pro/SKILL.md for complete documentation.
`, aiName)
}

func generateSkillDoc() string {
	return `# QuantPro Skill Documentation

## Overview

QuantPro provides instant access to curated quantitative trading knowledge across 5 domains:

1. **Strategies** (24 entries): HFT to long-term strategies
2. **Indicators** (23 entries): Technical indicators with formulas
3. **Risk Management** (24 entries): Position sizing and controls
4. **Data Sources** (24 entries): Market data types and vendors
5. **Anti-Patterns** (27 entries): Common mistakes to avoid

## Usage

### Basic Search

` + "```bash" + `
quantpro search "query"
` + "```" + `

### Domain-Specific Search

` + "```bash" + `
quantpro search "query" -d <domain>
` + "```" + `

Domains: strategy, indicator, risk, data, anti-pattern

### Advanced Options

` + "```bash" + `
quantpro search "query" -d domain -n 5        # Get 5 results
quantpro search "query" --data-dir /path      # Custom data path
` + "```" + `

## Knowledge Base

- **Total**: 122 curated entries
- **Format**: CSV (language-agnostic)
- **Search**: BM25 algorithm (<1ms)
- **Auto-detection**: Intelligent domain routing

## Data Location

` + "```" + `
.shared/quant-trading-pro/data/
├── strategies.csv
├── indicators.csv
├── risk-management.csv
├── data-sources.csv
└── anti-patterns.csv
` + "```" + `

## Examples

See workflow file at .agent/workflows/use-quant-skill.md
`
}

func printResults(result *search.Result) {
	cyan := color.New(color.FgCyan, color.Bold)
	yellow := color.New(color.FgYellow)
	green := color.New(color.FgGreen)

	cyan.Printf("\n%s\n", strings.Repeat("=", 100))
	cyan.Printf("QUERY: %s\n", result.Query)
	cyan.Printf("%s\n\n", strings.Repeat("=", 100))

	yellow.Printf("Domain: %s\n", result.Domain)
	yellow.Printf("Found %d results:\n\n", result.Count)

	if result.Count == 0 {
		color.Yellow("No results found. Try a different query or domain.\n")
		return
	}

	for i, r := range result.Results {
		green.Printf("%d. ", i+1)

		var primaryField string
		switch result.Domain {
		case "strategy":
			primaryField = r["Strategy Name"]
		case "indicator":
			primaryField = r["Indicator Name"]
		case "risk":
			primaryField = r["Risk Control"]
		case "data":
			primaryField = r["Data Type"]
		case "anti-pattern":
			primaryField = r["Issue"]
		default:
			primaryField = "Unknown"
		}

		fmt.Printf("%s\n", primaryField)

		fieldCount := 0
		for _, key := range getOrderedKeys(result.Domain) {
			if fieldCount >= 4 {
				break
			}
			if val, ok := r[key]; ok && val != "" && val != primaryField {
				fmt.Printf("   %s: %s\n", key, truncate(val, 100))
				fieldCount++
			}
		}
		fmt.Println()
	}

	color.Cyan("\nTip: Use -d flag to specify domain, -n flag to get more results\n")
}

func getOrderedKeys(domain string) []string {
	switch domain {
	case "strategy":
		return []string{"Category", "Time Horizon", "Complexity", "Best For", "Capital Requirements"}
	case "indicator":
		return []string{"Category", "Formula/Description", "Parameters", "Best For"}
	case "risk":
		return []string{"Category", "Description", "Best For", "Critical For"}
	case "data":
		return []string{"Source", "Frequency", "Best For", "Typical Cost"}
	case "anti-pattern":
		return []string{"Category", "Severity", "Don't", "Do"}
	default:
		return []string{}
	}
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max] + "..."
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
