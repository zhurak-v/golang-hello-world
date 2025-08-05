package jokes

import (
	"math/rand"
)

var jokeList = []string{
	"When the customer says something isn't working, then I work.",
	"What do frontend developers want for Christmas? npm which does not crash the project after the update.",
	"Python saves me from unwanted female attention.",
}

func GetRandomJoke() string {
	return jokeList[rand.Intn(len(jokeList))]
}
