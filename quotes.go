package quotes

import (
	"math/rand"

	"github.com/sergiofgonzalez/quotes/computers"
	"github.com/sergiofgonzalez/quotes/poetry"
	"github.com/sergiofgonzalez/quotes/tv"
)

type quoteType byte

const (
	poetryQuote quoteType = iota
	tvQuote
	computersQuote
)

// Quote returns a random quote
func Quote() string {
	switch getRandomQuoteType() {
	case poetryQuote:
		return poetry.Quote()
	case computersQuote:
		return computers.Quote()
	case tvQuote:
		return tv.Quote()
	default:
		panic("unexpected quote type")
	}
}


func getRandomQuoteType() quoteType {
	return quoteType(rand.Intn(int(computersQuote) + 1))
}