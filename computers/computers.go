package computers

var quotes = []string{
	`Optimization is by necessity. Performance is by design.`,
	`What a programmer can do in one month, two programmers can do in two months.`,
	`Code is like humour. When you have to explain it it's bad.`,
	`The problem with quick and diry is that the dirty remains long after the quick has been forgotten.`,
}

var i = -1

// Quote returns a different quote from its catalog of quotes using a round-robin strategy
func Quote() string {
	i = (i + 1) % len(quotes)
	return quotes[i]
}