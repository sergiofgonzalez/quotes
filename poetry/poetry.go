package poetry

var quotes = []string{
	"If you feel lonely when you are alone, you're in bad company.",
}

var i = -1

// Quote returns a poetry quote sequentially from the list of available quotes
func Quote() string {
	i = (i + 1) % len(quotes)
	return quotes[i]
}