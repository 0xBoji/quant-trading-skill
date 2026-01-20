package bm25

import (
	"math"
	"regexp"
	"strings"
)

// BM25 implements the BM25 ranking algorithm
type BM25 struct {
	k1         float64
	b          float64
	corpus     [][]string
	docLengths []int
	avgdl      float64
	idf        map[string]float64
	docFreqs   map[string]int
	N          int
}

// New creates a new BM25 instance
func New(k1, b float64) *BM25 {
	return &BM25{
		k1:       k1,
		b:        b,
		idf:      make(map[string]float64),
		docFreqs: make(map[string]int),
	}
}

// Tokenize converts text to tokens (lowercase, alphanumeric, >2 chars)
func (bm *BM25) Tokenize(text string) []string {
	// Remove punctuation, convert to lowercase
	re := regexp.MustCompile(`[^\w\s]`)
	text = re.ReplaceAllString(strings.ToLower(text), " ")

	// Split and filter
	words := strings.Fields(text)
	tokens := make([]string, 0, len(words))
	for _, w := range words {
		if len(w) > 2 {
			tokens = append(tokens, w)
		}
	}
	return tokens
}

// Fit builds the BM25 index from documents
func (bm *BM25) Fit(documents []string) {
	bm.N = len(documents)
	if bm.N == 0 {
		return
	}

	// Tokenize all documents
	bm.corpus = make([][]string, bm.N)
	bm.docLengths = make([]int, bm.N)
	totalLen := 0

	for i, doc := range documents {
		tokens := bm.Tokenize(doc)
		bm.corpus[i] = tokens
		bm.docLengths[i] = len(tokens)
		totalLen += len(tokens)
	}

	bm.avgdl = float64(totalLen) / float64(bm.N)

	// Calculate document frequencies
	for _, doc := range bm.corpus {
		seen := make(map[string]bool)
		for _, word := range doc {
			if !seen[word] {
				bm.docFreqs[word]++
				seen[word] = true
			}
		}
	}

	// Calculate IDF
	for word, freq := range bm.docFreqs {
		bm.idf[word] = math.Log((float64(bm.N-freq)+0.5)/(float64(freq)+0.5) + 1.0)
	}
}

// Result represents a search result with index and score
type Result struct {
	Index int
	Score float64
}

// Score scores all documents against a query
func (bm *BM25) Score(query string) []Result {
	queryTokens := bm.Tokenize(query)
	results := make([]Result, bm.N)

	for idx, doc := range bm.corpus {
		score := 0.0
		docLen := float64(bm.docLengths[idx])

		// Calculate term frequencies
		termFreqs := make(map[string]int)
		for _, word := range doc {
			termFreqs[word]++
		}

		// Calculate BM25 score
		for _, token := range queryTokens {
			if idf, ok := bm.idf[token]; ok {
				tf := float64(termFreqs[token])
				numerator := tf * (bm.k1 + 1.0)
				denominator := tf + bm.k1*(1.0-bm.b+bm.b*docLen/bm.avgdl)
				score += idf * numerator / denominator
			}
		}

		results[idx] = Result{Index: idx, Score: score}
	}

	return results
}
