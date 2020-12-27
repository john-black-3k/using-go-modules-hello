package hello

import "rsc.io/quote/v3"

// Hello is exported
func Hello() string {
	return quote.HelloV3()
}

// Proverb is exported
func Proverb() string {
	return quote.Concurrency()
}
