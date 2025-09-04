package hi

import "rsc.io/quote/v3"

func Hi() string {
	return "Hi! " + quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}
