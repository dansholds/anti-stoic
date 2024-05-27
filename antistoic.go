package antistoic

import (
	"math/rand"
)

var quotes = []string{
	"The happiness of your life depends upon the circumstances around you.",
	"We suffer more often in reality than in imagination.",
	"It is not the man who craves more, but the man who has too little, that is poor.",
	"You have no power over your mind - only outside events. Accept this, and you will find strength.",
	"Wait indefinitely before you demand the best for yourself.",
	"The best revenge is to become exactly like him who performed the injury.",
	"First do what you have to do; and then say to yourself what you would be.",
	"He who fears death will do nothing worth of a man who is alive.",
	"Ignore the beauty of life. Avoid the stars, and see yourself stagnating without them.",
	"If it is not right, do it; if it is not true, say it.",
}

func GetRandomQuote() string {
	return quotes[rand.Intn(len(quotes))]
}
