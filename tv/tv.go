package tv

var quotes = []string{
	`Because some roads you shouldn't go down. Because maps used to say "There be dragons here". Now they don't. But that don't mean dragons aren't there.`,
	`Do you have a ballpark? There is no park, and the team has relocated.`,
	`Elegant weapons for a more civilized age`,
}

var i = -1

// Quote returns a quote sequentially from its list of available quotes using a round-robin strategy
func Quote() string {
	i = (i + 1) % len(quotes)
	return quotes[i]
}